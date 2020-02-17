package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var pathFlag bool

// TODO could be table with props from envs
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all configured environments.",
	Long:    "List all configured environments.",
	Args:    cobra.NoArgs,
	Example: "dockma env list",
	Run: func(cmd *cobra.Command, args []string) {
		envs := config.GetEnvNames()
		maxEnvNameLength := getLongestWordLength(envs, 3)

		activeEnv := config.GetActiveEnv()
		activeEnvName := activeEnv.GetName()

		if len(envs) > 0 {
			for _, envName := range envs {
				if envName == activeEnvName {
					fmt.Printf("%s[active] %s", chalk.Cyan, chalk.ResetColor)
				} else {
					fmt.Print("         ")
				}

				fmt.Print(chalk.Bold.TextStyle(pad(envName, maxEnvNameLength)))

				if env, err := config.GetEnv(envName); err == nil {
					if env.IsRunning() {
						fmt.Printf("%s%s%s", chalk.Green, " running", chalk.ResetColor)
					} else {
						fmt.Print(" -------")
					}

					if pathFlag {
						fmt.Print(" " + env.GetHomeDir())
					}
				}

				fmt.Println()
			}
		} else {
			fmt.Printf("No environments configured. Add a new environment by running %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)
		}
	},
}

func init() {
	EnvCommand.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&pathFlag, "path", "p", false, "print path")
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
