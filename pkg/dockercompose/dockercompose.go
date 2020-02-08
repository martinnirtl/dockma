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

var servicesChache map[string]Services = make(map[string]Services)

func isEmpty(services Services) bool {
	if len(services.All) == 0 || len(services.Base) == 0 || len(services.Override) == 0 {
		return true
	}

	return false
}

func getDockerCompose(filepath string, override bool) (*viper.Viper, error) {
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
		return nil, fmt.Errorf("Could not read %s file", fileName)
	}

	return temp, nil
}

func GetVersion(filepath string) string {
	dockercompose, err := getDockerCompose(filepath, false)

	if err != nil {
		return ""
	}

	return dockercompose.GetString("version")
}

func GetServices(filepath string) (services Services, err error) {
	services = servicesChache[filepath]

	if !isEmpty(services) {
		return
	}

	dockercompose, err := getDockerCompose(filepath, false)

	services = Services{}

	if err != nil {
		return services, err
	}

	dockercomposeOverride, err := getDockerCompose(filepath, true)

	if err == nil {
		services.Base = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.Override = getServicesFromStringMap(dockercomposeOverride.GetStringMap("services"))

		services.All = mergeServiceSlices(services.Base, services.Override)
	} else {
		services.Base = getServicesFromStringMap(dockercompose.GetStringMap("services"))
		services.All = services.Base
	}

	servicesChache[filepath] = services

	return services, nil
}

func GetEnvironment(filepath string) map[string]string {
	return nil
}
