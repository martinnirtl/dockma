package envcmd

import (
	"github.com/spf13/cobra"
)

// EnvCommand implements the top level envs command
var EnvCommand = &cobra.Command{
	Use:     "env",
	Aliases: []string{"environment"},
	Short:   "Environments reflect docker-compose based projects.",
	Long:    "Environments reflect docker-compose based projects.",
}
