package configs

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Addresss string `yaml:"address"`
	Port     string `yaml:"port"`
	DSN      string `yaml:"dsn"`
	LogLevel string `yaml:"log_level"`
}

func NewConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}

	err = yaml.NewDecoder(file).Decode(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
