package logscmd

import (
	"fmt"
	"os"
	"sort"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var followFlag bool
var timestampsFlag bool
var tailFlag int

var LogsCommand = &cobra.Command{
	Use:   "logs [service...]",
	Short: "Logs output of either all or only selected services.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		err := os.Chdir(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		args = addFlagsToArgs(args)

		command := externalcommand.JoinCommandSlices("docker-compose logs", args...)

		_, err = externalcommand.Execute(command, nil, "")

		if err != nil {
			utils.Error(err)
		}
	},
}

func init() {
	LogsCommand.Flags().BoolVarP(&followFlag, "follow", "f", false, "Follow log output.")
	LogsCommand.Flags().BoolVarP(&timestampsFlag, "timestamps", "t", false, "Show timestamps.")
	LogsCommand.Flags().IntVar(&tailFlag, "tail", 0, "Number of lines to show from the end of the logs for each service.")
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
