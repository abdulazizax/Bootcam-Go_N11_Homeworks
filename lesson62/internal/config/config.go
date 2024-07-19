package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Server ServerConfig
}

type ServerConfig struct {
	Port string
}

func (c *Config) Load() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	c.Server.Port = ":" + os.Getenv("PORT")

	return nil
}

func New() (*Config, error) {
	config := &Config{}
	err := config.Load()
	if err != nil {
		return nil, err
	}
	return config, nil
}
