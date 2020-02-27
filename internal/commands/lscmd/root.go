package logscmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/spf13/cobra"
)

var listFlag bool
var allFlag bool

// GetLSCommand returns the top level ls command
func GetLSCommand() *cobra.Command {
	lsCommand := &cobra.Command{
		Use:     "ls",
		Short:   "List environments, profiles and services",
		Long:    "List environments, profiles and services",
		Example: "dockma ls -l",
		Hidden:  true,
		Args:    argvalidators.OnlyServices,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runLsCommand,
	}

	lsCommand.Flags().BoolVarP(&listFlag, "list", "l", false, "output as list")
	lsCommand.Flags().BoolVarP(&allFlag, "all", "a", false, "include profiles and services")

	return lsCommand
}

func runLsCommand(cmd *cobra.Command, args []string) {
	fmt.Println("not yet implemented")
}
