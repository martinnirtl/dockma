package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var setCmd = &cobra.Command{
	Use:     "set [environment]",
	Short:   "Set active environment",
	Long:    "Set active environment",
	Example: "dockma envs set",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		env := ""
		if len(args) == 0 {
			env = helpers.GetEnvironment("")
		} else {
			env = helpers.GetEnvironment(args[0])
		}

		activeEnv := config.GetActiveEnv()

		if env == activeEnv.GetName() {
			fmt.Printf("%sEnvironment already set as active: %s%s\n", chalk.Yellow, activeEnv.GetName(), chalk.ResetColor)

			return
		}

		if activeEnv.IsRunning() {
			utils.Warn(fmt.Sprintf("Switching from running environment."))
		}

		viper.Set("active", env)

		config.Save(fmt.Sprintf("New active environment: %s%s%s (old: %s)\n", chalk.Green, env, chalk.ResetColor, activeEnv.GetName()), fmt.Errorf("Failed to set active environment"))
	},
}

func init() {
	EnvCommand.AddCommand(setCmd)
}
