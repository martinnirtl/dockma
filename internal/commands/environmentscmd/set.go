package environmentscmd

import (
	"errors"
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var setCmd = &cobra.Command{
	Use:     "set [environment]",
	Short:   "Set active environment.",
	Long:    `-`,
	Example: "dockma envs set",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("Too many arguments")
		}

		return nil
	},
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
	EnvironmentsCommand.AddCommand(setCmd)
}
