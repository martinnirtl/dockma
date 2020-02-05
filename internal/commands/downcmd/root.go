package downcmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var DownCommand = &cobra.Command{
	Use:   "down",
	Short: "Stops active environment.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		logfileName := viper.GetString("logfile")
		filepath := utils.GetFullLogfilePath(logfileName)

		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("environments.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); !hideCmdOutput {
			timebridger = spinnertimebridger.New("Running command '%s'", "", 14, "cyan")
		}

		_, err = externalcommand.Execute("docker-compose down", timebridger, filepath)

		if err != nil {
			utils.Error(err)
		}
	},
}
