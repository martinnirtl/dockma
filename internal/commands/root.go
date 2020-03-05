package commands

import (
	"fmt"
	"path"

	"github.com/martinnirtl/dockma/internal/commands/completioncmd"
	"github.com/martinnirtl/dockma/internal/commands/configcmd"
	"github.com/martinnirtl/dockma/internal/commands/downcmd"
	"github.com/martinnirtl/dockma/internal/commands/envcmd"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/commands/initcmd"
	"github.com/martinnirtl/dockma/internal/commands/inspectcmd"
	"github.com/martinnirtl/dockma/internal/commands/logscmd"
	"github.com/martinnirtl/dockma/internal/commands/profilecmd"
	"github.com/martinnirtl/dockma/internal/commands/pscmd"
	"github.com/martinnirtl/dockma/internal/commands/restartcmd"
	"github.com/martinnirtl/dockma/internal/commands/scriptcmd"
	"github.com/martinnirtl/dockma/internal/commands/testcmd"
	"github.com/martinnirtl/dockma/internal/commands/upcmd"
	"github.com/martinnirtl/dockma/internal/commands/versioncmd"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// FIXME make flags work
var verboseFlag bool

// RootCommand is the root command of dockma
var RootCommand = &cobra.Command{
	Use:               "dockma",
	Short:             "Dockma brings your docker-compose game to the next level.",
	Long:              `A fast and flexible CLI tool to boost your productivity during development in docker-compose based environments.`,
	PersistentPreRun:  hooks.RequiresInit,
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
	RootCommand.AddCommand(restartcmd.GetRestartCommand())
	RootCommand.AddCommand(scriptcmd.GetScriptCommand())
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
