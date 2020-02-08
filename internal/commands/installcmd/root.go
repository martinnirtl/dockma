package installcmd

import (
	"github.com/spf13/cobra"
)

// InstallCommand implements the top level install command
var InstallCommand = &cobra.Command{
	Use:   "install",
	Short: "-",
	Long:  "-",
	Run:   func(cmd *cobra.Command, args []string) {},
}
