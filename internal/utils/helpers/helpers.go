package helpers

import (
	"fmt"

	"github.com/ttacon/chalk"
)

// PrintErrorList prints error list as red 'suberrors' to stdout (To be used with config.SaveNow).
func PrintErrorList(errorList []error) {
	for _, err := range errorList {
		fmt.Printf("%s> %s%s\n", chalk.Red, err, chalk.ResetColor)
	}
}

// PrintMessageList prints messages to stdout (To be used with config.SaveNow)
func PrintMessageList(messageList []string) {
	for _, msg := range messageList {
		fmt.Printf("%s\n", msg)
	}
}
