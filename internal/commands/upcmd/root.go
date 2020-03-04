package upcmd

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/argvalidators"
	"github.com/martinnirtl/dockma/internal/commands/envcmd"
	"github.com/martinnirtl/dockma/internal/commands/hooks"
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

var latestFlag bool
var skipFlag bool

// GetUpCommand returns the top level up command
func GetUpCommand() *cobra.Command {
	upCommand := &cobra.Command{
		Use:     "up [profile]",
		Aliases: []string{"run"},
		Short:   "Runs active environment with profile or service selection",
		Long:    "Runs active environment with profile or service selection",
		Example: "dockma up",
		Args:    argvalidators.OptionalProfile,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runUpCommand,
	}

	upCommand.Flags().BoolVarP(&latestFlag, "latest", "l", false, "use latest service selection")
	upCommand.Flags().BoolVarP(&skipFlag, "select", "s", false, "select services")

	return upCommand
}

func runUpCommand(cmd *cobra.Command, args []string) {
	filepath := config.GetSubcommandLogfile()
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	pull(activeEnv)

	services, err := dockercompose.GetServices(envHomeDir)
	utils.ErrorAndExit(err)

	if len(services.Override) > 0 {
		utils.Warn(fmt.Sprintf("Found %d services in docker-compose override file: %s", len(services.Override), strings.Join(services.Override, ", ")))
	}

	profileNames := activeEnv.GetProfileNames()

	// default
	profileName := "$latest"
	var preselected []string

	if latestFlag {
		profile, err := activeEnv.GetLatest()
		utils.ErrorAndExit(err)

		preselected = profile.Selected
	} else if len(args) > 0 {
		profileName = args[0]

		profile, err := activeEnv.GetProfile(profileName)
		utils.ErrorAndExit(err)

		preselected = profile.Selected
	} else if len(profileNames) > 0 {
		profileNames = append(profileNames, "$custom")
		profileNames = append(profileNames, "$latest")

		profileName = survey.Select(fmt.Sprintf("Select profile to run"), profileNames)

		if profileName == "$custom" {
			preselected = services.All
		} else if profileName == "$latest" {
			profile, err := activeEnv.GetLatest()
			utils.ErrorAndExit(err)

			preselected = profile.Selected
		} else {
			profile, err := activeEnv.GetProfile(profileName)
			utils.ErrorAndExit(err)

			preselected = profile.Selected
		}
	}

	// if len(preselected) == 0 {
	// 	preselected = services.All
	// }

	var selectedServices []string
	if skipFlag {
		selectedServices = preselected
	} else {
		selectedServices = survey.MultiSelect("Select services to start", services.All, preselected)
	}

	if profileName == "$custom" || (profileName == "$latest" && !skipFlag) {
		saveProfile := survey.Confirm("Save as profile", false)

		if saveProfile {
			profileName = survey.InputName("Enter profile name", "")

			viper.Set(fmt.Sprintf("envs.%s.profiles.%s", activeEnv.GetName(), profileName), selectedServices)
			config.Save(fmt.Sprintf("Saved profile: %s", chalk.Cyan.Color(profileName)), fmt.Errorf("Failed to save profile: %s", profileName))
		}
	}

	viper.Set(fmt.Sprintf("envs.%s.latest", activeEnv.GetName()), selectedServices)
	config.Save("", errors.New("Failed to save latest selection"))

	err = envvars.SetEnvVars("", services.All, selectedServices)
	utils.ErrorAndExit(err)

	err = os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	hideCmdOutput := config.GetHideSubcommandOutputSetting()

	var timebridger externalcommand.Timebridger
	if hideCmdOutput {
		timebridger = spinnertimebridger.Default(fmt.Sprintf("Running %s", chalk.Cyan.Color("docker-compose up")))
	}

	command := externalcommand.JoinCommand("docker-compose up -d", selectedServices...)

	output, err := externalcommand.Execute(command, timebridger, filepath)
	if err != nil && timebridger != nil {
		fmt.Print(string(output))
	}
	utils.ErrorAndExit(err)

	viper.Set(fmt.Sprintf("envs.%s.running", activeEnv.GetName()), true)
	config.Save("", fmt.Errorf("Failed to set environment %s to %s", chalk.Underline.TextStyle(activeEnv.GetName()), chalk.Underline.TextStyle("running")))

	utils.Success("Executed command: docker-compose up")
}

func pull(activeEnv config.Env) {
	envHomeDir := activeEnv.GetHomeDir()
	pull := activeEnv.GetPullSetting()

	var doPull bool
	switch pull {
	case "auto":
		doPull = true
	case "optional":
		doPull = survey.Confirm("Pull changes from git", true)
	case "manual":
		timePassed, err := activeEnv.LastUpdate()

		if err != nil {
			doPull = survey.Confirm(fmt.Sprintf("Environment %s not yet updated (%s). Pull now", chalk.Cyan.Color(activeEnv.GetName()), chalk.Cyan.Color("dockma env pull")), true)
		} else if timePassed.Hours() > 24*7 {
			doPull = survey.Confirm(fmt.Sprintf("Some time has passed since last %s. Pull now", chalk.Cyan.Color("git pull")), true)
		}
	case "off":
		doPull = false
	default:
		doPull = false
	}

	hideCmdOutput := config.GetHideSubcommandOutputSetting()

	if doPull {
		output, err := envcmd.Pull(envHomeDir, hideCmdOutput, false)
		if err != nil && hideCmdOutput {
			fmt.Print(string(output))
			utils.Warn(fmt.Sprintf("Could not execute command %s.", chalk.Underline.TextStyle("git pull")))

			fmt.Println() // Add empty line for better readability
		} else {
			utils.Success(fmt.Sprintf("Pulled environment: %s", activeEnv.GetName()))
		}
	}
}
