package survey

import (
	"fmt"
	"regexp"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/ttacon/chalk"
)

// NameRegex should be used to verify all names
var NameRegex string = "^[a-zA-Z].[-.a-zA-Z0-9]*[a-zA-Z0-9]$"

// ValidateName implements the validate func type
func ValidateName(val interface{}) error {
	switch name := val.(type) {
	case string:
		match, err := CheckName(name)

		if err != nil {
			return err
		}

		if !match {
			return fmt.Errorf("%s has to comply with regex: %s", chalk.Bold.TextStyle(name), NameRegex)
		}

		return nil
	default:
		return fmt.Errorf("Input is not a string")
	}
}

// CheckName verifies the name against utils.NameRegex
func CheckName(name string) (match bool, err error) {
	match, err = regexp.MatchString(NameRegex, name)

	if err != nil {
		return false, fmt.Errorf("Matching string with regex failed")
	}

	return
}

// Confirm abstracts survey's confirm and adds styling
func Confirm(message string, preselected bool) (confirm bool) {
	err := survey.AskOne(&survey.Confirm{
		Message: message,
		Default: preselected,
	}, &confirm, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "magenta+hb"
	}))

	checkError(err)

	return
}

// Input abstracts survey's input and adds styling
func Input(message string, suggestion string) (response string) {
	err := survey.AskOne(&survey.Input{
		Message: message,
		Default: suggestion,
	}, &response, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "magenta+hb"
	}))

	checkError(err)

	return
}

// InputName abstracts survey's input and validates input
func InputName(message string, suggestion string) (response string) {
	err := survey.AskOne(&survey.Input{
		Message: message,
		Default: suggestion,
		Help:    fmt.Sprintf("Name has to comply with regex: %s", NameRegex),
	}, &response, survey.WithValidator(ValidateName), survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "magenta+hb"
	}))

	checkError(err)

	return
}

// Select abstracts survey's select and adds styling
func Select(message string, options []string) (selection string) {
	err := survey.AskOne(&survey.Select{
		Message: message,
		Options: options,
	}, &selection, survey.WithIcons(func(icons *survey.IconSet) {
		icons.Question.Format = "magenta+hb"

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
		icons.Question.Format = "magenta+hb"

		icons.UnmarkedOption.Text = "◯"

		icons.MarkedOption.Text = "◉"
		icons.MarkedOption.Format = "cyan+b"

		icons.SelectFocus.Text = "❯"
		icons.SelectFocus.Format = "cyan+b"
	}))

	checkError(err)

	return
}

func checkError(err error) {
	if err == terminal.InterruptErr {
		utils.Abort()
	} else if err != nil {
		panic(err)
	}
}
