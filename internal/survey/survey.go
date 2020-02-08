package survey

import (
	"fmt"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func checkError(err error) {
	if err == terminal.InterruptErr {
		fmt.Println("Interrupted.")

		os.Exit(0)
	} else if err != nil {
		panic(err)
	}
}

// Confirm abstracts survey's confirm and adds styling
func Confirm(message string, preselected bool) (confirm bool) {
	err := survey.AskOne(&survey.Confirm{
		Message: message,
		Default: preselected,
	}, &confirm)

	checkError(err)

	return
}

// Input abstracts survey's input and adds styling
func Input(message string, suggestion string) (response string) {
	err := survey.AskOne(&survey.Input{
		Message: message,
		Default: suggestion,
	}, &response)

	checkError(err)

	return
}

// Select abstracts survey's select and adds styling
func Select(message string, options []string) (selection string) {
	err := survey.AskOne(&survey.Select{
		Message: message,
		Options: options,
	}, &selection, survey.WithIcons(func(icons *survey.IconSet) {
		icons.SelectFocus.Text = "❯"

		icons.SelectFocus.Format = "cyan+b"
	}))

	checkError(err)

	return
}

// MultiSelect abstracts survey's multiselect and adds styling
func MultiSelect(message string, options []string, preselected []string) (selection []string) {
	err := survey.AskOne(&survey.MultiSelect{
		Message:  message,
		Options:  options,
		Default:  preselected,
		PageSize: len(options),
	}, &selection, survey.WithIcons(func(icons *survey.IconSet) {
		icons.UnmarkedOption.Text = "◯"
		icons.MarkedOption.Text = "◉"
		icons.SelectFocus.Text = "❯"

		icons.MarkedOption.Format = "green+b"
		icons.SelectFocus.Format = "cyan+b"
	}))

	checkError(err)

	return
}
