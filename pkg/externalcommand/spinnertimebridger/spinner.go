package spinnertimebridger

import (
	"time"

	spinnerlib "github.com/briandowns/spinner"
)

// SpinnerTimebridger implements timebridger interface
type SpinnerTimebridger struct {
	command         string
	messageRunning  string
	messageFinished string
	spinner         *spinnerlib.Spinner
}

// Start should get called after external command execution starts
func (otb *SpinnerTimebridger) Start(command string) error {
	otb.command = command

	if otb.messageRunning != "" {
		otb.spinner.Suffix = " " + otb.messageRunning
	} else {
		otb.spinner.Suffix = " " + command
	}

	otb.spinner.Start()

	return nil
}

// Update should get called if description changes during command execution
func (otb *SpinnerTimebridger) Update(update string) error {
	otb.command = update

	if otb.messageRunning != "" {
		otb.spinner.Suffix = " " + otb.messageRunning
	} else {
		otb.spinner.Suffix = " " + update
	}

	return nil
}

// Stop should get called after external command execution finishes
func (otb *SpinnerTimebridger) Stop() error {
	if otb.messageFinished != "" {
		otb.spinner.FinalMSG = otb.messageFinished + "\n"
	}

	otb.spinner.Stop()

	return nil
}

// New creates a new OutputTimebidger
func New(messageRunning string, messageFinished string, spinner int, color string) *SpinnerTimebridger {
	s := spinnerlib.New(spinnerlib.CharSets[spinner], 100*time.Millisecond)
	s.Color(color, "bold")

	return &SpinnerTimebridger{
		messageRunning:  messageRunning,
		messageFinished: messageFinished,
		spinner:         s,
	}
}
