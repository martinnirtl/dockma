package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var setCmd = &cobra.Command{
	Use:     "set [environment]",
	Short:   "Set active environment.",
	Long:    "Set active environment.",
	Args:    cobra.RangeArgs(0, 1),
	Example: "dockma envs set",
	Run: func(cmd *cobra.Command, args []string) {
		env := ""
		if len(args) == 0 {
			env = utils.GetEnvironment("")
		} else {
			env = utils.GetEnvironment(args[0])
		}

		activeEnv := viper.GetString("active")

		if env == activeEnv {
			fmt.Printf("%sActive environment already set: %s%s\n", chalk.Yellow, activeEnv, chalk.ResetColor)

			return
		}

		fmt.Printf("%sNew active environment: %s%s (old: %s)\n", chalk.Green, env, chalk.ResetColor, activeEnv)

		viper.Set("active", env)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError setting active environment: %s%s\n", chalk.Red, env, chalk.ResetColor)
		}
	},
}

func init() {
	EnvCommand.AddCommand(setCmd)
}
