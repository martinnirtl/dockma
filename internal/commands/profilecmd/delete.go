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

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"del"},
	Short:   "Delete a profile of active environment",
	Long:    "Delete a profile of active environment",
	Example: "dockma profile delete",
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

		var profileName string
		if len(args) == 0 {
			profileName = survey.Select("Select profile to be deleted", profileNames)
		} else {
			profileName = args[0]

			proceed := survey.Confirm("Are you sure", true)
			if !proceed {
				utils.Abort()
			}
		}

		profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()))

		profileMap[profileName] = nil

		viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()), profileMap)

		config.Save(fmt.Sprintf("Deleted profile from %s environment: %s", chalk.Bold.TextStyle(activeEnv.GetName()), chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to delete profile '%s'", profileName))
	},
}

func init() {
	ProfileCommand.AddCommand(deleteCmd)
}
