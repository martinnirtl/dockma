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

// GetInspectCommand returns the top level inspect command
func GetInspectCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "inspect",
		Short:   "Print detailed output of previously executed external command",
		Long:    "Print detailed output of previously executed external command",
		Example: "dockma inspect",
		Args:    cobra.NoArgs,
		Run:     runInspectCommand,
	}
}

func runInspectCommand(cmd *cobra.Command, args []string) {
	filepath := config.GetSubcommandLogfile()

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Nothing to output yet.")
		} else {
			logfileName := viper.GetString("subcommandlogfile")

			utils.ErrorAndExit(fmt.Errorf("Could not read logfile %s", chalk.Underline.TextStyle(logfileName)))
		}
	} else {
		fmt.Println(chalk.Cyan.Color("Here comes the original output:"))
		fmt.Print(string(content))
	}
}
