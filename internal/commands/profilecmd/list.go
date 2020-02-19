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
	Short:   "List profiles of active environment",
	Long:    "List profiles of active environment",
	Example: "dockma profiles list",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("No profiles in %s. Create one with %s.\n", chalk.Cyan.Color(activeEnv.GetName()), chalk.Cyan.Color("dockma profile create"))
		}

		for _, profileName := range profileNames {
			fmt.Printf("%s%s%s\n", chalk.Cyan, profileName, chalk.ResetColor)

			if servicesFlag {
				profile, err := activeEnv.GetProfile(profileName)

				utils.ErrorAndExit(err)

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

	listCmd.Flags().BoolVarP(&servicesFlag, "services", "S", false, "print services")
}
