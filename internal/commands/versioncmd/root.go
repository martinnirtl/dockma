package versioncmd

import (
	"github.com/martinnirtl/dockma/internal/buildinfo"
	"github.com/spf13/cobra"
)

var fullFlag bool

// GetVersionCommand returns the top level version command
func GetVersionCommand() *cobra.Command {
	versionCommand := &cobra.Command{
		Use:              "version",
		Short:            "Print the version number of dockma.",
		Long:             "Print the version number of dockma.",
		Example:          "dockma version",
		Args:             cobra.NoArgs,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {},
		Run:              runVersionCommand,
	}

	versionCommand.Flags().BoolVarP(&fullFlag, "full", "f", false, "print full version description")

	return versionCommand
}

func runVersionCommand(cmd *cobra.Command, args []string) {
	if fullFlag {
		buildinfo.Version.PrintFull()
	} else {
		buildinfo.Version.Print()
	}
}
