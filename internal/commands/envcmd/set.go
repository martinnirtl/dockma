package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/internal/utils/helpers"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getSetCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "set [environment]",
		Short:   "Set active environment",
		Long:    "Set active environment",
		Example: "dockma envs set",
		// FIXME loading of viper config for dynamic ValidArgs
		Args: cobra.RangeArgs(0, 1),
		Run:  runSetCommand,
	}
}

func runSetCommand(cmd *cobra.Command, args []string) {
	env := ""
	if len(args) == 0 {
		env = helpers.GetEnvironment("")
	} else {
		env = helpers.GetEnvironment(args[0])
	}

	activeEnv := config.GetActiveEnv()

	if env == activeEnv.GetName() {
		fmt.Printf("Environment already set active: %s\n", chalk.Cyan.Color(activeEnv.GetName()))

		return
	}

	if activeEnv.IsRunning() {
		utils.Warn(fmt.Sprintf("Switching from running environment."))
	}

	viper.Set("active", env)

	config.Save(fmt.Sprintf("New active environment: %s (old: %s)", chalk.Cyan.Color(env), activeEnv.GetName()), fmt.Errorf("Failed to set active environment"))

}
