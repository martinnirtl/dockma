package profilecmd

import "github.com/spf13/cobra"

// ProfileCommand is the top level config command
var ProfileCommand = &cobra.Command{
	Use:   "profile",
	Short: "Manage profiles (predefined service selections).",
	Long:  "-",
}
