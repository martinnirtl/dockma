package logscmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
)

var followFlag bool
var timestampsFlag bool
var tailFlag int

// LogsCommand implements the top level logs command
var LogsCommand = &cobra.Command{
	Use:     "logs [service...]",
	Short:   "Logs output of all or only selected services",
	Long:    "Logs output of all or only selected services",
	Example: "dockma logs -f my-service",
	Args:    cobra.ArbitraryArgs,
	// Args:      cobra.OnlyValidArgs, // TODO investigate
	// ValidArgs: getValidArgs(),
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()

		if activeEnv.GetName() == "-" {
			utils.PrintNoActiveEnvSet()
		}

		envHomeDir := activeEnv.GetHomeDir()

		err := os.Chdir(envHomeDir)
		utils.ErrorAndExit(err)

		args = addFlagsToArgs(args)
		command := externalcommand.JoinCommand("docker-compose logs", args...)

		_, err = externalcommand.Execute(command, nil, "")
		utils.ErrorAndExit(err)
	},
}

func init() {
	LogsCommand.Flags().BoolVarP(&followFlag, "follow", "f", false, "follow log output")
	LogsCommand.Flags().BoolVarP(&timestampsFlag, "timestamps", "t", false, "show timestamps")
	LogsCommand.Flags().IntVar(&tailFlag, "tail", 0, "number of lines to show from the end of the logs for each service")
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

// func getValidArgs() []string {
// 	activeEnv := config.GetActiveEnv()

// 	if activeEnv == "-" {
// 		return []string{}
// 	}

// 	envHomeDir := config.GetEnvHomeDir(activeEnv)

// 	services, err := dockercompose.GetServices(envHomeDir)
// 	utils.ErrorAndExit(err)

// 	return services.All
// }
