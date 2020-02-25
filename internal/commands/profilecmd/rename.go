package profilecmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getRenameCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "rename [profile]",
		Short:   "Rename profile",
		Long:    "Rename profile",
		Example: "dockma profile rename",
		Args:    argvalidators.OptionalProfile,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runRenameCommand,
	}
}

func runRenameCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	profileNames := activeEnv.GetProfileNames()

	if len(profileNames) == 0 {
		fmt.Printf("No profiles in environment %s.\n", chalk.Cyan.Color(activeEnv.GetName()))

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

	config.Save(fmt.Sprintf("Renamed profile of environment %s from %s to %s.", chalk.Cyan.Color(activeEnv.GetName()), chalk.Yellow.Color(renameProfile), chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to rename profile: %s", renameProfile))
}
