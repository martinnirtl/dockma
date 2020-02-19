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

var configVars []string = []string{"hidesubcommandoutput", "logfile", "username"}

var setCmd = &cobra.Command{
	Use:       "set",
	Short:     "Set Dockma config vars in an interactive walkthrough",
	Long:      "Set Dockma config vars in an interactive walkthrough",
	Example:   "dockma config set",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: configVars,
	Run: func(cmd *cobra.Command, args []string) {
		var selected []string

		if len(args) == 0 {
			options := []string{
				fmt.Sprintf("hidesubcommandoutput: %t", viper.GetBool("hidesubcommandoutput")),
				fmt.Sprintf("logfile: %s", viper.GetString("logfile")),
				fmt.Sprintf("username: %s", viper.GetString("username")),
			}

			selected = survey.MultiSelect("Select config items to change", options, nil)
		} else {
			selected = args
		}

		for _, varnameRaw := range selected {
			varname := strings.Split(varnameRaw, ":")

			setConfigVar(varname[0])

			message := fmt.Sprintf("Set %s: %s", chalk.Cyan.Color(varname[0]), chalk.Cyan.Color(viper.GetString(varname[0])))
			err := fmt.Errorf("Failed to set '%s'", varname[0])
			config.Save(message, err)
		}
	},
}

func init() {
	ConfigCommand.AddCommand(setCmd)
}

func setConfigVar(varname string) {
	switch varname {
	case "hidesubcommandoutput":
		hide := survey.Confirm("Hide output of external commands [true: show output only on error, false: always pipe output]", viper.GetBool("hidesubcommandoutput"))

		viper.Set("hidesubcommandoutput", hide)

	case "logfile":
		logfile := survey.InputName("Enter name of logfile [stored in Dockma home dir]", viper.GetString("logfile"))

		viper.Set("logfile", logfile)

	case "username":
		username := survey.InputName("Enter new username", viper.GetString("username"))

		viper.Set("username", username)
	}
}
