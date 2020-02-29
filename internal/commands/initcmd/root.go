package initcmd

import (
	"errors"
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
		Short:            "Initialize dockma",
		Long:             "Initialize dockma",
		Example:          "dockma init",
		Args:             cobra.NoArgs,
		PersistentPreRun: initPreRunHook, // used to override root PreRun func
		Run:              runInitCommand,
	}
}

func initPreRunHook(cmd *cobra.Command, args []string) {
	if init := config.GetInitTime(); !init.IsZero() {
		proceed := survey.Confirm(fmt.Sprintf("%s Do you want to proceed", chalk.Yellow.Color("Dockma has already been initialized!")), false)
		if !proceed {
			utils.Abort()
		}
	} else {
		accept := survey.Confirm(fmt.Sprintf("Dockma config will be stored at: %s", config.GetHomeDir()), true)

		if !accept {
			fmt.Printf("Dockma's config location can be set by %s environment variable.\n", chalk.Cyan.Color("DOCKMA_HOME"))

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
		fmt.Println(err)
		utils.ErrorAndExit(errors.New("Could not create config dir"))
	}

	filepath := path.Join(home, "config.json")

	if err := viper.WriteConfigAs(filepath); err != nil {
		fmt.Println(err)
		utils.ErrorAndExit(fmt.Errorf("Could not save config.json to %s", chalk.Underline.TextStyle(home)))
	}

	fmt.Printf("%s has been initialized successfully!\n\nStart with adding a new environment by %s or run %s for some little docs.\n", chalk.Cyan.Color("Dockma"), chalk.Cyan.Color("dockma env init"), chalk.Cyan.Color("dockma help"))
}
