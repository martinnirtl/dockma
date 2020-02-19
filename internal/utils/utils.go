package utils

import (
	"fmt"
	"os"

	"github.com/ttacon/chalk"
)

// Abort prints 'Aborted.' to std out and exits process with 0
func Abort() {
	fmt.Printf("%sAborted.%s\n", chalk.Cyan, chalk.ResetColor)

	os.Exit(0)
}

// Success prints green colored text to std out
func Success(text string) {
	fmt.Println(chalk.Green.Color(text))
}

// Warn prints yellow colored text to std out
func Warn(text string) {
	fmt.Println(chalk.Yellow.Color(text))
}

// Error checks if err is not nil, prints 'Error: <message>' to std out
func Error(err error) {
	if err != nil {
		fmt.Println(chalk.Red.Color(fmt.Sprintf("Error: %s", err)))
	}
}

// ErrorAndExit checks if err is not nil, prints 'Error: <message>' to std out and exits process with 0
func ErrorAndExit(err error) {
	if err != nil {
		fmt.Println(chalk.Red.Color(fmt.Sprintf("Error: %s", err)))

		os.Exit(0)
	}
}

// PrintNoActiveEnvSet prints no envs configured and exits
func PrintNoActiveEnvSet() {
	fmt.Printf("No active environment configured. Add new environment with %s or set active environment with %s.\n", chalk.Cyan.Color("dockma env init"), chalk.Cyan.Color("dockma env set"))

	os.Exit(0)
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
