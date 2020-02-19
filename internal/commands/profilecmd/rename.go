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
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) == 1 && !config.GetActiveEnv().HasProfile(args[0]) {
			return fmt.Errorf("No such profile '%s'", args[0])
		}

		if len(args) > 1 {
			return errors.New("Command only takes one argument")
		}

		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("No profiles in environment: %s\n", chalk.Cyan.Color(activeEnv.GetName()))

			os.Exit(0)
		}

		var renameProfile string
		if len(args) == 0 {
			renameProfile = survey.Select("Select profile to update", profileNames)
		} else {
			renameProfile = args[0]
		}

		profileName := survey.InputName("Enter new name for profile", renameProfile)

		if activeEnv.HasProfile(profileName) {
			utils.ErrorAndExit(errors.New("Profile name already taken"))
		}

		profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()))
		profile := profileMap[renameProfile]
		profileMap[renameProfile] = nil
		profileMap[profileName] = profile

		viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()), profileMap)

		config.Save(fmt.Sprintf("Renamed profile from %s to %s %s", chalk.Cyan.Color(renameProfile), chalk.Cyan.Color(profileName), chalk.Bold.TextStyle(fmt.Sprintf("[%s]", activeEnv.GetName()))), fmt.Errorf("Failed to rename profile '%s'", renameProfile))
	},
}

func init() {
	ProfileCommand.AddCommand(renameCmd)
}
