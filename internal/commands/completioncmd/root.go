package completioncmd

import (
	"errors"
	"os"

	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
)

var shells []string = []string{"bash", "powershell", "zsh"}

// GetCompletionCommand returns the top level completion command
func GetCompletionCommand() *cobra.Command {
	return &cobra.Command{
		Use:       "completion [shell]",
		Short:     "Generate shell completion code",
		Long:      "Generate shell completion code",
		Example:   "dockma completion bash",
		Args:      cobra.OnlyValidArgs,
		ValidArgs: shells,
		Run:       runCompletionCommand,
	}
}

func runCompletionCommand(cmd *cobra.Command, args []string) {
	var shell string

	if len(args) > 0 {
		shell = args[0]
	} else {
		shell = survey.Select("Select shell to install completion for", shells)
	}

	var err error
	switch shell {
	case "bash":
		err = cmd.Root().GenBashCompletion(os.Stdout)

	case "powershell":
		err = cmd.Root().GenPowerShellCompletion(os.Stdout)

	case "zsh":
		err = cmd.Root().GenZshCompletion(os.Stdout)
	}

	if err != nil {
		utils.Error(errors.New("Failed to generate shell completion code"))
	}
}
