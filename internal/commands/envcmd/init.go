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
	var envName string

	path := "."
	if len(args) == 1 {
		path = args[0]

		err := os.Chdir(path)
		if err != nil {
			fmt.Println(err)
			utils.ErrorAndExit(fmt.Errorf("Could not change to path %s", chalk.Underline.TextStyle(path)))
		}
	} else {
		utils.Warn("No path provided. Default is: .")
	}

	envName = survey.InputName("Enter a name for the new environment (has to be unique)", "")

	pull := "off"
	if _, err := os.Stat(".git"); !os.IsNotExist(err) {
		pull = survey.Select(fmt.Sprintf("Run %s before %s", chalk.Cyan.Color("git pull"), chalk.Cyan.Color("dockma up")), []string{"auto", "optional", "manual", "off"})
	} else {
		pull = "no-git"
	}

	workingDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		utils.ErrorAndExit(errors.New("Could get current working dir"))
	}

	proceed := survey.Confirm(fmt.Sprintf("Add new environment %s (location: %s)", chalk.Cyan.Color(envName), workingDir), true)
	if !proceed {
		utils.Abort()
	}

	viper.Set(fmt.Sprintf("envs.%s.home", envName), workingDir)
	viper.Set(fmt.Sprintf("envs.%s.pull", envName), pull)
	viper.Set(fmt.Sprintf("envs.%s.running", envName), false)

	config.Save(fmt.Sprintf("Initialized new environment: %s", chalk.Cyan.Color(envName)), fmt.Errorf("Failed to save newly created environment: %s", envName))

	activeEnv := config.GetActiveEnv()
	oldEnvName := activeEnv.GetName()

	set := true
	if activeEnv.IsRunning() {
		message := fmt.Sprintf("Currently active environment %s is %s. Set new environment %s active", chalk.Cyan.Color(oldEnvName), chalk.Green.Color("running"), chalk.Cyan.Color(envName))

		set = survey.Confirm(message, false)
	}

	if set {
		viper.Set("active", envName)

		config.Save(fmt.Sprintf("Set active environment: %s (old: %s)", chalk.Cyan.Color(envName), oldEnvName), fmt.Errorf("Failed to set active environment: %s", envName))
	}

}
