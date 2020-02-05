package config

import "github.com/spf13/viper"

var homeDir = viper.GetString("home")

// GetActiveEnv returns id/name of the currently active environment
func GetActiveEnv() string {
	return viper.GetString("active")
}
