package dockercompose

import (
	"errors"
	"os"
	"os/exec"
	"strings"
)

// ReadServices returns defined services of docker-compose files
func ReadServices(filepath string) (services []string, err error) {

	os.Chdir(filepath)

	out, e := exec.Command("docker-compose", "config", "--services").Output()

	if e != nil {
		err = errors.New("Could not execute docker-compose")
	} else {
		for _, service := range strings.Split(string(out), "\n") {
			if service != "" {
				services = append(services, service)
			}
		}
	}

	return
}
