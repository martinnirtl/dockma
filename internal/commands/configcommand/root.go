package configcommand

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// ConfigCommand is the top level dockma config command
var ConfigCommand = &cobra.Command{
	Use: "config",
}

var configHomeCmd = &cobra.Command{
	Use:     "home",
	Short:   "Print home dir of dockma config",
	Long:    `-`,
	Example: "dockma config home",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Dockma config home dir: %s%s%s\n", chalk.Cyan, viper.GetString("home"), chalk.ResetColor)
	},
}

var configCatCmd = &cobra.Command{
	Use:     "cat",
	Short:   "Concatinate config.json of dockma",
	Long:    `-`,
	Example: "dockma config cat",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

// TODO finish 'config load' cmd
var configLoadCmd = &cobra.Command{
	Use:     "load",
	Short:   "load dockma config.json",
	Long:    `-`,
	Example: "dockma config cat",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	ConfigCommand.AddCommand(configHomeCmd)
	ConfigCommand.AddCommand(configCatCmd)
	ConfigCommand.AddCommand(configLoadCmd)
}
