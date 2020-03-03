package configcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
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
	fmt.Println(config.GetHomeDir())
}
