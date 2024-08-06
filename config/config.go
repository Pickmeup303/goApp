package config

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	AppPort string `yaml:"APP_PORT"`
	AppHost string `yaml:"APP_HOST"`
	AppURI  string `yaml:"APP_URI"`
}

func DefaultConfig() (*Config, error) {
	file, err := os.Open("config.yaml")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
