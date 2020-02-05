package environmentscmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var removeCmd = &cobra.Command{
	Use:     "remove [environment]",
	Aliases: []string{"rm"},
	Short:   "Remove environment.",
	Long:    `-`,
	Example: "dockma envs remove my-env",
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

		sure := false
		if err := survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Are you sure to remove '%s'", env),
		}, &sure); err != nil || !sure {
			fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

			os.Exit(0)
		}

		activeEnv := viper.GetString("active")

		if env == activeEnv {
			fmt.Printf("%sRemoved active environment: %s%s\n\n", chalk.Yellow, env, chalk.ResetColor)

			viper.Set("active", "-")

			fmt.Printf("%sUnset active environment%s\n", chalk.Cyan, chalk.ResetColor)
		} else {
			fmt.Printf("%sRemoved environment: %s%s\n", chalk.Cyan, env, chalk.ResetColor)
		}

		envs := viper.GetStringMap("environments")

		delete(envs, env)

		viper.Set("environments", envs)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError removing environment: %s%s\n", chalk.Red, env, chalk.ResetColor)
		}
	},
}

func init() {
	EnvironmentsCommand.AddCommand(removeCmd)
}
