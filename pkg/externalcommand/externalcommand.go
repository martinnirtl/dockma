package externalcommand

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/ttacon/chalk"
)

func startSpinner(cmd string) func() {
	s := spinner.New(spinner.CharSets[14], 100*time.Millisecond)
	s.Suffix = fmt.Sprintf(" Running '%s' (print output: %sdockma inspect%s)", cmd, chalk.Cyan, chalk.ResetColor)
	s.Color("cyan", "bold")
	s.Start()

	return func() {
		s.Stop()
	}
}

// Execute executes an external command string and optionally pipes the output to stdout or stderr
func Execute(cmd string, filename string) (err error) {
	splittedCmd := strings.Split(cmd, " ")

	command := exec.Command(splittedCmd[0], splittedCmd[1:]...)

	if filename == "" {
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr

		err = command.Run()
	} else {
		var output []byte

		stopSpinner := startSpinner(cmd)
		defer stopSpinner()

		output, err = command.CombinedOutput()

		if err != nil {
			fmt.Printf("%sError: Could not run command '%s'%s\n\t%s\n", chalk.Red, cmd, err, chalk.ResetColor)

			return
		}

		ioutil.WriteFile(filename, output, 0644)

		if err != nil {
			fmt.Printf("%sError: Could not save output!%s\n\t%s\n", chalk.Red, err, chalk.ResetColor)

			return
		}

	}

	return
}
