package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var removeCmd = &cobra.Command{
	Use:     "remove [environment]",
	Aliases: []string{"rm"},
	Short:   "Remove environment.",
	Long:    "Remove environment.",
	Args:    cobra.RangeArgs(0, 1),
	Example: "dockma envs remove my-env",
	Run: func(cmd *cobra.Command, args []string) {
		env := ""
		if len(args) == 0 {
			env = utils.GetEnvironment("")
		} else {
			env = utils.GetEnvironment(args[0])
		}

		sure := survey.Confirm(fmt.Sprintf("Are you sure to remove '%s'", env), false)

		if !sure {
			utils.Abort()
		}

		activeEnv := viper.GetString("active")

		if env == activeEnv {
			fmt.Printf("%sRemoved active environment: %s%s\n\n", chalk.Yellow, env, chalk.ResetColor)

			viper.Set("active", "-")

			fmt.Printf("%sUnset active environment%s\n", chalk.Cyan, chalk.ResetColor)
		} else {
			fmt.Printf("%sRemoved environment: %s%s\n", chalk.Cyan, env, chalk.ResetColor)
		}

		envs := viper.GetStringMap("envs")

		delete(envs, env)

		viper.Set("envs", envs)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError removing environment: %s%s\n", chalk.Red, env, chalk.ResetColor)
		}
	},
}

func init() {
	EnvCommand.AddCommand(removeCmd)
}