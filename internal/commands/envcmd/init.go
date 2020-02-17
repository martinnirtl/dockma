package envcmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var initCmd = &cobra.Command{
	Use:     "init [path-to-environment]",
	Short:   "Initialize new environment",
	Long:    "Initialize new environment",
	Example: "dockma env init path/to/env",
	Args:    cobra.RangeArgs(0, 1),

	Run: func(cmd *cobra.Command, args []string) {
		var env string

		path := "."
		if len(args) == 1 {
			path = args[0]

			err := os.Chdir(path)

			if err != nil {
				utils.ErrorAndExit(fmt.Errorf("Could not find directory: %s", path))
			}
		}

		env = survey.InputName("Enter a name for the new environment (has to be unique)", "")

		if env == "" {
			utils.ErrorAndExit(errors.New("Got empty string for environment name"))
		} else if env == "-" {
			utils.ErrorAndExit(errors.New("Invalid environment name '-'"))
		}

		workingDir, err := os.Getwd()

		if err != nil {
			utils.ErrorAndExit(errors.New("Could not read current working directory"))
		}

		pull := "off"
		if _, err := os.Stat(".git"); !os.IsNotExist(err) {
			pull = survey.Select(fmt.Sprintf("Run %sgit pull%s before %sdockma up%s", chalk.Cyan, chalk.Reset, chalk.Cyan, chalk.ResetColor), []string{"auto", "optional", "manual", "off"})
		} else {
			pull = "no-git"
		}

		proceed := survey.Confirm(fmt.Sprintf("Add new environment %s%s%s (location: %s)", chalk.Cyan, env, chalk.ResetColor, workingDir), true)

		if !proceed {
			utils.Abort()
		}

		viper.Set(fmt.Sprintf("envs.%s.home", env), workingDir)
		viper.Set(fmt.Sprintf("envs.%s.pull", env), pull)
		viper.Set(fmt.Sprintf("envs.%s.running", env), false)

		activeEnv := config.GetActiveEnv()
		oldEnv := activeEnv.GetName()

		var set bool
		if activeEnv.IsRunning() {
			message := fmt.Sprintf("Current active environment running: %s. Set newly initialized environment active", oldEnv)

			set = survey.Confirm(message, false)
		}

		if set {
			viper.Set("active", env)
		}

		config.Save(fmt.Sprintf("Initialized new environment: %s%s%s\n", chalk.Cyan, env, chalk.ResetColor), fmt.Errorf("Failed to save newly created environment"))
	},
}

func init() {
	EnvCommand.AddCommand(initCmd)
}
