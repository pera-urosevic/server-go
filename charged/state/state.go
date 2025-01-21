package state

import (
	"encoding/json"
	"errors"
	"os"

	"server/charged/log"
	"server/charged/types"
)

var _state types.State

func printConfig() {
	for _, item := range _state.Devices {
		name := item.Name
		log.Log("[DEVICE]", name, item.Low, "~", item.High)
	}
}

func configPath() (string, error) {
	path := os.Getenv("CHARGED_CONFIG_PATH")
	if path == "" {
		log.Log("config path env not found")
		return "", errors.New("config path env not found")
	}

	return path, nil
}

func configParse(data []byte) error {
	err := json.Unmarshal(data, &_state)
	if err != nil {
		log.Log(err)
		return err
	}

	return nil
}

func Load() error {
	path, err := configPath()
	if err != nil {
		log.Log(err)
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Log(err)
		return err
	}

	err = configParse(data)
	if err != nil {
		log.Log(err)
		return err
	}
	printConfig()

	return nil
}

func Time() int {
	return _state.Time
}

func Get() types.State {
	return _state
}
