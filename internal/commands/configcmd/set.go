package configcmd

import (
	"fmt"
	"strings"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var setCmd = &cobra.Command{
	Use:     "set",
	Short:   "Set Dockma config vars in an interactive walkthrough.",
	Long:    "Set Dockma config vars in an interactive walkthrough.",
	Example: "dockma config set",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		options := []string{
			fmt.Sprintf("hidesubcommandoutput: %t", viper.GetBool("hidesubcommandoutput")),
			fmt.Sprintf("logfile: %s", viper.GetString("logfile")),
			fmt.Sprintf("color: %s", viper.GetString("color")),
			fmt.Sprintf("username: %s", viper.GetString("username")),
		}

		selected := survey.MultiSelect("Select config items to change", options, nil)

		for _, varnameRaw := range selected {
			varname := strings.Split(varnameRaw, ":")

			setConfigVar(varname[0])
		}

		if len(selected) > 0 {
			config.Save()
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
		logfile := survey.InputName("Enter name of logfile [stored in dockma home dir]", viper.GetString("logfile"))

		viper.Set("logfile", logfile)

	case "color":
		color := survey.Select("Select a new primary color", config.PrimaryColors)

		viper.Set("color", color)

	case "username":
		username := survey.InputName("Enter new username", viper.GetString("username"))

		viper.Set("username", username)
	}
}
