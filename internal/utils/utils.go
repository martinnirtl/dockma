package utils

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/ttacon/chalk"
)

// Println works as custom fmt.Printf with chalk.ResetColor and \n automattically attached
func Println(text string) {
	fmt.Printf("%s%s\n", text, chalk.ResetColor)
}

// PrintCyan colors text in cyan
func PrintCyan(text string) string {
	return fmt.Sprintf("%s%s%s", chalk.Cyan, text, chalk.ResetColor)
}

// Abort prints 'Aborted.' to std out and exits process with 0
func Abort() {
	fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// Success prints green colored text to std out and exits process with 0
func Success(text string) {
	fmt.Printf("%s%s%s\n", chalk.Green, text, chalk.ResetColor)

	os.Exit(0)
}

// Error checks if err is not nil, prints 'Error: <message>' to std out and exits process with 0
func Error(err error) {
	if err != nil {
		fmt.Printf("%sError: %s%s\n", chalk.Red, err, chalk.ResetColor)

		os.Exit(0)
	}
}

// NoEnvs prints no envs configured and exits
func NoEnvs() {
	fmt.Printf("No environments configured. Add new environment with %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// GetEnvironment returns one environment
func GetEnvironment(env string) string {
	envs := config.GetEnvNames()

	for _, envName := range envs {
		if env == envName {
			return env
		}
	}

	fmt.Printf("%sNo such environment: %s%s\n", chalk.Yellow, env, chalk.ResetColor)

	env = survey.Select("Choose an environment", envs)

	return env
}

// Fallback returns fallback if val is nil
func Fallback(val string, fallback string) string {
	if val == "" {
		return fallback
	}

	return val
}

// Includes checks if string slice includes string
func Includes(slice []string, s string) bool {
	for _, val := range slice {
		if val == s {
			return true
		}
	}

	return false
}
