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

		if activeEnv.GetName() == "-" {
			utils.NoEnvs()
		}

		envHomeDir := activeEnv.GetHomeDir()

		services, err := dockercompose.GetServices(envHomeDir)
		utils.Error(err)

		if len(services.Override) > 0 {
			fmt.Printf("%sFound %d services in docker-compose override file: %s%s\n", chalk.Yellow, len(services.Override), strings.Join(services.Override, ", "), chalk.ResetColor)
		}

		pull := activeEnv.GetPullSetting()

		var doPull bool
		switch pull {
		case "auto":
			doPull = true
		case "optional":
			doPull = survey.Confirm("Pull changes from git", false)
		case "manual":
			timePassed, err := activeEnv.LastUpdate()

			if err != nil {
				doPull = survey.Confirm(fmt.Sprintf("Environment never got updated (%s). Wanna pull now", utils.PrintCyan("dockma env pull")), true)
			} else if timePassed.Hours() > 24*7 {
				doPull = survey.Confirm("Some time has passed since last git pull. Wanna pull now", true)
			}
		case "off":
			doPull = false
		default:
			doPull = false
		}

		if doPull {
			err := envcmd.Pull(envHomeDir, false)

			if err != nil {
				fmt.Printf("%sCould not execute git pull.%s\n", chalk.Yellow, chalk.ResetColor)
			} else {
				fmt.Printf("%sSuccessfully pulled environment.%s\n", chalk.Green, chalk.ResetColor)
			}
		}

		profileNames := activeEnv.GetProfileNames()

		var preselected []string

		// default
		profileName := "latest"
		if len(profileNames) > 0 {
			profileNames = append(profileNames, "latest")

			profileName = survey.Select(fmt.Sprintf("Select profile to run"), profileNames)

			if profileName != "latest" {
				profile, err := activeEnv.GetProfile(profileName)

				utils.Error(err)

				preselected = profile.Selected
			} else {
				profile, err := activeEnv.GetLatest()

				utils.Error(err)

				preselected = profile.Selected
			}
		}

		if len(preselected) == 0 {
			preselected = services.All
		}

		selectedServices := survey.MultiSelect("Select services to start", services.All, preselected)

		if profileName == "latest" {
			saveProfile := survey.Confirm("Save as profile", false)

			if saveProfile {
				profileName = survey.InputName("Enter profile name", "")

				viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selectedServices)
			} else {
				viper.Set(fmt.Sprintf("envs.%s.latest", activeEnv.GetName()), selectedServices)
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
			timebridger = spinnertimebridger.New("Running 'docker-compose up'", 14, "cyan")
		}

		command := externalcommand.JoinCommand("docker-compose up -d", selectedServices...)

		_, err = externalcommand.Execute(command, timebridger, filepath)

		utils.Error(err)
	},
}
