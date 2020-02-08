package envscmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var initCmd = &cobra.Command{
	Use:     "init [path-to-environment]",
	Short:   "Initialize new environment.",
	Long:    "Initialize new environment.",
	Args:    cobra.RangeArgs(0, 1),
	Example: "dockma envs init path/to/env",
	Run: func(cmd *cobra.Command, args []string) {
		var env string

		path := "."
		if len(args) == 1 {
			path = args[0]

			if err := os.Chdir(path); err != nil {
				fmt.Printf("%sError: Could not find directory: %s%s\n", chalk.Red, path, chalk.ResetColor)

				os.Exit(0)
			}
		}

		env = survey.Input("Enter a name for the new environment (has to be unique)", "")

		if env == "" {
			utils.Error(errors.New("Got empty string for environment name"))
		} else if env == "-" {
			utils.Error(errors.New("Invalid environment name '-'"))
		}

		workingDir, err := os.Getwd()

		if err != nil {
			fmt.Printf("%sError. Could not read current working directory%s\n", chalk.Red, chalk.ResetColor)

			os.Exit(0)
		}

		autoPull := survey.Confirm(fmt.Sprintf("Run %sgit pull%s before %sdockma up%s", chalk.Cyan, chalk.Reset, chalk.Cyan, chalk.ResetColor), false)

		proceed := survey.Confirm(fmt.Sprintf("Add new environment %s%s%s (location: %s)", chalk.Cyan, env, chalk.ResetColor, workingDir), true)

		if !proceed {
			utils.Abort()
		}

		viper.Set(fmt.Sprintf("envs.%s.home", env), workingDir)
		viper.Set(fmt.Sprintf("envs.%s.autopull", env), autoPull)

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
	EnvsCommand.AddCommand(initCmd)
}
