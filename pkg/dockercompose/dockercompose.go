package dockercompose

import (
	"fmt"

	"github.com/spf13/viper"
)

type Services struct {
	All      []string
	Base     []string
	Override []string
}

func GetDockerCompose(filepath string, override bool) (*viper.Viper, error) {
	configName := "docker-compose"
	if override {
		configName = "docker-compose.override"
	}

	temp := viper.New()
	temp.SetConfigName(configName)
	temp.SetConfigType("yaml")
	temp.AddConfigPath(filepath)

	readError := temp.ReadInConfig()

	if readError != nil {
		return nil, fmt.Errorf("could not read docker-compose?.override file")
	}

	return temp, nil
}

func GetVersion(filepath string) string {
	dockercompose, err := GetDockerCompose(filepath, false)

	if err != nil {
		return ""
	}

	return dockercompose.GetString("version")
}

func GetServices(filepath string) (services Services, err error) {
	dockercompose, err := GetDockerCompose(filepath, false)

	services = Services{}

	if err != nil {
		return services, err
	}

	dockercomposeOverride, err := GetDockerCompose(filepath, true)

	if err == nil {
		services.Base = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.Override = getServicesFromStringMap(dockercomposeOverride.GetStringMap("services"))

		services.All = mergeServiceSlices(services.Base, services.Override)
	} else {
		services.Base = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.All = services.Base
	}

	return services, nil
}

func GetEnvironment(filepath string) map[string]string {
	return nil
}
