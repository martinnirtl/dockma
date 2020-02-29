package hooks

import (
	"fmt"
	"os"
	"os/user"
	"strings"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// RequiresInit ensures dockma has been initialized.
func RequiresInit(cmd *cobra.Command, args []string) {
	// enable printing of help
	if cmd.Name() == "help" {
		return
	}

	if init := viper.GetTime("init"); init.IsZero() {
		if user, err := user.Current(); err == nil {
			fmt.Printf("Come on, %s! Run %sdockma init%s first to initialize dockma.\n", strings.Title(user.Username), chalk.Cyan, chalk.ResetColor)
		} else {
			fmt.Printf("Please run %sdockma init%s first to initialize dockma.\n", chalk.Cyan, chalk.ResetColor)
		}

		os.Exit(0)
	}
}

// RequiresEnv ensures that any environment has been configured. Additionally it also executes RequiresInit pre-run-hook.
func RequiresEnv(cmd *cobra.Command, args []string) {
	RequiresInit(cmd, args)

	if len(config.GetEnvNames()) == 0 {
		PrintNoEnvConfigured()
	}
}

// RequiresActiveEnv ensures that an active environment is set. Additionally it also executes RequiresInit pre-run-hook.
func RequiresActiveEnv(cmd *cobra.Command, args []string) {
	RequiresInit(cmd, args)

	if config.GetActiveEnv().GetName() == "-" {
		PrintNoActiveEnvSet()
	}
}

// PrintNoEnvConfigured informs user that no env has been configured and exits with 0.
func PrintNoEnvConfigured() {
	fmt.Printf("No environment configured. Add new environment with %s.\n", chalk.Cyan.Color("dockma env init"))

	os.Exit(0)
}

// PrintNoActiveEnvSet informs user that no env is set and exits with 0.
func PrintNoActiveEnvSet() {
	fmt.Printf("No active environment set. Add new environment with %s or set active environment with %s.\n", chalk.Cyan.Color("dockma env init"), chalk.Cyan.Color("dockma env set"))

	os.Exit(0)
}
