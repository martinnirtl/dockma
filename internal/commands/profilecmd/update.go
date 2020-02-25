package profilecmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

func getUpdateCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "update [profile]",
		Aliases: []string{"upd"},
		Short:   "Update profile's service selection",
		Long:    "Update profile's service selection",
		Example: "dockma profile update",
		Args:    argvalidators.OptionalProfile,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runUpdateCommand,
	}
}

func runUpdateCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	profileNames := activeEnv.GetProfileNames()

	if len(profileNames) == 0 {
		fmt.Printf("No profiles in environment %s.\n", chalk.Cyan.Color(activeEnv.GetName()))

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

	config.Save(fmt.Sprintf("Updated profile %s of environment %s.", chalk.Cyan.Color(activeEnv.GetName()), chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to update profile: %s", profileName))
}
