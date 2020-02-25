package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var pathFlag bool

// TODO could be table with props from envs
func getListCommand() *cobra.Command {
	listCommand := &cobra.Command{
		Use:     "list",
		Short:   "List all configured environments",
		Long:    "List all configured environments",
		Example: "dockma envs list",
		Args:    cobra.NoArgs,
		PreRun:  hooks.RequiresEnv,
		Run:     runListCommand,
	}

	listCommand.Flags().BoolVarP(&pathFlag, "path", "p", false, "print path")

	return listCommand
}

func runListCommand(cmd *cobra.Command, args []string) {
	envs := config.GetEnvNames()
	maxEnvNameLength := getLongestWordLength(envs, 3)

	activeEnv := config.GetActiveEnv()
	activeEnvName := activeEnv.GetName()

	if activeEnvName == "-" {
		hooks.PrintNoActiveEnvSet()
	}

	for _, envName := range envs {
		if envName == activeEnvName {
			fmt.Printf("%s ", chalk.Cyan.Color("[active]"))
		} else {
			fmt.Print("         ")
		}

		fmt.Print(chalk.Bold.TextStyle(pad(envName, maxEnvNameLength)))

		if env, err := config.GetEnv(envName); err == nil {
			if env.IsRunning() {
				fmt.Printf(" %s", chalk.Green.Color("running"))
			} else {
				fmt.Print(" -------")
			}

			if pathFlag {
				fmt.Print(" " + env.GetHomeDir())
			}
		}

		fmt.Println()
	}
}

func getLongestWordLength(words []string, minLength int) int {
	length := minLength
	for _, word := range words {
		if len(word) > length {
			length = len(word)
		}
	}

	return length
}

func pad(word string, n int) string {
	rest := n - len(word)

	for i := 0; i < rest; i++ {
		word = word + " "
	}

	return word
}
