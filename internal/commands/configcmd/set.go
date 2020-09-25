package configcmd

import (
	"fmt"
	"strings"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var configVars []string = []string{"hidecommandoutput"}

func getSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "set",
		Short:     "Set variables of dockma configuration",
		Long:      "Set variables of dockma configuration",
		Example:   "dockma config set",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: configVars,
		Run:       runSetCommand,
	}
}

func runSetCommand(cmd *cobra.Command, args []string) {
	var selected []string

	if len(args) == 0 {
		options := []string{
			fmt.Sprintf("Hide command output: %t", viper.GetBool("hidecommandoutput")),
		}

		selected = survey.MultiSelect("Select config items to change", options, nil)
	} else {
		selected = args
	}

	for _, varnameRaw := range selected {
		varname := strings.Split(varnameRaw, ":")

		setConfigVar(varname[0])

		message := fmt.Sprintf("Set %s: %s", chalk.Cyan.Color(varname[0]), chalk.Cyan.Color(viper.GetString(varname[0])))
		err := fmt.Errorf("Failed to set: %s", varname[0])
		config.Save(message, err)
	}
}

func setConfigVar(selection string) {
	switch selection {
	case "Hide command output":
		hide := survey.Confirm("Hide output of external commands [true: show output only on error, false: always pipe output]", viper.GetBool("hidecommandoutput"))

		viper.Set("hidecommandoutput", hide)
	}
}
