package envcmd

import (
	"errors"
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

func getRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "remove [environment]",
		Aliases: []string{"rm"},
		Short:   "Remove environment",
		Long:    "Remove environment",
		Example: "dockma envs remove",
		Args:    argvalidators.OptionalEnv,
		PreRun:  hooks.RequiresEnv,
		Run:     runRemoveCommand,
	}
}

func runRemoveCommand(cmd *cobra.Command, args []string) {
	var envName string
	if len(args) == 0 {
		envNames := config.GetEnvNames()
		envName = survey.Select("Choose an environment", envNames)
	} else {
		envName = args[0]
	}

	sure := survey.Confirm(fmt.Sprintf("Are you sure to remove %s", chalk.Cyan.Color(envName)), false)
	if !sure {
		utils.Abort()
	}

	activeEnv := config.GetActiveEnv()

	if envName == activeEnv.GetName() {
		viper.Set("active", "-")

		config.Save(chalk.Yellow.Color("Unset active environment."), errors.New("Failed to unset active environment"))
	}

	envs := viper.GetStringMap("envs")
	delete(envs, envName)
	viper.Set("envs", envs)

	config.Save(fmt.Sprintf("Removed environment: %s", chalk.Cyan.Color(envName)), fmt.Errorf("Failed to remove environment: %s", envName))
}
