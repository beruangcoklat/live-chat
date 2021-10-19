package config

import (
	"encoding/json"
	"os"
)

var cfg Config

func Init(filepath string) error {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		return err
	}

	return nil
}

func GetConfig() Config {
	return cfg
}
