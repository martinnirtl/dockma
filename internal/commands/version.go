package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:              "version",
	Short:            "Print the version number of Dockma.",
	Long:             "Print the version number of Dockma.",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Dockma v1.0.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
