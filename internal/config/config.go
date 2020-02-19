package config

import (
	"errors"
	"fmt"
	"path/filepath"
	"sort"
	"time"

	"github.com/martinnirtl/dockma/internal/utils"
	"github.com/martinnirtl/dockma/pkg/dockercompose"
	"github.com/spf13/viper"
)

// NOTE viper gets initialized in commands/root.go.

// SaveConfig indicates whether config should be saved or not.
var SaveConfig bool

// message buffers for delayed/one-time saving.
var onWriteConfigError []error = make([]error, 0)
var onWriteConfigSuccess []string = make([]string, 0)

type env struct {
	name string
}

// Env provides an interface for easier access to more complex environment config.
type Env interface {
	GetName() string
	GetHomeDir() string
	IsRunning() bool
	SetUpdated() (time.Time, error)
	LastUpdate() (time.Duration, error)
	GetPullSetting() string
	GetProfileNames() []string
	HasProfile(name string) bool
	GetProfile(name string) (Profile, error)
	GetLatest() (Profile, error)
}

// Save sets the config to be saved at end of command execution. Respective message is printed after writing config.
func Save(success string, err error) {
	if success != "" {
		onWriteConfigSuccess = append(onWriteConfigSuccess, success)
	}

	if err != nil {
		onWriteConfigError = append(onWriteConfigError, err)
	}

	SaveConfig = true
}

// SaveNow writes the config to the file and returns previously cached success and error messages.
func SaveNow() (successMessages []string, errorMessages []error, err error) {
	SaveConfig = false

	successMessages = onWriteConfigSuccess
	errorMessages = onWriteConfigError

	onWriteConfigError = make([]error, 0)
	onWriteConfigSuccess = make([]string, 0)

	writeError := viper.WriteConfig()

	if writeError != nil {
		err = errors.New("Could not save changes")
	}

	return
}

// GetUsername returns the user's name.
func GetUsername() string {
	return viper.GetString("username")
}

// GetHomeDir returns the full path to dockma home dir.
func GetHomeDir() string {
	return viper.GetString("home")
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

// GetEnvNames returns sorted envs.
func GetEnvNames() (envs []string) {
	envsMap := viper.GetStringMap("envs")

	envs = make([]string, 0, len(envsMap))

	for env := range envsMap {
		envs = append(envs, env)
	}

	sort.Strings(envs)

	return
}

// GetActiveEnv returns the active environment.
func GetActiveEnv() Env {
	return &env{
		name: viper.GetString("active"),
	}
}

// SetActiveEnv returns the new active environment and the previous one.
func SetActiveEnv(new string) (newEnv Env, oldEnv Env) {
	old := viper.GetString("active")

	viper.Set("active", new)

	return &env{
			name: new,
		}, &env{
			name: old,
		}
}

// GetHideSubcommandOutputSetting returns hidesubcommandoutput flag.
func GetHideSubcommandOutputSetting() bool {
	return viper.GetBool("hidesubcommandoutput")
}

// GetEnv returns the named environment.
func GetEnv(name string) (Env, error) {
	if !utils.Includes(GetEnvNames(), name) {
		return nil, errors.New("No such environment")
	}

	return &env{
		name: viper.GetString("active"),
	}, nil
}

// GetName returns the name of env.
func (e *env) GetName() string {
	return e.name
}

// GetEnvHomeDir returns the full path to dockma home dir.
func (e *env) GetHomeDir() string {
	return viper.GetString(fmt.Sprintf("envs.%s.home", e.name))
}

// IsRunning returns current state of env.
func (e *env) IsRunning() bool {
	return viper.GetBool(fmt.Sprintf("envs.%s.running", e.name))
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
