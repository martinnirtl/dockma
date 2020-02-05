package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// Println works as custom fmt.Printf with chalk.ResetColor and \n automattically attached
func Println(text string) {
	fmt.Printf("%s%s\n", text, chalk.ResetColor)
}

// Abort prints 'Aborted.' to std out and exits process with 0
func Abort() {
	fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// Error prints 'Error: <message>' to std out and exits process with 0
func Error(err error) {
	fmt.Printf("%sError: %s%s\n", chalk.Red, err, chalk.ResetColor)

	os.Exit(0)
}

// NoEnvs prints no envs configured and exits
func NoEnvs() {
	fmt.Printf("No environments configured. Add new environment with %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// GetFullLogfilePath returns the absolute path of the logfile location in dockma home dir
func GetFullLogfilePath(filename string) string {
	path := viper.GetString("home")

	return filepath.Join(path, filename)
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

// Fallback returns fallback if val is nil
func Fallback(val string, fallback string) string {
	if val == "" {
		return fallback
	}

	return val
}
