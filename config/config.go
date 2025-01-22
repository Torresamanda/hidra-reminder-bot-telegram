package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	TelegramToken string `json:"telegram_token"`
}

func LoadConfig() (Config, error) {
	file, err := os.Open("config.json")
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var cfg Config
	err = decoder.Decode(&cfg)

	return cfg, err
}
