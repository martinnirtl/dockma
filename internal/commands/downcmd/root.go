package downcmd

import (
	"fmt"
	"os"

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
		Use:     "down",
		Short:   "Stops active environment",
		Long:    "Stops active environment",
		Example: "dockma down",
		Args:    cobra.NoArgs,
		Run:     runDownCommand,
	}
}

func runDownCommand(cmd *cobra.Command, args []string) {
	filepath := config.GetSubcommandLogfile()

	activeEnv := config.GetActiveEnv()

	if activeEnv.GetName() == "-" {
		utils.PrintNoActiveEnvSet()
	}

	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	var timebridger externalcommand.Timebridger
	if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
		timebridger = spinnertimebridger.Default(fmt.Sprintf("Running %s", chalk.Cyan.Color("docker-compose down")))
	}

	output, err := externalcommand.Execute("docker-compose down", timebridger, filepath)

	utils.Error(err)
	if err != nil {
		fmt.Print(string(output))

		os.Exit(0)
	}

	utils.Success("Executed command: docker-compose down")

	viper.Set(fmt.Sprintf("envs.%s.running", activeEnv.GetName()), false)
	config.Save("", fmt.Errorf("Failed to set running to 'false' [%s]", activeEnv))
}
