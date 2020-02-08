package configcmd

import (
	"github.com/spf13/cobra"
)

// ConfigCommand implements the top level config command
var ConfigCommand = &cobra.Command{
	Use:     "config",
	Short:   "Dockma configuration details.",
	Long:    "Dockma configuration details.",
	Aliases: []string{"cfg"},
}
