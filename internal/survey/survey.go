package survey

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

// Confirm abstracts survey's confirm and adds styling
func Confirm(message string, preselected bool) (confirm bool, err error) {
	err = survey.AskOne(&survey.Confirm{
		Message: message,
		Default: preselected,
	}, &confirm)

	if err != nil {
		return false, fmt.Errorf("confirm got interrupted")
	}

	return
}

// Input abstracts survey's input and adds styling
func Input(message string, suggestion string) (response string, err error) {
	err = survey.AskOne(&survey.Input{
		Message: message,
		Default: suggestion,
	}, &response)

	if err != nil {
		return "", fmt.Errorf("input got interrupted")
	}

	return
}

// Select abstracts survey's select and adds styling
func Select(message string, options []string) (selection string, err error) {
	err = survey.AskOne(&survey.Select{
		Message: message,
		Options: options,
	}, &selection, survey.WithIcons(func(icons *survey.IconSet) {
		icons.SelectFocus.Text = "❯"

		icons.SelectFocus.Format = "cyan+b"
	}))

	if err != nil {
		return "", fmt.Errorf("select got interrupted")
	}

	return
}

// MultiSelect abstracts survey's multiselect and adds styling
func MultiSelect(message string, options []string, preselected []string) (selection []string, err error) {
	err = survey.AskOne(&survey.MultiSelect{
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

	if err != nil {
		return nil, fmt.Errorf("multiselect got interrupted")
	}

	return
}
