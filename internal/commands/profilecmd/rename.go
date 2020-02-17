package profilecmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

var renameCmd = &cobra.Command{
	Use:     "rename",
	Short:   "Rename profile",
	Long:    "Rename profile",
	Example: "dockma profile rename",
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()

		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("%sNo profiles created in environment: %s%s\n", chalk.Cyan, activeEnv.GetName(), chalk.ResetColor)

			os.Exit(0)
		}

		renameProfile := survey.Select("Select profile to update", profileNames)

		profileName := survey.InputName("Enter name for profile", renameProfile)

		// FIXME use regex
		if profileName == "" || profileName == "-" {
			utils.ErrorAndExit(errors.New("Invalid profile name"))
		}

		if activeEnv.HasProfile(profileName) {
			utils.ErrorAndExit(errors.New("Profile name already taken. Use 'update' to reselect services"))
		}

		profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()))

		profile := profileMap[renameProfile]

		profileMap[renameProfile] = nil

		profileMap[profileName] = profile

		viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()), profileMap)

		config.Save(fmt.Sprintf("Renamed profile from %s to %s%s%s [%s]", renameProfile, chalk.Cyan, profileName, chalk.ResetColor, activeEnv.GetName()), fmt.Errorf("Failed to rename profile '%s'", renameProfile))
	},
}

func init() {
	ProfileCommand.AddCommand(renameCmd)
}
