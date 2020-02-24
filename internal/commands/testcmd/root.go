package testcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
)

// GetTestCommand returns the top level test command
func GetTestCommand() *cobra.Command {
	activeEnv := config.GetActiveEnv()
	profileNames := activeEnv.GetProfileNames()

	return &cobra.Command{
		Use:              "test",
		Short:            "-",
		Long:             "-",
		Example:          "dockma test",
		Args:             cobra.OnlyValidArgs,
		ValidArgs:        profileNames,
		Hidden:           true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		Run:              runTestCommand,
	}
}

func runTestCommand(cmd *cobra.Command, args []string) {
	fmt.Println("Running test cmd")

	fmt.Println(cmd.ValidArgs)
}
