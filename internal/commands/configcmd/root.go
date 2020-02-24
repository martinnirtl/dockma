package configcmd

import (
	"github.com/spf13/cobra"
)

// GetConfigCommand returns the top level config command
func GetConfigCommand() *cobra.Command {
	configCommand := &cobra.Command{
		Use:     "config",
		Aliases: []string{"cfg"},
		Short:   "Dockma configuration details",
		Long:    "Dockma configuration details",
	}

	configCommand.AddCommand(getCatCommand())
	configCommand.AddCommand(getHomeCommand())
	configCommand.AddCommand(getSetCommand())

	return configCommand
}
