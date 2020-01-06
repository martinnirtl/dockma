package main

import (
	"github.com/spf13/viper"
	"github.com/AlecAivazis/survey/v2"
	"fmt"
)

// the questions to ask
var qs = []*survey.Question{
    {
        Name:     "Name",
        Prompt:   &survey.Input{Message: "What is your name?"},
        Validate: survey.Required,
        Transform: survey.Title,
    },
    {
        Name:     "SaveProfile",
        Prompt:   &survey.Confirm{Message: "Do you want to save your selection as a profile?"},
        Validate: survey.Required,
        Transform: survey.Title,
    },
    {
        Name:     "ProfileName",
        Prompt:   &survey.Input{Message: "Please enter a name for your profile?"},
        Validate: survey.Required,
        Transform: survey.Title,
    },
    {
        Name: "color",
        Prompt: &survey.Select{
            Message: "Choose a color:",
            Options: []string{"red", "blue", "green"},
            Default: "red",
        },
    },
    {
        Name: "age",
        Prompt: &survey.Input{Message: "How old are you?"},
		},
		{
			Name: "days",
			Prompt: &survey.MultiSelect{
					Message: "What days do you prefer:",
					Options: []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"},
					// Default: []string{"Sunday", "Monday"},
			},
		},
}

func main() {
	cmd.Execute()

	viper.SetConfigName("docker-compose")
	viper.AddConfigPath("/Users/martin/development/cube/dev-env-setup")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	serviceMaps := viper.GetStringMap("services")

	services := make([]string, 0, len(serviceMaps))

	for k := range serviceMaps {
		services = append(services, k)
	}

	for i := range services {
		fmt.Printf("%d.\t \"%s\"\n", i+1, services[i])
	}

	// the answers will be written to this struct
	answers := struct {
			Name          string                  // survey will match the question and field names
			SaveProfile   string                  // or you can tag fields to match a specific name
			ProfileName   string                  // or you can tag fields to match a specific name
			Age           int                     // if the types don't match, survey will convert it
			Days          []string
	}{}

	// perform the questions
	err = survey.Ask(qs, &answers, survey.WithIcons(func(icons *survey.IconSet) {
			icons.UnmarkedOption.Text = "◯"
			icons.MarkedOption.Text = "◉"
	}))
	if err != nil {
			fmt.Println(err.Error())
			return
	}

	// err = viper.WriteConfig()
	// if err != nil { // Handle errors reading the config file
	// 	panic(fmt.Errorf("Fatal error config file: %s \n", err))
	// }
}
