package envcmd

import (
	"github.com/spf13/cobra"
)

// GetEnvCommand returns the top level envs command
func GetEnvCommand() *cobra.Command {
	envCommand := &cobra.Command{
		Use:     "env",
		Aliases: []string{"envs"},
		Short:   "Environments reflect docker-compose based projects",
		Long:    "Environments reflect docker-compose based projects",
	}

	envCommand.AddCommand(getInitCommand())
	envCommand.AddCommand(getListCommand())
	envCommand.AddCommand(getPullCommand())
	envCommand.AddCommand(getRemoveCommand())
	// envCommand.AddCommand(getRenameCommand())
	envCommand.AddCommand(getSetCommand())

	return envCommand
}
