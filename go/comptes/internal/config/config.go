package config

import (
	"os"

	"go.yaml.in/yaml/v3"
)

type Config struct {
	Db struct {
		Connection string `yaml:"connection"`
	} `yaml:"db"`
}

func LoadConfig(path string) (*Config, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
