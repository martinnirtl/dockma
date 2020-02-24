package envcmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func getPullCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "pull",
		Short:   "Run 'git pull' in environment home dir",
		Long:    "Run 'git pull' in environment home dir",
		Example: "dockma env pull",
		Args:    cobra.RangeArgs(0, 1),
		Run:     runPullCommand,
	}
}

func runPullCommand(cmd *cobra.Command, args []string) {
	envName := ""
	if len(args) == 0 {
		envName = helpers.GetEnvironment("")
	} else {
		envName = helpers.GetEnvironment(args[0])
	}

	env, err := config.GetEnv(envName)
	utils.ErrorAndExit(err)

	envHomeDir := env.GetHomeDir()

	if !env.IsGitBased() {
		fmt.Printf("Environment %s is not a git repository.\n", chalk.Cyan.Color(envName))

		os.Exit(0)
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
		logfile = config.GetLogfile()
	}

	output, cmdErr := externalcommand.Execute("git pull", timebridger, logfile)

	if cmdErr != nil {
		err = errors.New("Could not execute 'git pull'")
	}

	// activeEnv.SetUpdated() // TODO make config to object

	return
}
