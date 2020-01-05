package commands

import (
	"errors"
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getEnvironments() (envs []string) {
	envsMap := viper.GetStringMap("environments")

	envs = make([]string, 0, len(envsMap))

	for env := range envsMap {
		envs = append(envs, env)
	}

	sort.Strings(envs)

	return
}

func getEnvironment(env string) string {
	envs := getEnvironments()

	for _, envName := range envs {
		if env == envName {
			return env
		}
	}

	survey.AskOne(&survey.Select{
		Message: "Choose an environment:",
		Options: envs,
	}, &env)

	if env == "" {
		fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

		os.Exit(0)
	}

	return env
}

// GetEnvironmentsCommand returns the environments base command
func GetEnvironmentsCommand() *cobra.Command {
	environmentsCommand := &cobra.Command{
		Use: "environments",
	}

	environmentsCommand.AddCommand(getListCommand())
	environmentsCommand.AddCommand(getSetCommand())
	environmentsCommand.AddCommand(getRemoveCommand())

	return environmentsCommand
}

// TODO could be table with props from envs
func getListCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all configured environments",
		Long:  `-`,
		Run: func(cmd *cobra.Command, args []string) {
			envs := getEnvironments()

			activeEnv := viper.GetString("activeEnvironment")

			for _, env := range envs {
				if env == activeEnv {
					fmt.Printf("%s%s [active]%s\n", chalk.Cyan, env, chalk.ResetColor)
				} else {
					fmt.Println(env)
				}
			}
		},
	}
}

func getInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "init",
		Short: "Initialize environment",
		Long:  `-`,
		// TODO add flag to prevent setting active
		Args: cobra.NoArgs,
		Run:  func(cmd *cobra.Command, args []string) {},
	}
}

func getSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "set [environment]",
		Short: "Set active environment",
		Long:  `-`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("Too many arguments")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			env := ""
			if len(args) == 0 {
				env = getEnvironment("")
			} else {
				env = getEnvironment(args[0])
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
}

func getRemoveCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "remove [environment]",
		Short: "Remove environment",
		Long:  `-`,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) > 1 {
				return errors.New("Too many arguments")
			}

			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			env := ""
			if len(args) == 0 {
				env = getEnvironment("")
			} else {
				env = getEnvironment(args[0])
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
}
