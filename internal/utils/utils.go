package utils

import (
	"fmt"
	"os"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// Abort prints 'Aborted.' to std out and exits process with 0
func Abort() {
	fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// GetEnvironments returns configured environments
func GetEnvironments() (envs []string) {
	envsMap := viper.GetStringMap("environments")

	envs = make([]string, 0, len(envsMap))

	for env := range envsMap {
		envs = append(envs, env)
	}

	sort.Strings(envs)

	return
}

// GetEnvironment returns one environment
func GetEnvironment(env string) string {
	envs := GetEnvironments()

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
