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
	Short:   "Update profile's service selection.",
	Long:    "Update profile's service selection.",
	Aliases: []string{"upd"},
	Args:    cobra.NoArgs,
	Example: "dockma profile update",
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
		utils.Error(err)

		profile, err := activeEnv.GetProfile(profileName)
		utils.Error(err)

		selected := survey.MultiSelect(fmt.Sprintf("Select services for profile %s%s%s", chalk.Cyan, profileName, chalk.ResetColor), services.All, profile.Selected)

		if len(selected) == 0 {
			fmt.Printf("%sNo services selected%s\n\n", chalk.Yellow, chalk.ResetColor)
		}

		viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selected)

		err = config.Save()

		utils.Error(err)

		utils.Success(fmt.Sprintf("Successfully updated profile: %s [%s]", profileName, activeEnv.GetName()))
	},
}

func init() {
	ProfileCommand.AddCommand(updateCmd)
}
