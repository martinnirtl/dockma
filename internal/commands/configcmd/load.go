package configcmd

import "github.com/spf13/cobra"

// TODO finish 'config load' cmd
var loadCmd = &cobra.Command{
	Use:     "load",
	Short:   "load dockma config.json",
	Long:    `-`,
	Example: "dockma config cat",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ConfigCommand.AddCommand(loadCmd)
}
