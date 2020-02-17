package envcmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// TODO could be table with props from envs
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all configured environments.",
	Long:    "List all configured environments.",
	Args:    cobra.NoArgs,
	Example: "dockma envs list",
	Run: func(cmd *cobra.Command, args []string) {
		envs := config.GetEnvNames()

		activeEnv := config.GetActiveEnv()
		activeEnvName := activeEnv.GetName()

		if len(envs) > 0 {
			for _, envName := range envs {
				if envName == activeEnvName {
					fmt.Printf("%s%s [active]%s", chalk.Cyan, envName, chalk.ResetColor)
				} else {
					fmt.Print(envName)
				}

				if env, err := config.GetEnv(envName); err != nil {

				} else {
					if env.IsRunning() {
						fmt.Printf("%s%s%s\n", chalk.Green, " running", chalk.ResetColor)
					} else {
						fmt.Println()
					}
				}
			}
		} else {
			fmt.Printf("No environments configured. Add a new environment by running %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)
		}
	},
}

func init() {
	EnvCommand.AddCommand(listCmd)
}
