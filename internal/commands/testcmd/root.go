package testcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TestCommand = &cobra.Command{
	Use:   "test",
	Short: "Just for testing.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		logfileName := viper.GetString("logfile")
		filepath := utils.GetFullLogfilePath(logfileName)

		err := externalcommand.Execute("docker-compose up -d pg", filepath)

		if err != nil {
			fmt.Println(err)
		}
	},
}
