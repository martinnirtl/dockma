package profilecmd

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

var servicesFlag bool

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List profiles of active environment.",
	Long:    `-`,
	Example: "dockma profile list",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		profileNames := config.GetProfilesNames(activeEnv)

		for _, profileName := range profileNames {
			fmt.Printf("%s%s%s\n", chalk.Cyan, profileName, chalk.ResetColor)

			if servicesFlag {
				profile, err := config.GetProfile(activeEnv, profileName)

				if err != nil {
					utils.Error(err)
				}

				for _, service := range profile.Services {
					if utils.Includes(profile.Selected, service) {
						fmt.Printf("- %s%s%s\n", chalk.Green, service, chalk.ResetColor)
					} else {
						fmt.Printf("- %s\n", service)
					}
				}

				fmt.Println()
			}

		}
	},
}

func init() {
	ProfileCommand.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&servicesFlag, "services", "S", false, "Print services")
}
