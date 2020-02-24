package profilecmd

import "github.com/spf13/cobra"

// GetProfileCommand returns the top level profile command
func GetProfileCommand() *cobra.Command {
	profileCommand := &cobra.Command{
		Use:     "profile",
		Aliases: []string{"profiles"},
		Short:   "Manage profiles (predefined service selections)",
		Long:    "Manage profiles (predefined service selections)",
	}

	profileCommand.AddCommand(getCreateCommand())
	profileCommand.AddCommand(getDeleteCommand())
	profileCommand.AddCommand(getListCommand())
	profileCommand.AddCommand(getRenameCommand())
	profileCommand.AddCommand(getUpdateCommand())

	return profileCommand
}
