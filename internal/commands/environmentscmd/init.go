package environmentscmd

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

var initCmd = &cobra.Command{
	Use:     "init [path-to-environment]",
	Short:   "Initialize new environment.",
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
			utils.Abort()
		}

		workingDir, err := os.Getwd()

		if err != nil {
			fmt.Printf("%sError. Could not read current working directory%s\n", chalk.Red, chalk.ResetColor)

			os.Exit(0)
		}

		// TODO read docker-compose.yaml
		services, err := dockercompose.GetServices(workingDir)

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
		viper.Set(fmt.Sprintf("environments.%s.services", env), services.All)

		oldEnv := viper.GetString("active")

		viper.Set("active", env)

		if err := viper.WriteConfig(); err != nil {
			fmt.Printf("%sError on initializing environment: %s (old: %s)%s\n", chalk.Red, env, oldEnv, chalk.ResetColor)

			os.Exit(1)
		}

		fmt.Printf("%sSet active environment: %s%s\n", chalk.Cyan, env, chalk.ResetColor)
	},
}

func init() {
	EnvironmentsCommand.AddCommand(initCmd)
}
