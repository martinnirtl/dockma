package profilecmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var updateCmd = &cobra.Command{
	Use:     "update",
	Aliases: []string{"upd"},
	Short:   "Update profile's service selection",
	Long:    "Update profile's service selection",
	Example: "dockma profile update",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		envHomeDir := activeEnv.GetHomeDir()

		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("%sNo profiles created in environment: %s%s\n", chalk.Cyan, activeEnv.GetName(), chalk.ResetColor)

			os.Exit(0)
		}

		profileName := survey.Select("Select profile to update", profileNames)

		services, err := dockercompose.GetServices(envHomeDir)
		utils.ErrorAndExit(err)

		profile, err := activeEnv.GetProfile(profileName)
		utils.ErrorAndExit(err)

		selected := survey.MultiSelect(fmt.Sprintf("Select services for profile %s%s%s", chalk.Cyan, profileName, chalk.ResetColor), services.All, profile.Selected)

		if len(selected) == 0 {
			fmt.Printf("%sNo services selected%s\n\n", chalk.Yellow, chalk.ResetColor)
		}

		viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selected)

		config.Save(fmt.Sprintf("Updated profile: %s [%s]", chalk.Cyan, profileName, chalk.ResetColor, activeEnv.GetName()), fmt.Errorf("Failed to update profile '%s'", profileName))
	},
}

func init() {
	ProfileCommand.AddCommand(updateCmd)
}
