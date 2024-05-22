package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	Token    string `yaml:"token" env-required:"true"`
	LogLevel string `yaml:"log_level" env-default:"info"`
}

func LoadConfig() (*Config, error) {
	configPath := os.Getenv("CONFIG_PATH")

	cfg := new(Config)

	err := cleanenv.ReadConfig(configPath, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
