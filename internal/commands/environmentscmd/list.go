package environmentscmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// TODO could be table with props from envs
var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List all configured environments",
	Long:    `-`,
	Example: "dockma envs list",
	Run: func(cmd *cobra.Command, args []string) {
		envs := utils.GetEnvironments()

		activeEnv := viper.GetString("active")

		if len(envs) > 0 {
			for _, env := range envs {
				if env == activeEnv {
					fmt.Printf("%s%s [active]%s\n", chalk.Cyan, env, chalk.ResetColor)
				} else {
					fmt.Println(env)
				}
			}
		} else {
			fmt.Printf("No environments configured. Add a new environment by running %sdockma envs init%s.\n", chalk.Cyan, chalk.ResetColor)
		}
	},
}

func init() {
	EnvironmentsCommand.AddCommand(listCmd)
}
