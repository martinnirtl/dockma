package initcommand

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strings"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var InitCommand = &cobra.Command{
	Use:              "init",
	Short:            "Initialize the Dockma CLI",
	Long:             "-",
	PersistentPreRun: initPreRunHook, // used to override root PreRun func
	Run:              initCommandHandler,
}

func initPreRunHook(cmd *cobra.Command, args []string) {
	if init := viper.GetTime("init"); !init.IsZero() {
		proceed := false
		err := survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("%sDockma CLI has already been initialized!%s Do you want to proceed", chalk.Yellow, chalk.ResetColor),
			Default: false,
		}, &proceed, survey.WithValidator(survey.Required))

		if err != nil || !proceed {
			utils.Abort()
		}
	} else {
		accept := false
		err := survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Dockma CLI config will be stored at: %s", viper.GetString("home")),
			Default: true,
		}, &accept, survey.WithValidator(survey.Required))

		if err != nil {
			utils.Abort()
		} else if !accept {
			fmt.Printf("Ok, you can set the config location via %sDOCKMA_HOME%s environment variable.\n", chalk.Cyan, chalk.ResetColor)

			os.Exit(0)
		}
	}
}

func initCommandHandler(cmd *cobra.Command, args []string) {
	var username string
	survey.AskOne(&survey.Input{
		Message: "What is your name",
		Default: "User",
	}, &username, survey.WithValidator(func(val interface{}) error {
		switch text := val.(type) {
		case string:
			if len(text) > 0 {
				return nil
			}
		default:
			return errors.New("invalid input for username")
		}

		return nil
	}))

	viper.Set("username", strings.Title(username))
	viper.Set("init", time.Now())

	home := viper.GetString("home")

	if err := os.MkdirAll(home, os.FileMode(0755)); err != nil {
		fmt.Printf("%sCould not create config dir: %s%s\n", chalk.Red, err, chalk.ResetColor)

		os.Exit(1)
	}

	filepath := path.Join(home, "config.json")

	if err := viper.WriteConfigAs(filepath); err != nil {
		fmt.Printf("%sCould not save config.json at: %s%s\n", chalk.Red, home, chalk.ResetColor)

		os.Exit(1)
	}
}
