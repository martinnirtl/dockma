package profilecmd

import "github.com/spf13/cobra"

// GetProfileCommand returns the top level profile command
func GetProfileCommand() *cobra.Command {
	profileCommand := &cobra.Command{
		Use:     "profile",
		Aliases: []string{"profiles", "pr"},
		Short:   "Manage profiles (named service selections)",
		Long:    "Manage profiles (named service selections)",
	}

	profileCommand.AddCommand(getCreateCommand())
	profileCommand.AddCommand(getDeleteCommand())
	profileCommand.AddCommand(getListCommand())
	profileCommand.AddCommand(getRenameCommand())
	profileCommand.AddCommand(getUpdateCommand())

	return profileCommand
}
