package scriptcmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/martinnirtl/dockma/internal/commands/hooks"
	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

// GetScriptCommand returns the top level script command
func GetScriptCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "script [scriptname]",
		Aliases: []string{"sc"},
		Short:   "Run script (.sh) located in scripts dir of active environment",
		Long:    "Run script (.sh) located in scripts dir of active environment",
		Example: "dockma script",
		Args:    cobra.ArbitraryArgs,
		PreRun:  hooks.RequiresActiveEnv,
		Run:     runScriptCommand,
	}
}

func runScriptCommand(cmd *cobra.Command, args []string) {
	activeEnv := config.GetActiveEnv()
	envHomeDir := activeEnv.GetHomeDir()

	scriptsDir := filepath.Join(envHomeDir, "scripts")

	var scriptname string
	var scriptArgs []string
	if len(args) > 0 {
		scriptname = args[0]
		scriptArgs = args[1:]

		if !strings.HasSuffix(scriptname, ".sh") {
			scriptname = scriptname + ".sh"
		}
	} else {
		files, err := ioutil.ReadDir(scriptsDir)
		if err != nil {
			fmt.Println(err)

			utils.ErrorAndExit(fmt.Errorf("Could not read dir %s", chalk.Underline.TextStyle(envHomeDir+"/scripts")))
		}

		scripts := make([]string, 0, len(files))
		for _, file := range files {
			if !file.IsDir() && strings.HasSuffix(file.Name(), ".sh") {
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

	err := os.Chdir(scriptsDir)
	utils.ErrorAndExit(err)

	baseCommand := fmt.Sprintf("./%s", scriptname)
	command := externalcommand.JoinCommand(baseCommand, scriptArgs...)

	_, err = externalcommand.Execute(command, nil, "")
	utils.ErrorAndExit(err)

	utils.Success(fmt.Sprintf("Executed script: %s", scriptname))
}
