package versioncmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCommand implements the top level version command
var VersionCommand = &cobra.Command{
	Use:              "version",
	Short:            "Print the version number of dockma.",
	Long:             "Print the version number of dockma.",
	Example:          "dockma version",
	Args:             cobra.NoArgs,
	Hidden:           true,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		// TODO use version from build flags
		fmt.Println("🐳 Dockma v0.0.0")
	},
}
