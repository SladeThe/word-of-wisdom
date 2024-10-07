package config

import (
	"github.com/SladeThe/yav"
	"github.com/jinzhu/configor"

	"github.com/SladeThe/word-of-wisdom/internal/client"
)

type Config struct {
	Client client.Config
}

func (cfg Config) Validate() error {
	return yav.Nested("Client", cfg.Client.Validate())
}

func New() (Config, error) {
	var cfg Config

	if errLoad := configor.Load(&cfg); errLoad != nil {
		return cfg, errLoad
	}

	return cfg, cfg.Validate()
}
