package config

import (
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
)

func LoadConfig() (*Config, error) {
	configPath := flag.String("config", "./config.yml", "config file path")
	flag.Parse()

	cfg := &Config{}

	err := cleanenv.ReadConfig(*configPath, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
