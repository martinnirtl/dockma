package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
)

func getHomeCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "home [environment]",
		Short:   "Get environment home dir",
		Long:    "Get environment home dir",
		Example: "dockma env home",
		Args:    argvalidators.OptionalEnv,
		PreRun:  hooks.RequiresEnv,
		Run:     runHomeCommand,
	}
}

func runHomeCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()

	envName := ""
	if len(args) == 0 {
		if activeEnv.GetName() == "-" {
			envNames := config.GetEnvNames()
			envName = survey.Select("Choose an environment", envNames)
		} else {
			envName = activeEnv.GetName()
		}
	} else {
		envName = args[0]
	}

	env, err := config.GetEnv(envName)
	utils.ErrorAndExit(err)

	fmt.Println(env.GetHomeDir())
}
