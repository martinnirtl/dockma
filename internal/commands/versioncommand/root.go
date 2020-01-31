package versioncommand

import (
	"fmt"

	"github.com/spf13/cobra"
)

// VersionCommand is the top level version command
var VersionCommand = &cobra.Command{
	Use:              "version",
	Short:            "Print the version number of Dockma.",
	Long:             "Print the version number of Dockma.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dockma v1.0.0")
	},
}
