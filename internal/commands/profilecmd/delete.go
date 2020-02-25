package profilecmd

import (
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

func getDeleteCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "delete [profiles...]",
		Aliases: []string{"del"},
		Short:   "Delete a profile of active environment",
		Long:    "Delete a profile of active environment",
		Example: "dockma profile delete",
		Args:    argvalidators.OnlyProfiles,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runDeleteCommand,
	}
}

func runDeleteCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	profileNames := activeEnv.GetProfileNames()

	if len(profileNames) == 0 {
		fmt.Printf("No profiles in environment %s.\n", chalk.Cyan.Color(activeEnv.GetName()))

		os.Exit(0)
	}

	var selected []string
	if len(args) == 0 {
		selected = survey.MultiSelect("Select profiles to be deleted", profileNames, nil)
	} else {
		selected = args
	}

	proceed := survey.Confirm("Are you sure", true)
	if !proceed {
		utils.Abort()
	}

	activeEnvName := activeEnv.GetName()
	profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnvName))
	for _, profileName := range selected {
		profileMap[profileName] = nil

		config.Save(fmt.Sprintf("Deleted profile %s from environment %s.", chalk.Cyan.Color(profileName), chalk.Cyan.Color(activeEnvName)), fmt.Errorf("Failed to delete profile: %s", profileName))
	}

	viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnvName), profileMap)
}
