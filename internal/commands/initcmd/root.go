package initcmd

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"time"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// GetInitCommand returns the top level init command
func GetInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:              "init",
		Short:            "Initialize the Dockma CLI.",
		Long:             "Initialize the Dockma CLI.",
		Example:          "dockma init",
		Args:             cobra.NoArgs,
		PersistentPreRun: initPreRunHook, // used to override root PreRun func
		Run:              runInitCommand,
	}
}

func initPreRunHook(cmd *cobra.Command, args []string) {
	if init := config.GetInitTime(); !init.IsZero() {
		proceed := survey.Confirm(fmt.Sprintf("%sDockma CLI has already been initialized!%s Do you want to proceed", chalk.Yellow, chalk.ResetColor), false)
		if !proceed {
			utils.Abort()
		}
	} else {
		accept := survey.Confirm(fmt.Sprintf("Dockma CLI config will be stored at: %s", config.GetHomeDir()), true)

		if !accept {
			fmt.Printf("Ok, you can change the config default location by setting %sDOCKMA_HOME%s environment variable.\n", chalk.Cyan, chalk.ResetColor)

			os.Exit(0)
		}
	}
}

func runInitCommand(cmd *cobra.Command, args []string) {
	username := "User"
	if sysUser, err := user.Current(); err == nil {
		username = sysUser.Username
	}

	username = survey.InputName("What is your name", username)

	viper.Set("username", username)
	viper.Set("init", time.Now())

	home := config.GetHomeDir()

	if err := os.MkdirAll(home, os.FileMode(0755)); err != nil {
		utils.ErrorAndExit(fmt.Errorf("Could not create config dir: %s", err))
	}

	filepath := path.Join(home, "config.json")

	if err := viper.WriteConfigAs(filepath); err != nil {
		utils.ErrorAndExit(fmt.Errorf("Could not save config.json at: %s", home))
	}
}
