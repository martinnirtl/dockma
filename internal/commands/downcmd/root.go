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
)

// DownCommand implements the top level down command
var DownCommand = &cobra.Command{
	Use:     "down",
	Short:   "Stops active environment.",
	Long:    "Stops active environment.",
	Example: "dockma down",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetLogfile()

		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		utils.ErrorAndExit(err)

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
			timebridger = spinnertimebridger.New("Running 'docker-compose down'", 14, "cyan")
		}

		output, err := externalcommand.Execute("docker-compose down", timebridger, filepath)

		utils.Error(err)
		if err != nil {
			fmt.Println(output)

			os.Exit(0)
		}

		utils.Success("Successfully executed 'docker-compose down'")
	},
}
