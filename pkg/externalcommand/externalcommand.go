package externalcommand

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// JoinCommand joins all command slices with spaces
func JoinCommand(command string, arguments ...string) string {
	if len(arguments) > 0 {
		return fmt.Sprintf("%s ", command) + strings.Join(arguments, " ")
	} else {
		return command
	}
}

// Execute runs an external command string and automatically provides some output via console or logfile
func Execute(command string, timebridger Timebridger, logfile string) ([]byte, error) {
	splittedCmd := strings.Split(command, " ")

	cmd := exec.Command(splittedCmd[0], splittedCmd[1:]...)

	var output []byte
	var commandError error
	if timebridger == nil {
		var buffer bytes.Buffer

		cmd.Stdout = io.MultiWriter(os.Stdout, &buffer)
		cmd.Stderr = io.MultiWriter(os.Stderr, &buffer)

		commandError = cmd.Run()

		output = buffer.Bytes()
	} else {
		timebridger.Start(command)
		defer timebridger.Stop()

		output, commandError = cmd.CombinedOutput()
	}

	if logfile != "" {
		fileError := ioutil.WriteFile(logfile, output, 0644)

		if fileError != nil {
			return output, fmt.Errorf("Could not save output to logfile: %s", logfile)
		}
	}

	if commandError != nil {
		return output, fmt.Errorf("Could not run command: %s", command)
	}

	return output, nil
}
