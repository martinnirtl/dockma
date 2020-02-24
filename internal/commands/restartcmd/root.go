package restartcmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/commands/argsvalidator"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// GetRestartCommand returns the top level restart command
func GetRestartCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "restart [service...]",
		Short:   "Restart all or only selected services",
		Long:    "Restart all or only selected services",
		Example: "dockma restart database",
		Args:    argsvalidator.OnlyServices,
		Run:     runRestartCommand,
	}
}

func runRestartCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()

	if activeEnv.GetName() == "-" {
		utils.PrintNoActiveEnvSet()
	}

	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	logfile := config.GetSubcommandLogfile()

	command := externalcommand.JoinCommand("docker-compose restart", args...)

	if len(args) == 0 {
		restartAll := survey.Confirm("Restart all services", true)

		if !restartAll {
			// services, err := dockercompose.GetServices(envHomeDir)
			profile, err := config.GetActiveEnv().GetLatest()
			utils.ErrorAndExit(err)

			selected := survey.MultiSelect("Select services to restart", profile.Services, profile.Selected)

			command = externalcommand.JoinCommand("docker-compose restart", selected...)
		}
	}

	var timebridger externalcommand.Timebridger
	if config.GetHideSubcommandOutputSetting() {
		timebridger = spinnertimebridger.Default(fmt.Sprintf("Running %s", chalk.Cyan.Color("docker-compose restart")))
	}

	output, err := externalcommand.Execute(command, timebridger, logfile)
	if err != nil && timebridger != nil {
		fmt.Print(output)
	}
	utils.ErrorAndExit(err)

	utils.Success(fmt.Sprintf("Executed command: %s", command))
}
