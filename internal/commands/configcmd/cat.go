package configcmd

import "github.com/spf13/cobra"

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Concatinate config.json of dockma",
	Long:    `-`,
	Example: "dockma config cat",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ConfigCommand.AddCommand(catCmd)
}
