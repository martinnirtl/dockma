package profilecmd

import (
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
	Short:   "Delete a profile of active environment.",
	Long:    `-`,
	Aliases: []string{"del"},
	Example: "dockma profile delete",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		profileNames := config.GetProfilesNames(activeEnv)

		if len(profileNames) == 0 {
			fmt.Printf("%sNo profiles created in environment: %s%s\n", chalk.Cyan, activeEnv, chalk.ResetColor)

			os.Exit(0)
		}

		profileName, err := survey.Select("Select profile to be deleted", profileNames)

		if err != nil {
			utils.Error(err)
		}

		profileMap := viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", activeEnv))

		profileMap[profileName] = nil

		viper.Set(fmt.Sprintf("envs.%s.profiles", activeEnv), profileMap)

		err = config.Save()

		if err != nil {
			utils.Error(err)
		}

		utils.Success(fmt.Sprintf("Successfully deleted profile: %s [%s]", profileName, activeEnv))
	},
}

func init() {
	ProfileCommand.AddCommand(deleteCmd)
}