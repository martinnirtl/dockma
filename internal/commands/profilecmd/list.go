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

		fmt.Printf("Profiles of %s environment:\n", chalk.Bold.TextStyle(activeEnv.GetName()))

		for _, profileName := range profileNames {
			if servicesFlag {
				fmt.Println()

				fmt.Println(chalk.Bold.TextStyle(profileName))

				profile, err := activeEnv.GetProfile(profileName)

				utils.ErrorAndExit(err)

				for _, service := range profile.Services {
					if utils.Includes(profile.Selected, service) {
						fmt.Printf("- %s\n", chalk.Cyan.Color(service))
					} else {
						fmt.Printf("- %s\n", service)
					}
				}
			} else {
				fmt.Printf("- %s\n", chalk.Cyan.Color(profileName))
			}

		}
	},
}

func init() {
	ProfileCommand.AddCommand(listCmd)

	listCmd.Flags().BoolVarP(&servicesFlag, "services", "S", false, "print services")
}
