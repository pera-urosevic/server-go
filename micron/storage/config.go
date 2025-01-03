package storage

import (
	"encoding/json"
	"errors"
	"os"

	"server/micron/log"
	"server/micron/types"
)

var Config types.Config

func PrintConfig() {
	for _, item := range Config.Monitor {
		name := item.Name
		if !item.Enabled {
			name = "# " + name
		}
		log.Log("[MONITOR]", name)
	}

	for _, item := range Config.Daily {
		name := item.Name
		if !item.Enabled {
			name = "# " + name
		}
		log.Log("[DAILY]", name)
	}

	for _, item := range Config.Weekly {
		name := item.Name
		if !item.Enabled {
			name = "# " + name
		}
		log.Log("[WEEKLY]", name)
	}
}

func configPath() (string, error) {
	path := os.Getenv("MICRON_CONFIG_PATH")
	if path == "" {
		return "", errors.New("config path env not found")
	}
	return path, nil
}

func ConfigLoad() error {
	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = ConfigParse(data)
	if err != nil {
		return err
	}
	PrintConfig()

	return nil
}

func ConfigParse(data []byte) error {
	err := json.Unmarshal(data, &Config)
	if err != nil {
		return err
	}

	Config.Changed = false
	return nil
}

func ConfigSave() error {
	if !Config.Changed {
		return nil
	}

	path, err := configPath()
	if err != nil {
		return err
	}

	data, err := json.MarshalIndent(Config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	Config.Changed = false
	return nil
}
