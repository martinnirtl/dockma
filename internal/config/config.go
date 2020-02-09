package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/viper"
)

// NOTE viper gets initialized in commands/root.go.

// Save saves the config.
func Save() error {
	err := viper.WriteConfig()

	if err != nil {
		return errors.New("Could not save changes to dockma config")
	}

	return nil
}

// GetUsername returns the user's name.
func GetUsername() string {
	return viper.GetString("username")
}

// GetHomeDir returns the full path to dockma home dir.
func GetHomeDir() string {
	return viper.GetString("home")
}

// GetActiveEnv returns the name of the active environment.
func GetActiveEnv() string {
	return viper.GetString("active")
}

// GetDockmaFilepath returns full path of the given filename joined with dockma home dir.
func GetDockmaFilepath(filename string) string {
	path := viper.GetString("home")

	return filepath.Join(path, filename)
}

// GetLogfile returns full path to std dockma logfile.
func GetLogfile() string {
	filename := viper.GetString("logfile")

	return GetDockmaFilepath(filename)
}

// GetEnvs returns configured envs.
func GetEnvs() (envs []string) {
	envsMap := viper.GetStringMap("envs")

	envs = make([]string, 0, len(envsMap))

	for env := range envsMap {
		envs = append(envs, env)
	}

	sort.Strings(envs)

	return
}

// GetEnvHomeDir returns the full path to dockma home dir.
func GetEnvHomeDir(envName string) string {
	return viper.GetString(fmt.Sprintf("envs.%s.home", envName))
}

// SetEnvUpdated updates update timestamp of env. Should be used when running git pull.
func SetEnvUpdated(envName string) {
	viper.Set(fmt.Sprintf("envs.%s.home", envName), time.Now())
}

// GetDurationPassedSinceLastUpdate returns duration passed since last git pull exec in given env.
func GetDurationPassedSinceLastUpdate(envName string) (time.Duration, error) {
	update := viper.GetTime(fmt.Sprintf("envs.%s.home", envName))

	if update.IsZero() {
		return time.Duration(0), fmt.Errorf("No update done yet")
	}

	now := time.Now()

	return now.Sub(update), nil
}

// GetAutoPullSetting returns wether to run git pull or not.
func GetAutoPullSetting(envName string) string {
	return viper.GetString(fmt.Sprintf("envs.%s.pull", envName))
}

// GetProfilesNames returns profile names for given env.
func GetProfilesNames(env string) (profiles []string) {
	for profile := range viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", env)) {
		profiles = append(profiles, profile)
	}

	return
}

// Profile consists of selected and all services of env.
type Profile struct {
	Services []string
	Selected []string
}

// GetLatest returns special profile with latest chosen services.
func GetLatest(env string) (profile Profile, err error) {
	profile = Profile{}

	services, err := dockercompose.GetServices(GetEnvHomeDir(env))

	if err != nil {
		return
	}

	profile.Services = services.All

	profile.Selected = viper.GetStringSlice(fmt.Sprintf("envs.%s.latest", env))

	return
}

// GetProfile returns services for given profile.
func GetProfile(env string, profileName string) (profile Profile, err error) {
	profile = Profile{}

	services, err := dockercompose.GetServices(GetEnvHomeDir(env))

	if err != nil {
		return
	}

	profile.Services = services.All

	profile.Selected = viper.GetStringSlice(fmt.Sprintf("envs.%s.profiles.%s", env, profileName))

	return
}

// HasProfileName tells if profile with name exists in env.
func HasProfileName(env string, name string) bool {
	profile := viper.GetStringSlice(fmt.Sprintf("envs.%s.profiles.%s", env, name))

	if profile != nil {
		return true
	}

	return false
}
