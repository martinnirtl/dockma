package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func GetVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the version number of Dockma.",
		Long:  "Print the version number of Dockma.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Dockma v1.0.0")
		},
	}
}
