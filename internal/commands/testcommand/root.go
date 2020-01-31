package testcommand

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TestCommand = &cobra.Command{
	Use:              "test",
	Short:            "Just for testing.",
	Long:             "-",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		logfileName := viper.GetString("logfile")
		filepath := utils.GetFullLogfilePath(logfileName)

		err := externalcommand.Execute("ls -la", filepath)

		if err != nil {
			fmt.Println(err)
		}
	},
}
