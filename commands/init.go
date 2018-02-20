package commands

import (
	"fmt"

	"github.com/lionize/gitwrap/config"
	"github.com/lionize/gitwrap/git"
	"gopkg.in/AlecAivazis/survey.v1"
)

// Init initializes a git repository using interactive user input
var Init = Command{
	Name:        "init",
	Description: "description",
	Run: func(args []string) {
		profiles, keys := initProfiles()
		profile := userProfileSelect(profiles, keys)
		git.RunGitInit(profile)
	},
}

func initProfiles() (map[string]config.Profile, []string) {
	profiles := make(map[string]config.Profile)
	keys := make([]string, 0)

	for _, profile := range config.Profiles() {
		key := fmt.Sprintf("%s <%s>", profile.Name, profile.Email)
		profiles[key] = profile
		keys = append(keys, key)
	}

	keys = append(keys, "Create a new user profile")

	return profiles, keys
}

func userProfileSelect(profiles map[string]config.Profile, keys []string) config.Profile {
	response := ""

	prompt := &survey.Select{
		Message: "Choose a user profile to use for this repo:",
		Options: keys,
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

		survey.Ask(qs, &profile)
		config.CreateProfile(profile)
		return profile
	}

	profile := profiles[response]
	return profile
}
