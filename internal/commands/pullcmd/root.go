package pullcmd

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

// Pull runs git pull in given path and optionally logs output
func Pull(path string, log bool) error {
	err := os.Chdir(path)

	if err != nil {
		return errors.New("could not change working dir")
	}

	var timebridger externalcommand.Timebridger
	if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
		timebridger = spinnertimebridger.New(fmt.Sprintf("Running %sgit pull%s", chalk.Cyan, chalk.ResetColor), "", 14, "cyan")
	}

	var logfile string
	if log {
		logfile = config.GetLogfile()
	}

	_, err = externalcommand.Execute("git pull", timebridger, logfile)

	if err != nil {
		return errors.New("Could not execute 'git pull' in active environment home dir")
	}

	return nil
}

// PullCommand is a top level dockma command
var PullCommand = &cobra.Command{
	Use:   "pull",
	Short: "Run 'git pull' in active environment home dir.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {

		activeEnv := config.GetActiveEnv()

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := config.GetEnvHomeDir(activeEnv)

		err := Pull(envHomeDir, true)

		if err != nil {
			utils.Error(err)
		}
	},
}
