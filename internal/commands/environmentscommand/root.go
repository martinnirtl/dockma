package environmentscommand

import (
	"errors"
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var EnvironmentsCommand = &cobra.Command{
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
			fmt.Printf("No environments configured. Add a new environment by running %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)
		}
	},
}

var environmentsInitCmd = &cobra.Command{
	Use:     "init [path-to-environment]",
	Short:   "Initialize environment",
	Long:    `-`,
	Example: "dockma envs init path/to/env",
	// TODO add flag to prevent setting active
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 1 {
			return errors.New("Too many arguments")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		var env string

		path := "."
		if len(args) == 1 {
			path = args[0]

			if err := os.Chdir(path); err != nil {
				fmt.Printf("%sError. Could not change directory to: %s%s\n", chalk.Red, path, chalk.ResetColor)

				os.Exit(0)
			}
		}

		err := survey.AskOne(&survey.Input{
			Message: "Enter a name for the new environment (has to be unique)",
		}, &env)

		if err != nil {
			fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

			os.Exit(0)
		}

		workingDir, err := os.Getwd()

		if err != nil {
			fmt.Printf("%sError. Could not read current working directory%s\n", chalk.Red, chalk.ResetColor)

			os.Exit(0)
		}

		// TODO read docker-compose.yaml
		services, err := dockercompose.ReadServices(workingDir)

		if err != nil {
			fmt.Print(err)

			os.Exit(1)
		}

		proceed := false
		err = survey.AskOne(&survey.Confirm{
			Message: fmt.Sprintf("Add new environment %s%s%s (location: %s)", chalk.Cyan, env, chalk.ResetColor, workingDir),
			Default: true,
		}, &proceed)

		if !proceed {
			utils.Abort()
		} else if err != nil {
			fmt.Printf("%sError. %s%s\n", chalk.Red, err, chalk.ResetColor)

			os.Exit(0)
		}

		viper.Set(fmt.Sprintf("environments.%s.home", env), workingDir)
		viper.Set(fmt.Sprintf("environments.%s.services", env), services)

		oldEnv := viper.GetString("activeEnvironment")

		viper.Set("activeEnvironment", env)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError on initializing environment: %s (old: %s)%s\n", chalk.Red, env, oldEnv, chalk.ResetColor)

			os.Exit(1)
		}

		fmt.Printf("%sSet active environment: %s%s\n", chalk.Cyan, env, chalk.ResetColor)
	},
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
	EnvironmentsCommand.AddCommand(environmentsListCmd)
	EnvironmentsCommand.AddCommand(environmentsInitCmd)
	EnvironmentsCommand.AddCommand(environmentsSetCmd)
	EnvironmentsCommand.AddCommand(environmentsRemoveCmd)
}
