package envvars

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// SetEnvVars sets env vars for given env
func SetEnvVars(services []string, selected []string) error {
	for _, service := range services {
		key := fmt.Sprintf("%s_HOST", strings.ToUpper(strings.ReplaceAll(service, "-", "_")))

		var val string
		for _, selectedService := range selected {
			if service == selectedService {
				val = service

				break
			}
		}

		if val == "" {
			val = "docker.host.internal"
		}

		err := os.Setenv(key, val)

		if err != nil {
			return errors.New("setting environment variables for docker dns failed")
		}
	}

	return nil
}
