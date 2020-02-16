package spinnertimebridger

import (
	"time"

	spinnerlib "github.com/briandowns/spinner"
	"github.com/martinnirtl/dockma/pkg/externalcommand"
)

var defaultSpinnter int = 14
var defaultColor string = "cyan"

// SpinnerTimebridger implements timebridger interface
type spinnerTimebridger struct {
	command string
	message string
	spinner *spinnerlib.Spinner
}

// Start should get called after external command execution starts
func (otb *spinnerTimebridger) Start(command string) error {
	otb.command = command

	if otb.message != "" {
		otb.spinner.Suffix = " " + otb.message
	} else {
		otb.spinner.Suffix = " " + command
	}

	otb.spinner.Start()

	return nil
}

// Update should get called if description changes during command execution
func (otb *spinnerTimebridger) Update(update string) error {
	otb.command = update

	if otb.message == "" {
		otb.spinner.Suffix = " " + update
	}

	return nil
}

// Stop should get called after external command execution finishes
func (otb *spinnerTimebridger) Stop() error {
	otb.spinner.Stop()

	return nil
}

// New creates a new OutputTimebidger
func New(message string, spinner int, color string) externalcommand.Timebridger {
	s := spinnerlib.New(spinnerlib.CharSets[spinner], 100*time.Millisecond)
	s.Color(color, "bold")

	return &spinnerTimebridger{
		message: message,
		spinner: s,
	}
}

// SetDefault sets spinner type and color of spinner
func SetDefault(spinner int, color string) {
	defaultSpinnter = spinner
	defaultColor = color
}

// Default creates the default OutputTimebidger
func Default(message string) externalcommand.Timebridger {
	s := spinnerlib.New(spinnerlib.CharSets[defaultSpinnter], 100*time.Millisecond)
	s.Color(defaultColor, "bold")

	return &spinnerTimebridger{
		message: message,
		spinner: s,
	}
}
