package commands

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/configcommand"
	"github.com/martinnirtl/dockma/internal/commands/initcommand"
	"github.com/martinnirtl/dockma/internal/commands/inspectcommand"
	"github.com/martinnirtl/dockma/internal/commands/testcommand"
	"github.com/martinnirtl/dockma/internal/commands/upcommand"
	"github.com/martinnirtl/dockma/internal/commands/versioncommand"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// RootCommand is the root command of dockma
var RootCommand = &cobra.Command{
	Use:              "dockma",
	Short:            "Dockma is bringing your docker-compose game to the next level.",
	Long:             `A fast and flexible CLI tool to boost your productivity during development with docker containers built with Go. Full documentation is available at https://dockma.dev`,
	PersistentPreRun: rootPreRunHook,
}

func init() {
	RootCommand.AddCommand(configcommand.ConfigCommand)
	RootCommand.AddCommand(initcommand.InitCommand)
	RootCommand.AddCommand(inspectcommand.InspectCommand)
	RootCommand.AddCommand(testcommand.TestCommand)
	RootCommand.AddCommand(upcommand.UpCommand)
	RootCommand.AddCommand(versioncommand.VersionCommand)

	cobra.OnInitialize(initConfig)

	// FLAGS GO HERE

	// FIXME does not consider envs set via .bash_profile
	viper.SetEnvPrefix("dockma")
	viper.BindEnv("home")

	if homeDir := viper.GetString("home"); homeDir == "" {
		userHome, err := homedir.Dir()
		if err != nil {
			fmt.Printf("%sCould not detect home dir: %s%s\n", chalk.Red, err, chalk.ResetColor)
		}

		viper.SetDefault("home", path.Join(userHome, ".dockma"))
	}

	viper.SetDefault("logfile", "log.txt")
	viper.SetDefault("username", "User")
	viper.SetDefault("activeEnvironment", "-")
	viper.SetDefault("environments", map[string]interface{}{})
}

func initConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")

	dockmaConfig := viper.GetString("home")

	viper.AddConfigPath(dockmaConfig)

	viper.ReadInConfig()
}

func rootPreRunHook(cmd *cobra.Command, args []string) {
	if init := viper.GetTime("init"); init.IsZero() {
		if user, err := user.Current(); err == nil {
			fmt.Printf("Come on, %s! Run %sdockma init%s first to initialize the Dockma CLI.\n", strings.Title(user.Username), chalk.Cyan, chalk.ResetColor)
		} else {
			fmt.Printf("Please run %sdockma init%s first to initialize the Dockma CLI.\n", chalk.Cyan, chalk.ResetColor)
		}

		os.Exit(0)
	}
}

// Execute starts cobra command execution
func Execute() {
	if err := RootCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
