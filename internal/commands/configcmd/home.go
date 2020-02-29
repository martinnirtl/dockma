package configcmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getHomeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "home",
		Short:   "Print home dir of dockma config",
		Long:    "Print home dir of dockma config",
		Example: "dockma config home",
		Args:    cobra.NoArgs,
		Run:     runHomeCommand,
	}
}

func runHomeCommand(cmd *cobra.Command, args []string) {
	fmt.Printf("Dockma config home dir: %s\n", chalk.Cyan.Color(viper.GetString("home")))
}
