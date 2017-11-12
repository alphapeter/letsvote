package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Settings struct {
	Binding string `json: binding`
}

var settings *Settings = nil

func GetSettings() Settings {
	if settings == nil {
		settings = getSettingsFromFile()
	}
	return *settings
}

func getSettingsFromFile() *Settings {
	var settingsFilePath string
	args := os.Args[1:]
	if len(args) == 2 && args[0] == "--settings" {
		settingsFilePath = args[1]
	} else if len(args) == 0 {
		settingsFilePath = "settings.json"
	} else {
		fmt.Println("invalid arguments")
		fmt.Println("use --settings path-to-settings/settings.json")
		fmt.Println("default config path is ./settings.json")
		os.Exit(2)
	}

	settingsFile, err := ioutil.ReadFile(settingsFilePath)

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
	var settings Settings
	err = json.Unmarshal(settingsFile, &settings)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(13)
	}
	return &settings
}
