package config

import (
	"os"

	"github.com/adrg/xdg"
	"gopkg.in/yaml.v3"
)

type Config struct {
	ApiKey string `yaml:"apikey"`
}

const configFilename = "gandi-livedns-cli/config.yaml"

func LoadApiKey() (string, error) {
	apiKey := os.Getenv("GANDI_V5_APIKEY")
	if apiKey != "" {
		return apiKey, nil
	}

	// read config file.
	path, err := xdg.SearchConfigFile(configFilename)
	if err != nil {
		return "", err
	}
	f, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	var t Config
	err = yaml.Unmarshal(f, &t)
	if err != nil {
		return "", err
	}
	return t.ApiKey, nil
}

func SaveApiKey(apikey string) (string, error) {
	var t Config
	t.ApiKey = apikey

	path, err := xdg.ConfigFile(configFilename)
	if err != nil {
		return "", err
	}

	b, err := yaml.Marshal(t)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(path, b, 0644)
	if err != nil {
		return "", err
	}
	return path, nil
}

func DeleteApiKey() error {
	path, err := xdg.SearchConfigFile(configFilename)
	if err != nil {
		return err
	}
	return os.Remove(path)
}
