package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "set [environment]",
		Short:   "Set active environment",
		Long:    "Set active environment",
		Example: "dockma envs set",
		Args:    argvalidators.OptionalEnv,
		PreRun:  hooks.RequiresEnv,
		Run:     runSetCommand,
	}
}

func runSetCommand(cmd *cobra.Command, args []string) {
	envName := ""
	if len(args) == 0 {
		envNames := config.GetEnvNames()
		envName = survey.Select("Choose an environment", envNames)
	} else {
		envName = args[0]
	}

	activeEnv := config.GetActiveEnv()

	if envName == activeEnv.GetName() {
		fmt.Printf("Environment already set active: %s\n", chalk.Cyan.Color(activeEnv.GetName()))

		return
	}

	if activeEnv.IsRunning() {
		utils.Warn(fmt.Sprintf("Leaving running environment: %s", activeEnv.GetName()))
	}

	viper.Set("active", envName)

	config.Save(fmt.Sprintf("New active environment: %s (old: %s)", chalk.Cyan.Color(envName), activeEnv.GetName()), fmt.Errorf("Failed to set active environment"))

	env, err := config.GetEnv(envName)
	utils.ErrorAndExit(err)

	if env.IsRunning() {
		utils.Warn("New active environment is presently running.")
	}
}
