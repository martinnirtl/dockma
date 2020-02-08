package commands

import (
	"fmt"
	"os"
	"os/user"
	"path"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/configcmd"
	"github.com/martinnirtl/dockma/internal/commands/downcmd"
	"github.com/martinnirtl/dockma/internal/commands/envscmd"
	"github.com/martinnirtl/dockma/internal/commands/initcmd"
	"github.com/martinnirtl/dockma/internal/commands/inspectcmd"
	"github.com/martinnirtl/dockma/internal/commands/logscmd"
	"github.com/martinnirtl/dockma/internal/commands/profilecmd"
	"github.com/martinnirtl/dockma/internal/commands/pscmd"
	"github.com/martinnirtl/dockma/internal/commands/pullcmd"
	"github.com/martinnirtl/dockma/internal/commands/upcmd"
	"github.com/martinnirtl/dockma/internal/commands/versioncmd"
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
	RootCommand.AddCommand(configcmd.ConfigCommand)
	RootCommand.AddCommand(downcmd.DownCommand)
	RootCommand.AddCommand(envscmd.EnvsCommand)
	RootCommand.AddCommand(initcmd.InitCommand)
	RootCommand.AddCommand(inspectcmd.InspectCommand)
	RootCommand.AddCommand(logscmd.LogsCommand)
	RootCommand.AddCommand(profilecmd.ProfileCommand)
	RootCommand.AddCommand(pscmd.PSCommand)
	RootCommand.AddCommand(pullcmd.PullCommand)
	RootCommand.AddCommand(upcmd.UpCommand)
	RootCommand.AddCommand(versioncmd.VersionCommand)

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

	viper.SetDefault("username", "User")
	viper.SetDefault("hidesubcommandoutput", false)
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

// Execute starts cobra command execution
func Execute() {
	RootCommand.Execute()
}
