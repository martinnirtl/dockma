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

type env struct {
	name string
}

type Env interface {
	GetName() string
	GetHomeDir() string
	SetUpdated() (time.Time, error)
	LastUpdate() (time.Duration, error)
	GetPullSetting() string
	GetProfileNames() []string
	HasProfile(name string) bool
	GetProfile(name string) (Profile, error)
	GetLatest() (Profile, error)
}

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
func GetActiveEnv() Env {
	return &env{
		name: viper.GetString("active"),
	}
}

// GetFile returns full path of the given filename joined with dockma home dir.
func GetFile(filename string) string {
	path := viper.GetString("home")

	return filepath.Join(path, filename)
}

// GetLogfile returns full path to std dockma logfile.
func GetLogfile() string {
	filename := viper.GetString("logfile")

	return GetFile(filename)
}

// GetEnvs returns configured envs.
func GetEnvNames() (envs []string) {
	envsMap := viper.GetStringMap("envs")

	envs = make([]string, 0, len(envsMap))

	for env := range envsMap {
		envs = append(envs, env)
	}

	sort.Strings(envs)

	return
}

// GetName returns the name of env.
func (e *env) GetName() string {
	return e.name
}

// GetEnvHomeDir returns the full path to dockma home dir.
func (e *env) GetHomeDir() string {
	return viper.GetString(fmt.Sprintf("envs.%s.home", e.name))
}

// SetEnvUpdated updates update timestamp of env. Should be used when running git pull.
func (e *env) SetUpdated() (time.Time, error) {
	now := time.Now()

	viper.Set(fmt.Sprintf("envs.%s.home", e.name), now)

	return now, nil
}

// GetDurationPassedSinceLastUpdate returns duration passed since last git pull exec in given env.
func (e *env) LastUpdate() (time.Duration, error) {
	update := viper.GetTime(fmt.Sprintf("envs.%s.home", e.name))

	if update.IsZero() {
		return time.Duration(0), fmt.Errorf("No update done yet")
	}

	now := time.Now()

	return now.Sub(update), nil
}

// GetAutoPullSetting returns wether to run git pull or not.
func (e *env) GetPullSetting() string {
	return viper.GetString(fmt.Sprintf("envs.%s.pull", e.name))
}

// GetProfilesNames returns profile names for given env.
func (e *env) GetProfileNames() (profiles []string) {
	for profile := range viper.GetStringMap(fmt.Sprintf("envs.%s.profiles", e.name)) {
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
func (e *env) GetLatest() (profile Profile, err error) {
	profile = Profile{}

	services, err := dockercompose.GetServices(e.GetHomeDir())

	if err != nil {
		return
	}

	profile.Services = services.All

	profile.Selected = viper.GetStringSlice(fmt.Sprintf("envs.%s.latest", e.name))

	return
}

// GetProfile returns services for given profile.
func (e *env) GetProfile(profileName string) (profile Profile, err error) {
	profile = Profile{}

	services, err := dockercompose.GetServices(e.GetHomeDir())

	if err != nil {
		return
	}

	profile.Services = services.All

	profile.Selected = viper.GetStringSlice(fmt.Sprintf("envs.%s.profiles.%s", e.name, profileName))

	return
}

// HasProfileName tells if profile with name exists in env.
func (e *env) HasProfile(name string) bool {
	profile := viper.GetStringSlice(fmt.Sprintf("envs.%s.profiles.%s", e.name, name))

	if profile != nil {
		return true
	}

	return false
}
