package spinnertimebridger

import (
	"fmt"
	"time"

	spinnerlib "github.com/briandowns/spinner"
)

// SpinnerTimebridger implements timebridger interface
type SpinnerTimebridger struct {
	command          string
	templateRunning  string
	templateFinished string
	spinner          *spinnerlib.Spinner
}

// Start should get called after external command execution starts
func (otb *SpinnerTimebridger) Start(command string) error {
	otb.command = command

	if otb.templateRunning != "" {
		otb.spinner.Suffix = " " + fmt.Sprintf(otb.templateRunning, command)
	} else {
		otb.spinner.Suffix = " " + command
	}

	otb.spinner.Start()

	return nil
}

// Update should get called if description changes during command execution
func (otb *SpinnerTimebridger) Update(update string) error {
	otb.command = update

	if otb.templateRunning != "" {
		otb.spinner.Suffix = " " + fmt.Sprintf(otb.templateRunning, update)
	} else {
		otb.spinner.Suffix = " " + update
	}

	return nil
}

// Stop should get called after external command execution finishes
func (otb *SpinnerTimebridger) Stop() error {
	if otb.templateFinished != "" {
		otb.spinner.FinalMSG = fmt.Sprintf(otb.templateFinished+"\n", otb.command)
	}

	otb.spinner.Stop()

	return nil
}

// New creates a new OutputTimebidger
func New(templateRunning string, templateFinished string, spinner int, color string) *SpinnerTimebridger {
	s := spinnerlib.New(spinnerlib.CharSets[spinner], 100*time.Millisecond)
	s.Color(color, "bold")

	return &SpinnerTimebridger{
		templateRunning:  templateRunning,
		templateFinished: templateFinished,
		spinner:          s,
	}
}
