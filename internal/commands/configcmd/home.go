package configcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var homeCmd = &cobra.Command{
	Use:     "home",
	Short:   "Print home dir of dockma config",
	Long:    `-`,
	Example: "dockma config home",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Dockma config home dir: %s%s%s\n", chalk.Cyan, viper.GetString("home"), chalk.ResetColor)
	},
}

func init() {
	ConfigCommand.AddCommand(homeCmd)
}
