package helpers

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/ttacon/chalk"
)

// GetEnvironment returns one environment
func GetEnvironment(env string) string {
	envs := config.GetEnvNames()

	for _, envName := range envs {
		if env == envName {
			return env
		}
	}

	if env != "" {
		fmt.Printf("%sNo such environment: %s%s\n", chalk.Yellow, env, chalk.ResetColor)
	}

	env = survey.Select("Choose an environment", envs)

	return env
}

// PrintErrorList prints error list as red 'suberrors' to stdout (To be used with config.SaveNow).
func PrintErrorList(errorList []error) {
	for _, err := range errorList {
		fmt.Printf("%s> %s%s\n", chalk.Red, err, chalk.ResetColor)
	}
}

// PrintMessageList prints messages to stdout (To be used with config.SaveNow)
func PrintMessageList(messageList []string) {
	for _, msg := range messageList {
		fmt.Printf("%s\n", msg)
	}
}
