package envcmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getPullCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "pull [environment]",
		Short:   "Run 'git pull' in environment home dir",
		Long:    "Run 'git pull' in environment home dir",
		Example: "dockma env pull",
		Args:    argvalidators.OptionalEnv,
		PreRun:  hooks.RequiresEnv,
		Run:     runPullCommand,
	}
}

func runPullCommand(cmd *cobra.Command, args []string) {
	// activeEnv := config.GetActiveEnv()

	envName := ""
	if len(args) == 0 {
		// if activeEnv.GetName() != "-" {
		// 	pullActive := survey.Confirm(fmt.Sprintf("Pull active env %s", chalk.Cyan.Color(activeEnv.GetName())), true)

		// 	if pullActive {
		// 		envName = activeEnv.GetName()
		// 	}
		// }

		if envName == "" {
			envNames := config.GetEnvNames()
			envName = survey.Select("Choose an environment", envNames)
		}
	} else {
		envName = args[0]
	}

	env, err := config.GetEnv(envName)
	utils.ErrorAndExit(err)

	envHomeDir := env.GetHomeDir()

	if !env.IsGitBased() {
		err := os.Chdir(envHomeDir)
		if err != nil {
			fmt.Println(err)
			utils.ErrorAndExit(fmt.Errorf("Could not change to path %s", chalk.Underline.TextStyle(envHomeDir)))
		}

		// FIXME duplicated code > see env init cmd
		if _, err := os.Stat(".git"); !os.IsNotExist(err) {
			pull := survey.Select(fmt.Sprintf("Environment is now git-based. Run %s before %s", chalk.Cyan.Color("git pull"), chalk.Cyan.Color("dockma up")), []string{"auto", "optional", "manual", "off"})

			viper.Set(fmt.Sprintf("envs.%s.pull", envName), pull)

			config.Save(fmt.Sprintf("Saved pull setting for environment: %s", chalk.Cyan.Color(envName)), errors.New("Failed to save pull setting"))
		} else {
			fmt.Printf("Environment %s is %s.\n", chalk.Cyan.Color(envName), chalk.Red.Color("not a git repository"))

			os.Exit(0)
		}
	}

	hideCmdOutput := config.GetHideSubcommandOutputSetting()

	output, err := Pull(envHomeDir, hideCmdOutput, true)
	if err != nil && hideCmdOutput {
		fmt.Print(string(output))
	}
	utils.ErrorAndExit(err)

	utils.Success(fmt.Sprintf("Pulled environment: %s", env.GetName()))
}

// Pull runs git pull in given path and optionally logs output
func Pull(path string, hideCmdOutput bool, writeToDockmaLog bool) (output []byte, err error) {
	chdirErr := os.Chdir(path)
	if chdirErr != nil {
		err = errors.New("Could not change working dir")

		return
	}

	var timebridger externalcommand.Timebridger
	if hideCmdOutput {
		timebridger = spinnertimebridger.Default(fmt.Sprintf("Running %s", chalk.Cyan.Color("git pull")))
	}

	var logfile string
	if writeToDockmaLog {
		logfile = config.GetSubcommandLogfile()
	}

	output, err = externalcommand.Execute("git pull", timebridger, logfile)

	// activeEnv.SetUpdated() // TODO make config to object

	return
}
