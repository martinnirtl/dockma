package upcommand

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UpCommand = &cobra.Command{
	Use:              "up",
	Short:            "Runs active environment with service selection",
	Long:             "-",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		// logfileName := viper.GetString("logfile")
		// filepath := utils.GetFullLogfilePath(logfileName)

		activeEnv := viper.GetString("activeEnvironment")
		envHomeDir := viper.GetString(fmt.Sprintf("environments.%s.home", activeEnv))
		os.Chdir(envHomeDir)

		err := externalcommand.Execute("docker-compose up -d", "")

		if err != nil {
			fmt.Println(err)
		}
	},
}
