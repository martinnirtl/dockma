package envvars

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func includes(slice []string, s string) bool {
	for _, val := range slice {
		if val == s {
			return true
		}
	}

	return false
}

func TestSetEnvVars(t *testing.T) {
	services := []string{"database", "backend", "frontend"}
	selected := []string{"database", "frontend"}

	SetEnvVars(services, selected)

	for _, service := range services {
		envName := fmt.Sprintf("%s_HOST", strings.ToUpper(strings.ReplaceAll(service, "-", "_")))
		envVal := os.Getenv(envName)

		if includes(selected, service) {
			if envVal != service {
				t.Errorf("%s set to %s instead of %s", envName, envVal, service)
			}
		} else {
			if envVal != "host.docker.internal" {
				t.Errorf("%s set to %s instead of %s", envName, envVal, "host.docker.internal")
			}
		}
	}
}
