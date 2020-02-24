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

func getCreateCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "create [name]",
		Short:   "Create named service selection",
		Long:    "Create named service selection",
		Example: "dockma profile create my-profile",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) == 1 {
				match, err := survey.CheckName(args[0])

				if !match {
					return fmt.Errorf("Given name does not match regex: %s", survey.NameRegex)
				}

				if err != nil {
					return errors.New("Invalid input")
				}
			}

			if len(args) > 1 {
				return errors.New("Command only takes one argument")
			}

			return nil
		},
		Run: runCreateCommand,
	}
}

func runCreateCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	var profileName string
	if len(args) == 0 {
		profileName = survey.InputName("Enter name for profile", "")
	} else {
		profileName = args[0]
	}

	if activeEnv.HasProfile(profileName) {
		utils.ErrorAndExit(errors.New("Profile name already taken"))
	}

	services, err := dockercompose.GetServices(envHomeDir)
	utils.ErrorAndExit(err)

	selected := survey.MultiSelect(fmt.Sprintf("Select services for profile %s", chalk.Cyan.Color(profileName)), services.All, nil)

	if len(selected) == 0 {
		utils.Warn("No services selected.")
	}

	viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selected)

	config.Save(fmt.Sprintf("Saved profile to %s environment: %s", chalk.Bold.TextStyle(activeEnv.GetName()), chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to save profile: %s", profileName))
}
