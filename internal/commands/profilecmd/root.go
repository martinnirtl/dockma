package profilecmd

import "github.com/spf13/cobra"

// ProfileCommand implements the top level profile command
var ProfileCommand = &cobra.Command{
	Use:     "profile",
	Aliases: []string{"profiles"},
	Short:   "Manage profiles (predefined service selections)",
	Long:    "Manage profiles (predefined service selections)",
}
