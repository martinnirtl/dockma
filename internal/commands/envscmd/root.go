package envscmd

import (
	"github.com/spf13/cobra"
)

// EnvsCommand is the top level Envs command
var EnvsCommand = &cobra.Command{
	Use:   "envs",
	Short: "Environments reflect docker-compose based projects.",
	Long:  "-",
}
