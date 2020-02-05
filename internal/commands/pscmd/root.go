package pscmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var PSCommand = &cobra.Command{
	Use:   "ps",
	Short: "List running services of active environment.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("environments.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		_, err = externalcommand.Execute("docker-compose ps", nil, "")

		if err != nil {
			utils.Error(err)
		}
	},
}
