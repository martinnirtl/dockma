package commands

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/completioncmd"
	"github.com/martinnirtl/dockma/internal/commands/configcmd"
	"github.com/martinnirtl/dockma/internal/commands/downcmd"
	"github.com/martinnirtl/dockma/internal/commands/envcmd"
	"github.com/martinnirtl/dockma/internal/commands/initcmd"
	"github.com/martinnirtl/dockma/internal/commands/inspectcmd"
	"github.com/martinnirtl/dockma/internal/commands/logscmd"
	"github.com/martinnirtl/dockma/internal/commands/profilecmd"
	"github.com/martinnirtl/dockma/internal/commands/pscmd"
	"github.com/martinnirtl/dockma/internal/commands/upcmd"
	"github.com/martinnirtl/dockma/internal/commands/versioncmd"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// RootCommand is the root command of dockma
var RootCommand = &cobra.Command{
	Use:               "dockma",
	Short:             "Dockma is bringing your docker-compose game to the next level.",
	Long:              `A fast and flexible CLI tool to boost your productivity during development with docker containers built with Go. Full documentation is available at https://dockma.dev`,
	PersistentPreRun:  rootPreRunHook,
	PersistentPostRun: rootPostRunHook,
}

func init() {
	RootCommand.AddCommand(completioncmd.CompletionCommand)
	RootCommand.AddCommand(configcmd.ConfigCommand)
	RootCommand.AddCommand(downcmd.DownCommand)
	RootCommand.AddCommand(envcmd.EnvCommand)
	RootCommand.AddCommand(initcmd.InitCommand)
	RootCommand.AddCommand(inspectcmd.InspectCommand)
	RootCommand.AddCommand(logscmd.LogsCommand)
	RootCommand.AddCommand(profilecmd.ProfileCommand)
	RootCommand.AddCommand(pscmd.PSCommand)
	RootCommand.AddCommand(upcmd.UpCommand)
	RootCommand.AddCommand(versioncmd.VersionCommand)

	cobra.OnInitialize(initConfig)

	// TODO FLAGS GO HERE
	// author flag
	// version flag

	// TODO behavior not clear. no consideration of envs set via .bash_profile !?
	viper.SetEnvPrefix("dockma")
	viper.BindEnv("home")

	if homeDir := viper.GetString("home"); homeDir == "" {
		userHome, err := homedir.Dir()
		if err != nil {
			utils.Error(fmt.Errorf("Could not detect home dir: %s", err))
		}

		viper.SetDefault("home", path.Join(userHome, ".dockma"))
	}

	viper.SetDefault("username", "User")
	viper.SetDefault("hidesubcommandoutput", true)
	viper.SetDefault("logfile", "log.txt")
	viper.SetDefault("active", "-")
	viper.SetDefault("envs", map[string]interface{}{})
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	dockmaConfig := viper.GetString("home")

	viper.AddConfigPath(dockmaConfig)

	// NOTE read errors get ignored and execution is mainly prevented by rootPreRunHook func
	_ = viper.ReadInConfig()
}

func rootPreRunHook(cmd *cobra.Command, args []string) {
	// enable printing of help
	if "help" == cmd.Name() {
		return
	}

	if init := viper.GetTime("init"); init.IsZero() {
		if user, err := user.Current(); err == nil {
			fmt.Printf("Come on, %s! Run %sdockma init%s first to initialize the Dockma CLI.\n", strings.Title(user.Username), chalk.Cyan, chalk.ResetColor)
		} else {
			fmt.Printf("Please run %sdockma init%s first to initialize the Dockma CLI.\n", chalk.Cyan, chalk.ResetColor)
		}

		os.Exit(0)
	}
}

func rootPostRunHook(cmd *cobra.Command, args []string) {
	if config.SaveConfig {
		messages, errors, err := config.SaveNow()
		utils.Error(err)

		if err != nil {
			helpers.PrintErrorList(errors)
		} else {
			helpers.PrintMessageList(messages)
		}
	}
}

// Execute starts cobra command execution
func Execute() {
	RootCommand.Execute()
}
