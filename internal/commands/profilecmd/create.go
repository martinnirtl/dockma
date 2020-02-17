package profilecmd

import (
	"errors"
	"fmt"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create named service selection.",
	Long:    "Create named service selection.",
	Args:    cobra.NoArgs,
	Example: "dockma profile create",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		envHomeDir := activeEnv.GetHomeDir()

		profileName := survey.InputName("Enter name for profile", "")

		if activeEnv.HasProfile(profileName) {
			utils.ErrorAndExit(errors.New("Profile name already taken. Use 'update' to reselect services"))
		}

		// FIXME use regex
		if profileName == "" || profileName == "-" {
			utils.ErrorAndExit(errors.New("Invalid profile name"))
		}

		services, err := dockercompose.GetServices(envHomeDir)

		if err != nil {
			utils.ErrorAndExit(errors.New("Could not read services"))
		}

		selected := survey.MultiSelect(fmt.Sprintf("Select services for profile %s%s%s", chalk.Cyan, profileName, chalk.ResetColor), services.All, nil)

		if len(selected) == 0 {
			fmt.Printf("%sNo services selected%s\n\n", chalk.Yellow, chalk.ResetColor)
		}

		viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selected)

		config.Save(fmt.Sprintf("Saved profile: %s%s%s [%s]", chalk.Cyan, profileName, chalk.ResetColor, activeEnv.GetName()), fmt.Errorf("Failed to save profile '%s'", profileName))
	},
}

func init() {
	ProfileCommand.AddCommand(createCmd)
}
