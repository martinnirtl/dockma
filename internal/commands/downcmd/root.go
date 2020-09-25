package downcmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// GetDownCommand returns the top level down command
func GetDownCommand() *cobra.Command {
	return &cobra.Command{
		Use:              "down",
		Short:            "Downs active environment",
		Long:             "Downs active environment",
		Example:          "dockma down",
		Args:             cobra.NoArgs,
		PersistentPreRun: hooks.RequiresActiveEnv,
		Run:              runDownCommand,
	}
}

func runDownCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	var timebridger externalcommand.Timebridger
	if hideCmdOutput := viper.GetBool("hidecommandoutput"); hideCmdOutput {
		timebridger = spinnertimebridger.Default(fmt.Sprintf("Running %s", chalk.Cyan.Color("docker-compose down")))
	}

	logfile := config.GetSubcommandLogfile()

	output, err := externalcommand.Execute("docker-compose down", timebridger, logfile)
	if err != nil && timebridger != nil {
		fmt.Print(string(output))
	}
	utils.ErrorAndExit(err)

	utils.Success(fmt.Sprintf("Executed command: %s", chalk.Cyan.Color("docker-compose down")))

	viper.Set(fmt.Sprintf("envs.%s.running", activeEnv.GetName()), false)
	config.Save("", fmt.Errorf("Failed to set environment %s to %s", chalk.Underline.TextStyle(activeEnv.GetName()), chalk.Underline.TextStyle("not running")))
}
