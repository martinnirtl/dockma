package profilecmd

import (
	"fmt"
	"os"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
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
	Args:    cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		profileNames := activeEnv.GetProfileNames()

		if len(profileNames) == 0 {
			fmt.Printf("%sNo profiles created in environment: %s%s\n", chalk.Cyan, activeEnv, chalk.ResetColor)

			os.Exit(0)
		}

		profileName := survey.Select("Select profile to be deleted", profileNames)

		profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()))

		profileMap[profileName] = nil

		viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnv.GetName()), profileMap)

		config.Save(fmt.Sprintf("Deleted profile: %s%s%s [%s]", chalk.Cyan, profileName, chalk.ResetColor, activeEnv.GetName()), fmt.Errorf("Failed to delete profile '%s'", profileName))
	},
}

func init() {
	ProfileCommand.AddCommand(deleteCmd)
}
