package environmentscmd

import (
	"github.com/spf13/cobra"
)

// EnvironmentsCommand is the top level environments command
var EnvironmentsCommand = &cobra.Command{
	Use:   "envs",
	Short: "Environments reflect docker-compose based projects.",
	Long:  "-",
}
