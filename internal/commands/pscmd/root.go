package pscmd

import (
	"os"

	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
)

// GetPSCommand returns the top level ps command
func GetPSCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "ps",
		Short:   "List running services of active environment",
		Long:    "List running services of active environment",
		Example: "dockma ps",
		Args:    cobra.NoArgs,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runPSCommand,
	}
}

func runPSCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	_, err = externalcommand.Execute("docker-compose ps", nil, "")
	utils.ErrorAndExit(err)
}
