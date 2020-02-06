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
	Long:    `-`,
	Example: "dockma profile create [name]",
	Run: func(cmd *cobra.Command, args []string) {
		activeEnv := config.GetActiveEnv()
		envHomeDir := config.GetEnvHomeDir(activeEnv)

		profileName, err := survey.Input("Enter name for profile", "")

		if err != nil {
			utils.Abort()
		}

		if config.HasProfileName(activeEnv, profileName) {
			utils.Error(errors.New("Profile name already taken. Use 'update' to reselect services"))
		}

		// FIXME use regex
		if profileName == "" || profileName == "-" {
			utils.Error(errors.New("Invalid profile name"))
		}

		services, err := dockercompose.GetServices(envHomeDir)

		if err != nil {
			utils.Error(errors.New("Could not read services"))
		}

		selected, err := survey.MultiSelect(fmt.Sprintf("Select services for profile %s%s%s", chalk.Cyan, profileName, chalk.ResetColor), services.All, nil)

		if err != nil {
			utils.Abort()
		}

		if len(selected) == 0 {
			fmt.Printf("%sNo services selected%s\n\n", chalk.Yellow, chalk.ResetColor)

			utils.Abort()
		}

		viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv, profileName), selected)

		config.Save()
	},
}

func init() {
	ProfileCommand.AddCommand(createCmd)
}
