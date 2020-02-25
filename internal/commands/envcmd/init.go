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

func getInitCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "init [path-to-environment]",
		Short:   "Initialize new environment",
		Long:    "Initialize new environment",
		Example: "dockma env init path/to/env",
		Args:    cobra.RangeArgs(0, 1),
		Run:     runInitCommand,
	}
}

func runInitCommand(cmd *cobra.Command, args []string) {
	var env string

	path := "."
	if len(args) == 1 {
		path = args[0]

		err := os.Chdir(path)

		if err != nil {
			utils.ErrorAndExit(fmt.Errorf("Could not find directory: %s", path))
		}
	} else {
		utils.Warn("No path provided. Default is: .")
	}

	env = survey.InputName("Enter a name for the new environment (has to be unique)", "")

	workingDir, err := os.Getwd()
	if err != nil {
		utils.ErrorAndExit(errors.New("Could not read current working directory"))
	}

	pull := "off"
	if _, err := os.Stat(".git"); !os.IsNotExist(err) {
		pull = survey.Select(fmt.Sprintf("Run %s before %s", chalk.Cyan.Color("git pull"), chalk.Cyan.Color("dockma up")), []string{"auto", "optional", "manual", "off"})
	} else {
		pull = "no-git"
	}

	proceed := survey.Confirm(fmt.Sprintf("Add new environment %s (location: %s)", chalk.Cyan.Color(env), workingDir), true)
	if !proceed {
		utils.Abort()
	}

	viper.Set(fmt.Sprintf("envs.%s.home", env), workingDir)
	viper.Set(fmt.Sprintf("envs.%s.pull", env), pull)
	viper.Set(fmt.Sprintf("envs.%s.running", env), false)

	config.Save(fmt.Sprintf("Initialized new environment: %s", chalk.Cyan.Color(env)), fmt.Errorf("Failed to save newly created environment"))

	activeEnv := config.GetActiveEnv()
	oldEnv := activeEnv.GetName()

	set := true
	if activeEnv.IsRunning() {
		message := fmt.Sprintf("Current active environment running: %s. Set newly initialized environment active", oldEnv)

		set = survey.Confirm(message, false)
	}

	if set {
		viper.Set("active", env)

		config.Save(fmt.Sprintf("Set active environment: %s", chalk.Cyan.Color(env)), fmt.Errorf("Failed to set active environment"))
	}

}
