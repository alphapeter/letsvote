package cfg

import (
	"fmt"
	"io/ioutil"
	"os"
	"encoding/json"
)

type Settings struct {
	Roots []Root `json: "root"`
	Binding string `json: binding`
}

type Root struct {
	Name string `json: "name"`
	Path string `json: "path"`
}

var settings *Settings = nil

func GetSettings() Settings {
	if settings == nil {
		settings = getSettingsFromFile()
	}
	return *settings
}

func GetPathForRoot(name string) (string, error) {
	s := GetSettings()
	for _, r := range s.Roots {
		if r.Name == name {
			return r.Path, nil
		}
	}
	return "", RootNotFoundError{name}
}

type RootNotFoundError struct {
	rootName string
}

func (e RootNotFoundError) Error() string {
	return "Root " + e.rootName + " not configured!"
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
