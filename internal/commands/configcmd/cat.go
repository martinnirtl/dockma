package configcmd

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/martinnirtl/dockma/internal/config"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/spf13/cobra"
	"github.com/ttacon/chalk"
)

func getCatCommand() *cobra.Command {
	return &cobra.Command{
		Use:     "cat",
		Short:   "Print config.json of Dockma",
		Long:    "Print config.json of Dockma",
		Example: "dockma config cat",
		Args:    cobra.NoArgs,
		Run:     runCatCommand,
	}
}

func runCatCommand(cmd *cobra.Command, args []string) {
	filepath := config.GetFile("config.json")

	content, err := ioutil.ReadFile(filepath)

	if err != nil {
		fmt.Println(err)

		utils.ErrorAndExit(errors.New("Could not read config file"))
	}

	fmt.Printf("Here comes the %s file:\n", chalk.Cyan.Color("Dockma config"))
	fmt.Println(string(content))
}
