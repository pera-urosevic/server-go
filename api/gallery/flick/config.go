package flick

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	Api struct {
		Key    string `json:"key"`
		Secret string `json:"secret"`
	} `json:"api"`
	Access struct {
		Token  string `json:"token"`
		Secret string `json:"secret"`
	} `json:"access"`
}

var config Configuration

func configPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s\\flick.json", dirname), nil
}

func hasFlickrAccess() bool {
	return len(config.Access.Token) > 0 && len(config.Access.Secret) > 0
}

func configLoad() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return err
	}

	return nil
}

func configSave() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
