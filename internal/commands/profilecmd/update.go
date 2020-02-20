package profilecmd

import (
	"errors"
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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && !config.GetActiveEnv().HasProfile(args[0]) {
			return fmt.Errorf("No such profile: %s", args[0])
		}

		if len(args) > 1 {
			return errors.New("Command only takes one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		envHomeDir := activeEnv.GetHomeDir()

		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("No profiles in environment: %s\n", chalk.Cyan.Color(activeEnv.GetName()))

			os.Exit(0)
		}

		var profileName string
		if len(args) == 0 {
			profileName = survey.Select("Select profile to update", profileNames)
		} else {
			profileName = args[0]
		}

		services, err := dockercompose.GetServices(envHomeDir)
		utils.ErrorAndExit(err)

		profile, err := activeEnv.GetProfile(profileName)
		utils.ErrorAndExit(err)

		selected := survey.MultiSelect(fmt.Sprintf("Select services for profile %s", chalk.Cyan.Color(profileName)), services.All, profile.Selected)

		if len(selected) == 0 {
			utils.Warn("No services selected.")
		}

		viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selected)

		config.Save(fmt.Sprintf("Updated profile of env %s: %s", chalk.Bold.TextStyle(activeEnv.GetName()), chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to update profile: %s", profileName))
	},
}

func init() {
	ProfileCommand.AddCommand(updateCmd)
}
