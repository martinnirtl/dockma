package configcmd

import (
	"fmt"
	"io/ioutil"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var catCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Print config.json of Dockma.",
	Long:    "Print config.json of Dockma.",
	Example: "dockma config cat",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetDockmaFilepath(("config.json"))

		content, err := ioutil.ReadFile(filepath)

		if err != nil {
			fmt.Printf("%sError: Could not read config file!%s\n", chalk.Red, chalk.ResetColor)
		} else {
			fmt.Printf("%sHere comes the dockma configuration file:%s\n", chalk.Cyan, chalk.ResetColor)

			fmt.Println(string(content))
		}
	},
}

func init() {
	ConfigCommand.AddCommand(catCmd)
}
