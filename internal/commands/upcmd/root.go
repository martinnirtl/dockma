package upcmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/envcmd"
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

// UpCommand implements the top level up command
var UpCommand = &cobra.Command{
	Use:   "up",
	Short: "Runs active environment with service selection.",
	Long:  "Runs active environment with service selection.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		filepath := config.GetLogfile()

		activeEnv := config.GetActiveEnv()

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("envs.%s.home", activeEnv))

		autoPull := config.GetAutoPullSetting(activeEnv)

		var pull bool
		switch autoPull {
		case "auto":
			pull = true
		case "optional":
			pull = survey.Confirm("Pull changes from git", false)
		case "manual":
			timePassed, err := config.GetDurationPassedSinceLastUpdate(activeEnv)

			if err != nil {
				pull = survey.Confirm(fmt.Sprintf("Environment never got updated (%s). Wanna pull now", utils.PrintCyan("dockma env pull")), true)
			} else if timePassed.Hours() > 24*7 {
				pull = survey.Confirm("Some time has passed since last git pull. Wanna pull now", true)
			}
		case "off":
			pull = false
		default:
			pull = false
		}

		if pull {
			err := envcmd.Pull(envHomeDir, false)

			if err != nil {
				fmt.Printf("%sCould not execute git pull.%s\n", chalk.Yellow, chalk.ResetColor)
			}
		}

		profileNames := config.GetProfilesNames(activeEnv)

		var preselected []string

		// default
		profileName := "latest"
		if len(profileNames) > 0 {
			profileNames = append(profileNames, "latest")

			profileName := survey.Select(fmt.Sprintf("Select profile to run or %slatest%s", chalk.Cyan, chalk.ResetColor), profileNames)

			if profileName != "latest" {
				profile, err := config.GetProfile(activeEnv, profileName)

				utils.Error(err)

				preselected = profile.Selected
			} else {
				profile, err := config.GetLatest(activeEnv)

				utils.Error(err)

				preselected = profile.Selected
			}
		}

		services, err := dockercompose.GetServices(envHomeDir)

		if len(preselected) == 0 {
			preselected = services.All
		}

		utils.Error(err)

		if len(services.Override) > 0 {
			fmt.Printf("%sFound %d services in docker-compose.override.y(a)ml: %s%s\n\n", chalk.Yellow, len(services.Override), strings.Join(services.Override, ", "), chalk.ResetColor)
		}

		selectedServices := survey.MultiSelect("Select services to start", services.All, preselected)

		if profileName == "latest" {
			saveProfile := survey.Confirm("Save as profile", false)

			if saveProfile {
				profileName = survey.Input("Enter profile name", "")

				viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv, profileName), selectedServices)
			} else {
				viper.Set(fmt.Sprintf("envs.%s.latest", activeEnv), selectedServices)
			}

			err = config.Save()

			utils.Error(err)
		}

		err = envvars.SetEnvVars(services.All, selectedServices)

		utils.Error(err)

		err = os.Chdir(envHomeDir)

		utils.Error(err)

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
			timebridger = spinnertimebridger.New("Running 'docker-compose up'", fmt.Sprintf("%sSuccessfully executed 'docker-compose up'%s", chalk.Green, chalk.ResetColor), 14, "cyan")
		}

		command := externalcommand.JoinCommand("docker-compose up -d", selectedServices...)

		_, err = externalcommand.Execute(command, timebridger, filepath)

		utils.Error(err)
	},
}
