package scriptcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
)

// GetScriptCommand returns the top level script command
func GetScriptCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "script [scriptname]",
		Short:   "Run script located in scripts dir of active environment",
		Long:    "Run script located in scripts dir of active environment",
		Example: "dockma script",
		Args:    cobra.ArbitraryArgs,
		Run:     runScriptCommand,
	}
}

func runScriptCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()

	if activeEnv.GetName() == "-" {
		utils.PrintNoActiveEnvSet()
	}

	envHomeDir := activeEnv.GetHomeDir()

	err := os.Chdir(envHomeDir)
	utils.ErrorAndExit(err)

	var scriptname string
	var scriptArgs []string
	if len(args) > 0 {
		scriptname = args[0]
		scriptArgs = args[1:]

		if !strings.HasSuffix(scriptname, ".sh") {
			scriptname = scriptname + ".sh"
		}
	} else {
		files, err := ioutil.ReadDir(filepath.Join(envHomeDir, "scripts"))
		utils.ErrorAndExit(err)

		scripts := make([]string, 0, len(files))
		for _, file := range files {
			if !file.IsDir() && len(file.Name()) > 0 {
				scripts = append(scripts, file.Name())
			}
		}

		scriptname = survey.Select("Select script to run", scripts)

		addArgs := survey.Confirm("Input additional args", false)
		if addArgs {
			scriptArgsString := survey.Input("Input script arguments", "")
			scriptArgs = strings.Split(scriptArgsString, " ")
		}
	}

	baseCommand := fmt.Sprintf("./scripts/%s", scriptname)
	command := externalcommand.JoinCommand(baseCommand, scriptArgs...)

	_, err = externalcommand.Execute(command, nil, "")
	utils.ErrorAndExit(err)

	utils.Success(fmt.Sprintf("Executed script: %s", scriptname))
}
