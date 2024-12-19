package flick

import (
	"testing"
)

func TestConfigPath(t *testing.T) {
	path, err := configPath()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(path) == 0 {
		t.Errorf("path is empty")
	}
}

func TestConfigLoad(t *testing.T) {
	err := configLoad()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	if len(config.Api.Key) == 0 {
		t.Errorf("api key is missing")
	}
	if len(config.Api.Secret) == 0 {
		t.Errorf("api secret is missing")
	}
	if len(config.Access.Token) < 1 {
		t.Errorf("access token is missing")
	}
	if len(config.Access.Secret) < 1 {
		t.Errorf("access secret is missing")
	}
}

func TestConfigSave(t *testing.T) {
	err := configLoad()
	if err != nil {
		t.Errorf("err: %v", err)
	}
	err = configSave()
	if err != nil {
		t.Errorf("err: %v", err)
	}
}
