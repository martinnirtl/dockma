package argsvalidator

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// OnlyEnvs checks args for valid envs
func OnlyEnvs(cmd *cobra.Command, args []string) error {
	for _, arg := range args {
		if !utils.Includes(config.GetEnvNames(), arg) {
			return fmt.Errorf("No such environment %s", chalk.Underline.TextStyle(arg))
		}
	}

	return nil
}

// OptionalEnv checks for one optional env
func OptionalEnv(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("Expected 0 or 1 argument. Got %d", len(args))
	}

	for _, arg := range args {
		if !utils.Includes(config.GetEnvNames(), arg) {
			return fmt.Errorf("No such environment %s", chalk.Underline.TextStyle(arg))
		}
	}

	return nil
}

// OnlyProfiles checks args for valid profiles of active environment
func OnlyProfiles(cmd *cobra.Command, args []string) error {
	activeEnv := config.GetActiveEnv()

	for _, arg := range args {
		if !utils.Includes(activeEnv.GetProfileNames(), arg) {
			return fmt.Errorf("No such environment %s", chalk.Underline.TextStyle(arg))
		}
	}

	return nil
}

// OptionalProfile checks for one optional profile of active environment
func OptionalProfile(cmd *cobra.Command, args []string) error {
	if len(args) > 1 {
		return fmt.Errorf("Expected 0 or 1 argument. Got %d", len(args))
	}

	activeEnv := config.GetActiveEnv()

	for _, arg := range args {
		if !utils.Includes(activeEnv.GetProfileNames(), arg) {
			return fmt.Errorf("No such environment %s", chalk.Underline.TextStyle(arg))
		}
	}

	return nil
}

// OnlyServices checks args for valid services of active environment
func OnlyServices(cmd *cobra.Command, args []string) error {
	activeEnv := config.GetActiveEnv()

	for _, arg := range args {
		if !utils.Includes(helpers.GetEnvServices(activeEnv), arg) {
			return fmt.Errorf("No such environment %s", chalk.Underline.TextStyle(arg))
		}
	}

	return nil
}
