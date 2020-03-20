package helpers

import (
	"fmt"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/ttacon/chalk"
)

// PrintErrorList prints error list as red 'suberrors' to stdout (To be used with config.SaveNow).
func PrintErrorList(errorList []error) {
	for _, err := range errorList {
		fmt.Println(chalk.Red.Color(fmt.Sprintf("> %s", err)))
	}
}

// PrintMessageList prints messages to stdout (To be used with config.SaveNow)
func PrintMessageList(messageList []string) {
	for _, msg := range messageList {
		utils.Success(msg)
	}
}
