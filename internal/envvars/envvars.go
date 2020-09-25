package envvars

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// SetEnvVars sets env vars for given service selection
func SetEnvVars(services []string, selected []string) error {
	for _, service := range services {
		key := fmt.Sprintf("DOCKMA_HOST_%s", strings.ToUpper(strings.ReplaceAll(service, "-", "_")))

		var val string
		for _, selectedService := range selected {
			if service == selectedService {
				val = service

				break
			}
		}

		if val == "" {
			val = "host.docker.internal"
		}

		err := os.Setenv(key, val)

		if err != nil {
			return errors.New("setting environment variables for docker dns failed")
		}
	}

	return nil
}
