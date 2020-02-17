package envcmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var pullCommand = &cobra.Command{
	Use:     "pull",
	Short:   "Run 'git pull' in active environment home dir",
	Long:    "Run 'git pull' in active environment home dir",
	Example: "dockma env pull",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {

		activeEnv := config.GetActiveEnv()

		if activeEnv.GetName() == "-" {
			utils.PrintNoEnvs()
		}

		envHomeDir := activeEnv.GetHomeDir()

		err := Pull(envHomeDir, true)

		utils.ErrorAndExit(err)

		utils.Success(fmt.Sprintf("Successfully pulled env: %s", activeEnv.GetName()))
	},
}

func init() {
	EnvCommand.AddCommand(pullCommand)
}

// Pull runs git pull in given path and optionally logs output
func Pull(path string, log bool) error {
	err := os.Chdir(path)

	if err != nil {
		return errors.New("Could not change working dir")
	}

	var timebridger externalcommand.Timebridger
	if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
		timebridger = spinnertimebridger.New(fmt.Sprintf("Running %sgit pull%s", chalk.Cyan, chalk.ResetColor), 14, "cyan")
	}

	var logfile string
	if log {
		logfile = config.GetLogfile()
	}

	_, err = externalcommand.Execute("git pull", timebridger, logfile)

	if err != nil {
		return errors.New("Could not execute 'git pull' in active environment home dir")
	}

	// activeEnv.SetUpdated() // TODO make config to object

	return nil
}
