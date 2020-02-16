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

	fmt.Printf("%sNo such environment: %s%s\n", chalk.Yellow, env, chalk.ResetColor)

	env = survey.Select("Choose an environment", envs)

	return env
}
