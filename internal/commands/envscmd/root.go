package envscmd

import (
	"github.com/spf13/cobra"
)

// EnvsCommand implements the top level envs command
var EnvsCommand = &cobra.Command{
	Use:     "env",
	Aliases: []string{"environment"},
	Short:   "Environments reflect docker-compose based projects.",
	Long:    "Environments reflect docker-compose based projects.",
}
