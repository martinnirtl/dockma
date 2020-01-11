package commands

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

var environmentsCmd = &cobra.Command{
	Use:     "environments",
	Aliases: []string{"envs"},
}

// TODO could be table with props from envs
var environmentsListCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all configured environments",
	Long:    `-`,
	Example: "dockma envs list",
	Run: func(cmd *cobra.Command, args []string) {
		envs := utils.GetEnvironments()

		activeEnv := viper.GetString("activeEnvironment")

		if len(envs) > 0 {
			for _, env := range envs {
				if env == activeEnv {
					fmt.Printf("%s%s [active]%s\n", chalk.Cyan, env, chalk.ResetColor)
				} else {
					fmt.Println(env)
				}
			}
		} else {
			fmt.Printf("No environments configured. Add new environments by running %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)
		}
	},
}

var environmentsInitCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize environment",
	Long:    `-`,
	Example: "dockma envs init .",
	// TODO add flag to prevent setting active
	Args: cobra.NoArgs,
	Run:  func(cmd *cobra.Command, args []string) {},
}

var environmentsSetCmd = &cobra.Command{
	Use:     "set [environment]",
	Short:   "Set active environment",
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

		activeEnv := viper.GetString("activeEnvironment")

		if env == activeEnv {
			fmt.Printf("%sActive environment already set: %s%s\n", chalk.Yellow, activeEnv, chalk.ResetColor)

			return
		}

		fmt.Printf("%sNew active environment: %s%s (old: %s)\n", chalk.Green, env, chalk.ResetColor, activeEnv)

		viper.Set("activeEnvironment", env)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError setting active environment: %s%s\n", chalk.Red, env, chalk.ResetColor)
		}
	},
}

var environmentsRemoveCmd = &cobra.Command{
	Use:     "remove [environment]",
	Aliases: []string{"rm"},
	Short:   "Remove environment",
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

		activeEnv := viper.GetString("activeEnvironment")

		if env == activeEnv {
			fmt.Printf("%sRemoved active environment: %s%s\n\n", chalk.Yellow, env, chalk.ResetColor)

			viper.Set("activeEnvironment", nil)

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
	rootCmd.AddCommand(environmentsCmd)

	environmentsCmd.AddCommand(environmentsListCmd)
	environmentsCmd.AddCommand(environmentsInitCmd)
	environmentsCmd.AddCommand(environmentsSetCmd)
	environmentsCmd.AddCommand(environmentsRemoveCmd)
}
