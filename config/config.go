package config

import (
	"encoding/json"
	"os"
)

const ConfigFile = "config.json"

type Config struct {
	OpenAISecret string `json:"open_ai_secret"`
}

func LoadConfig() (Config, error) {
	var config Config
	file, err := os.Open(ConfigFile)
	if err != nil {
		return config, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		return config, err
	}
	return config, nil
}
