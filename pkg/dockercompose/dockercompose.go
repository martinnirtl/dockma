package dockercompose

import (
	"fmt"

	"github.com/spf13/viper"
)

// Services of compose and override file
type Services struct {
	All      []string
	Compose  []string
	Override []string
}

var composeCache map[string]*viper.Viper = make(map[string]*viper.Viper)
var overrideCache map[string]*viper.Viper = make(map[string]*viper.Viper)

func getDockerCompose(filepath string, override bool) (*viper.Viper, error) {
	if !override && composeCache[filepath] != nil {
		return composeCache[filepath], nil
	} else if override && overrideCache[filepath] != nil {
		return overrideCache[filepath], nil
	}

	fileName := "docker-compose"
	if override {
		fileName = "docker-compose.override"
	}

	temp := viper.New()
	temp.SetConfigName(fileName)
	temp.SetConfigType("yaml")
	temp.AddConfigPath(filepath)

	readError := temp.ReadInConfig()

	if readError != nil {
		return nil, fmt.Errorf("Could not read file: %s", fileName)
	}

	return temp, nil
}

// GetVersion returns the compose file's version
func GetVersion(filepath string) string {
	dockercompose, err := getDockerCompose(filepath, false)

	if err != nil {
		return ""
	}

	return dockercompose.GetString("version")
}

// GetServices returns the Services of compose/override files
func GetServices(filepath string) (services Services, err error) {
	dockercompose, err := getDockerCompose(filepath, false)

	services = Services{}

	if err != nil {
		return services, err
	}

	dockercomposeOverride, err := getDockerCompose(filepath, true)

	if err == nil {
		services.Compose = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.Override = getServicesFromStringMap(dockercomposeOverride.GetStringMap("services"))

		services.All = mergeServiceSlices(services.Compose, services.Override)
	} else {
		services.Compose = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.All = services.Compose
	}

	return services, nil
}

// IsEmpty returns whether the services are empty or not
func (s *Services) IsEmpty() bool {
	if len(s.All) == 0 || len(s.Compose) == 0 || len(s.Override) == 0 {
		return true
	}

	return false
}
