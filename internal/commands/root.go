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
	"github.com/martinnirtl/dockma/internal/commands/testcmd"
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

// FIXME make flags work
var verboseFlag bool

// RootCommand is the root command of dockma
var RootCommand = &cobra.Command{
	Use:               "dockma",
	Short:             "Dockma is bringing your docker-compose game to the next level.",
	Long:              `A fast and flexible CLI tool to boost your productivity during development with docker containers built with Go. Full documentation is available at https://dockma.dev`,
	PersistentPreRun:  rootPreRunHook,
	PersistentPostRun: rootPostRunHook,
}

func init() {
	// cobra.OnInitialize(initConfig)

	// TODO FLAGS GO HERE
	RootCommand.PersistentFlags().BoolVarP(&verboseFlag, "verbose", "v", false, "verbose output")
	// RootCommand.PersistentFlags().BoolVar(&authorFlag, "author", false, "print author")

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
	viper.SetDefault("subcommandlogfile", "subcommand.log")
	viper.SetDefault("dockmalogfile", "dockma.log")
	viper.SetDefault("active", "-")
	viper.SetDefault("envs", map[string]interface{}{})

	initConfig()
	initRootCmd()
}

func initRootCmd() {
	RootCommand.AddCommand(completioncmd.GetCompletionCommand())
	RootCommand.AddCommand(configcmd.GetConfigCommand())
	RootCommand.AddCommand(downcmd.GetDownCommand())
	RootCommand.AddCommand(envcmd.GetEnvCommand())
	RootCommand.AddCommand(initcmd.GetInitCommand())
	RootCommand.AddCommand(inspectcmd.GetInspectCommand())
	RootCommand.AddCommand(logscmd.GetLogsCommand())
	RootCommand.AddCommand(profilecmd.GetProfileCommand())
	RootCommand.AddCommand(pscmd.GetPSCommand())
	RootCommand.AddCommand(upcmd.GetUpCommand())
	RootCommand.AddCommand(versioncmd.GetVersionCommand())

	RootCommand.AddCommand(testcmd.GetTestCommand())
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
	if cmd.Name() == "help" {
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
