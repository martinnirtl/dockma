package envcmd

import (
	"github.com/spf13/cobra"
)

// EnvCommand implements the top level envs command
var EnvCommand = &cobra.Command{
	Use:     "env",
	Aliases: []string{"envs"},
	Short:   "Environments reflect docker-compose based projects",
	Long:    "Environments reflect docker-compose based projects",
}
