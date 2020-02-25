package logscmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
)

var followFlag bool
var timestampsFlag bool
var tailFlag int

// GetLogsCommand returns the top level logs command
func GetLogsCommand() *cobra.Command {
	logsCommand := &cobra.Command{
		Use:     "logs [service...]",
		Short:   "Logs output of all or only selected services",
		Long:    "Logs output of all or only selected services",
		Example: "dockma logs -f database",
		Args:    argvalidators.OnlyServices,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runLogsCommand,
	}

	logsCommand.Flags().BoolVarP(&followFlag, "follow", "f", false, "follow log output")
	logsCommand.Flags().BoolVarP(&timestampsFlag, "timestamps", "t", false, "show timestamps")
	logsCommand.Flags().IntVar(&tailFlag, "tail", 0, "number of lines to show from the end of the logs for each service")

	return logsCommand
}

func runLogsCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()

	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	args = addFlagsToArgs(args)
	command := externalcommand.JoinCommand("docker-compose logs", args...)

	_, err = externalcommand.Execute(command, nil, "")
	utils.ErrorAndExit(err)
}

func addFlagsToArgs(args []string) []string {
	if followFlag {
		args = append(args, "--follow")
	}

	if timestampsFlag {
		args = append(args, "--timestamps")
	}

	if tailFlag > 0 {
		args = append(args, fmt.Sprintf("--tail=%d", tailFlag))
	}

	if followFlag || timestampsFlag || tailFlag > 0 {
		sort.Strings(args)
	}

	return args
}
