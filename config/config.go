package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// Profile struct represents user profile information.
type Profile struct {
	Name  string
	Email string
}

type configJSON struct {
	profiles []Profile
}

var config = configJSON{
	profiles: []Profile{},
}

func init() {
	cfgFile := getConfigPath()

	_, err := os.Stat(cfgFile)

	if err != nil {
		createConfig()
	}

	jsonBytes, err := ioutil.ReadFile(cfgFile)

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(jsonBytes, &config)

	if err != nil {
		fmt.Println("ERROR: ", err.Error())
		os.Exit(1)
	}

	fmt.Println(config)
}

// Profiles returns the user profiles stored
// in the user's config file
func Profiles() []Profile {
	return config.profiles
}

func createConfig() {
	os.MkdirAll(getConfigDir(), os.ModePerm)
	saveConfig()
}

func saveConfig() error {
	jsn, err := json.MarshalIndent(config, "", "\t")

	if err != nil {
		return err
	}

	return ioutil.WriteFile(getConfigPath(), jsn, 0644)
}

func getConfigDir() string {
	home := os.Getenv("HOME")

	return filepath.Join(home, ".config", "gitwrap")
}

func getConfigPath() string {
	return filepath.Join(getConfigDir(), "config.json")
}
