package downcmd

import (
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

var DownCommand = &cobra.Command{
	Use:   "down",
	Short: "Stops active environment.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetLogfile()

		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		utils.Error(err)

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
			timebridger = spinnertimebridger.New("Running 'docker-compose down'", fmt.Sprintf("%sSuccessfully executed 'docker-compose down'%s", chalk.Green, chalk.ResetColor), 14, "cyan")
		}

		_, err = externalcommand.Execute("docker-compose down", timebridger, filepath)

		utils.Error(err)
	},
}
