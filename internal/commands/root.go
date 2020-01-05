package commands

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir" // alias import
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func Execute() {
	rootCmd := &cobra.Command{
		Use:   "dockma",
		Short: "Dockma is bringing your docker-compose game to the next level.",
		Long:  `A fast and flexible CLI tool to boost your productivity during development with docker containers built with Go. Full documentation is available at https://dockma.dev`,
	}

	rootCmd.AddCommand(GetVersionCommand())
	rootCmd.AddCommand(GetEnvironmentsCommand())

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cobra.yaml)")
	// rootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "author name for copyright attribution")
	// rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "name of license for the project")
	// rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	// viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	// viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	// viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	// viper.SetDefault("license", "apache")
}

func initConfig() {
	viper.SetConfigName(".dockma")

	home, err := homedir.Dir()
	if err != nil {
		fmt.Println("An error occured.")
	}

	viper.AddConfigPath(".")
	viper.AddConfigPath(home)

	// check also BindEnv
	// viper.SetEnvPrefix("DOCKMA_")
	// viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		// TODO log config file used depending on flag

		// fmt.Println("Using config file: ", viper.ConfigFileUsed())
	}
}
