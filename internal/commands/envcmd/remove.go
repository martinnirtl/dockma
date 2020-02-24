package envcmd

import (
	"errors"
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
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
		Example: "dockma envs remove my-env",
		// FIXME rewrite/rethink loading of dockma config for dynamic generation of ValidArgs
		Args: cobra.RangeArgs(0, 1),
		Run:  runRemoveCommand,
	}
}

func runRemoveCommand(cmd *cobra.Command, args []string) {
	env := ""
	if len(args) == 0 {
		env = helpers.GetEnvironment("")
	} else {
		env = helpers.GetEnvironment(args[0])
	}

	sure := survey.Confirm(fmt.Sprintf("Are you sure to remove '%s'", env), false)
	if !sure {
		utils.Abort()
	}

	activeEnv := config.GetActiveEnv()

	if env == activeEnv.GetName() {
		viper.Set("active", "-")
		config.Save(chalk.Yellow.Color("Unset active environment."), errors.New("Failed to unset active environment"))
	}

	config.Save(fmt.Sprintf("Removed environment: %s", chalk.Cyan.Color(env)), errors.New("Failed to remove environment"))

	envs := viper.GetStringMap("envs")
	delete(envs, env)
	viper.Set("envs", envs)
}
