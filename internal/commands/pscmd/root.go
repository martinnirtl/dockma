package pscmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// PSCommand implements the top level ps command
var PSCommand = &cobra.Command{
	Use:     "ps",
	Short:   "List running services of active environment",
	Long:    "List running services of active environment",
	Example: "dockma ps",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.PrintNoActiveEnvSet()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		utils.ErrorAndExit(err)

		_, err = externalcommand.Execute("docker-compose ps", nil, "")

		utils.ErrorAndExit(err)
	},
}
