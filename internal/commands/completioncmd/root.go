package completioncmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/martinnirtl/dockma/internal/survey"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
)

var shells []string = []string{"bash", "powershell", "zsh"}

// CompletionCommand implements the top level install command
var CompletionCommand = &cobra.Command{
	Use:     "completion [shell]",
	Short:   "Generate shell completion code",
	Long:    "Generate shell completion code",
	Example: "dockma completion bash",
	Args:    cobra.RangeArgs(0, 1),
	Run: func(cmd *cobra.Command, args []string) {
		var shell string

		if len(args) > 0 {
			if utils.Includes(shells, args[0]) {
				shell = args[0]
			} else {
				utils.ErrorAndExit(fmt.Errorf("Provided type not supported [%s]", strings.Join(shells, "|")))
			}
		} else {
			shell = survey.Select("Select shell to install completion for", shells)
		}

		switch shell {
		case "bash":
			cmd.Root().GenBashCompletion(os.Stdout)

		case "powershell":
			cmd.Root().GenPowerShellCompletion(os.Stdout)

		case "zsh":
			cmd.Root().GenZshCompletion(os.Stdout)
		}
	},
}
