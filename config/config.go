package config

import (
	"encoding/json"
	"log"
	"os"
)

type DataBaseConfig struct {
	User string `json:"user"`
	Pass string `json:"pass"`
	Host string `json:"host"`
	Port string `json:"port"`
	DB   string `json:"db"`
}

type Config struct {
	DataBaseConfig DataBaseConfig `json:"data_base_config"`
}

var loadedConfig *Config

func GetConfig() *Config {
	if loadedConfig == nil {
		config := &Config{}
		raw, err := os.ReadFile("public/config/conf.json")
		if err != nil {
			log.Fatalf("Error occured while reading config")
		}
		json.Unmarshal(raw, config)
		loadedConfig = config
		log.Printf("Config loaded: %v\n", loadedConfig)
	}
	return loadedConfig
}
