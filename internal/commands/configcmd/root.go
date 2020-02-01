package configcmd

import (
	"github.com/spf13/cobra"
)

// ConfigCommand is the top level config command
var ConfigCommand = &cobra.Command{
	Use: "config",
}
