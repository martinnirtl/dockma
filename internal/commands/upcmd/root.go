package upcmd

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/martinnirtl/dockma/pkg/externalcommand/spinnertimebridger"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var UpCommand = &cobra.Command{
	Use:   "up",
	Short: "Runs active environment with service selection.",
	Long:  "-",
	Run: func(cmd *cobra.Command, args []string) {
		logfileName := viper.GetString("logfile")
		filepath := utils.GetFullLogfilePath(logfileName)

		activeEnv := viper.GetString("active")

		if activeEnv == "-" {
			utils.NoEnvs()
		}

		envHomeDir := viper.GetString(fmt.Sprintf("environments.%s.home", activeEnv))

		services, err := dockercompose.GetServices(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		var selection []string
		survey.AskOne(&survey.MultiSelect{
			Message:  "What days do you prefer:",
			Options:  services.All,
			Default:  services.All,
			PageSize: len(services.All),
		}, &selection)

		err = os.Chdir(envHomeDir)

		if err != nil {
			utils.Error(err)
		}

		var timebridger externalcommand.Timebridger
		if hideCmdOutput := viper.GetBool("hidesubcommandoutput"); hideCmdOutput {
			timebridger = spinnertimebridger.New("Running command '%s'", "", 14, "cyan")
		}

		command := externalcommand.JoinCommandSlices("docker-compose up -d", selection...)

		_, err = externalcommand.Execute(command, timebridger, filepath)

		if err != nil {
			utils.Error(err)
		}
	},
}
