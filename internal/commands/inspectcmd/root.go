package inspectcmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// InspectCommand implements the top level inspect command
var InspectCommand = &cobra.Command{
	Use:   "inspect",
	Short: "Print detailed output of previously executed command [up|down|pull].",
	Long:  "Print detailed output of previously executed command [up|down|pull].",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetLogfile()

		content, err := ioutil.ReadFile(filepath)

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("%sNothing to output yet. Sry, %s!%s\n", chalk.Cyan, viper.GetString("username"), chalk.ResetColor)
			} else {
				utils.ErrorAndExit(fmt.Errorf("Could not read logfile: %s", viper.GetString("logfile")))
			}
		} else {
			fmt.Printf("%sHere come the logs:%s\n", chalk.Cyan, chalk.ResetColor)

			fmt.Print(string(content))
		}
	},
}
