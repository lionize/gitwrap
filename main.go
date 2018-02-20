package main

import (
	"fmt"
	"os"

	"github.com/lionize/gitwrap/config"
	"github.com/lionize/gitwrap/git"
	"gopkg.in/AlecAivazis/survey.v1"
)

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(0)
	}

	script := os.Args[1]

	if script == "init" {
		cmdInit()
	} else {
		git.RunGitPassthrough()
	}
}

func cmdInit() {
	profiles := make(map[string]config.Profile)
	profileKeys := make([]string, 0)

	for _, profile := range config.Profiles() {
		key := fmt.Sprintf("%s <%s>", profile.Name, profile.Email)
		profiles[key] = profile
		profileKeys = append(profileKeys, key)
	}

	profileKeys = append(profileKeys, "Create a new user profile")
	response := ""
	prompt := &survey.Select{
		Message: "Choose a user profile to use for this repo:",
		Options: profileKeys,
	}
	survey.AskOne(prompt, &response, nil)

	if response == "Create a new user profile" {
		profile := config.Profile{}
		qs := []*survey.Question{
			{
				Name:   "name",
				Prompt: &survey.Input{Message: "Full name for profile:"},
			},
			{
				Name:   "email",
				Prompt: &survey.Input{Message: "Email for profile:"},
			},
		}

		err := survey.Ask(qs, &profile)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		config.CreateProfile(profile)
		git.RunGitInit(profile)
	} else {
		profile := profiles[response]
		git.RunGitInit(profile)
	}
}

func usage() {
	usageString := `usage: %s [script]

A tiny wrapper around git. If the provided script doesn't exist, all arguments will be passed through to git.

Available scripts:
	init`
	usageString = fmt.Sprintf(usageString, os.Args[0])
	fmt.Println(usageString)
}
