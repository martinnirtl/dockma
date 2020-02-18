package configcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var homeCmd = &cobra.Command{
	Use:     "home",
	Short:   "Print home dir of Dockma config",
	Long:    "Print home dir of Dockma config",
	Example: "dockma config home",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Dockma config home dir: %s\n", chalk.Cyan.Color(viper.GetString("home")))
	},
}

func init() {
	ConfigCommand.AddCommand(homeCmd)
}
