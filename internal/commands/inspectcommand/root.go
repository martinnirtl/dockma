package inspectcommand

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var InspectCommand = &cobra.Command{
	Use:              "inspect",
	Short:            "Print detailed output of previously executed command",
	Long:             "-",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {},
	Run: func(cmd *cobra.Command, args []string) {
		filename := viper.GetString("logfile")
		filepath := utils.GetFullLogfilePath(filename)

		content, err := ioutil.ReadFile(filepath)

		if err != nil {
			if errors.Is(err, os.ErrNotExist) {
				fmt.Printf("%sNothing to output yet. Sry, %s!%s\n", chalk.Cyan, viper.GetString("username"), chalk.ResetColor)
			} else {
				fmt.Printf("%sError: %s%s\n", chalk.Red, viper.GetString("username"), chalk.ResetColor)
			}
		} else {
			fmt.Printf("%sHere come the logs:%s\n", chalk.Cyan, chalk.ResetColor)

			fmt.Print(string(content))
		}
	},
}
