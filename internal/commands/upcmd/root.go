package upcmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/pullcmd"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/envvars"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/ttacon/chalk"
)

// UpCommand implements the top level dockma command up
var UpCommand = &cobra.Command{
	Use:   "up",
	Short: "Runs active environment with service selection.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetLogfile()

		activeEnv := config.GetActiveEnv()

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		autoPull := config.IsAutoPullSet(activeEnv)

		if autoPull {
			err := pullcmd.Pull(envHomeDir, false)

			if err != nil {
				fmt.Printf("%sCould not execute git pull.%s\n", chalk.Yellow, chalk.ResetColor)
			}
		}

		profileNames := config.GetProfilesNames(activeEnv)

		var preselected []string

		// default
		profileName := "CUSTOM"

		if len(profileNames) > 0 {
			profileNames = append(profileNames, "CUSTOM")

			profileName, err := survey.Select("Select profile to run or CUSTOM", profileNames)

			if err != nil {
				utils.Abort()
			}

			profile, err := config.GetProfile(activeEnv, profileName)

			if err != nil {
				utils.Error(err)
			}

			preselected = profile.Selected
		}

		services, err := dockercompose.GetServices(envHomeDir)

		if len(preselected) == 0 {
			preselected = services.All
		}

		if err != nil {
			utils.Error(err)
		}

		if len(services.Override) > 0 {
			fmt.Printf("%sFound %d services in docker-compose.override.y(a)ml: %s%s\n\n", chalk.Yellow, len(services.Override), strings.Join(services.Override, ", "), chalk.ResetColor)
		}

		selectedServices, err := survey.MultiSelect("Select services to start", services.All, preselected)

		if err != nil {
			utils.Abort()
		}

		if profileName == "CUSTOM" {
			saveProfile, err := survey.Confirm("Save as profile", false)

			if err != nil {
				utils.Abort()
			}

			if saveProfile {
				profileName, err := survey.Input("Enter profile name", "")

				if err != nil {
					utils.Abort()
				}

				viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv, profileName), selectedServices)
			}
		}

		err = envvars.SetEnvVars(services.All, selectedServices)

		if err != nil {
			utils.Error(err)
		}

		err = os.Chdir(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
			timebridger = spinnertimebridger.New("Running 'docker-compose up'", fmt.Sprintf("%sSuccessfully executed 'docker-compose up'%s", chalk.Green, chalk.ResetColor), 14, "cyan")
		}

		command := externalcommand.JoinCommandSlices("docker-compose up -d", selectedServices...)

		_, err = externalcommand.Execute(command, timebridger, filepath)

		if err != nil {
			utils.Error(err)
		}
	},
}
