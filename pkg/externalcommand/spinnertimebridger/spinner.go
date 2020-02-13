package spinnertimebridger

import (
	"time"

	spinnerlib "github.com/briandowns/spinner"
)

// SpinnerTimebridger implements timebridger interface
type SpinnerTimebridger struct {
	command string
	message string
	spinner *spinnerlib.Spinner
}

// Start should get called after external command execution starts
func (otb *SpinnerTimebridger) Start(command string) error {
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
func (otb *SpinnerTimebridger) Update(update string) error {
	otb.command = update

	if otb.message == "" {
		otb.spinner.Suffix = " " + update
	}

	return nil
}

// Stop should get called after external command execution finishes
func (otb *SpinnerTimebridger) Stop() error {
	otb.spinner.Stop()

	return nil
}

// New creates a new OutputTimebidger
func New(message string, spinner int, color string) *SpinnerTimebridger {
	s := spinnerlib.New(spinnerlib.CharSets[spinner], 100*time.Millisecond)
	s.Color(color, "bold")

	return &SpinnerTimebridger{
		message: message,
		spinner: s,
	}
}

// Default creates the default OutputTimebidger
func Default(message string) *SpinnerTimebridger {
	s := spinnerlib.New(spinnerlib.CharSets[14], 100*time.Millisecond)
	s.Color("cyan", "bold")

	return &SpinnerTimebridger{
		message: message,
		spinner: s,
	}
}
