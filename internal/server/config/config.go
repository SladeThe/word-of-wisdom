package config

import (
	"context"
	"errors"

	"github.com/SladeThe/yav"
	"github.com/jinzhu/configor"

	"github.com/SladeThe/word-of-wisdom/internal/server"
)

type contextKey string

var (
	ErrMissingConfig = errors.New("missing config")

	configKey = contextKey("wow.config")
)

type Config struct {
	Server   server.Config
	Services Services
}

func (cfg Config) Validate() error {
	return yav.Join(
		yav.Nested("Server", cfg.Server.Validate()),
		yav.Nested("Services", cfg.Services.Validate()),
	)
}

func New() (Config, error) {
	var cfg Config

	if errLoad := configor.Load(&cfg); errLoad != nil {
		return cfg, errLoad
	}

	return cfg, cfg.Validate()
}

func Set(ctx context.Context, cfg Config) context.Context {
	return context.WithValue(ctx, configKey, cfg)
}

func Get(ctx context.Context) (Config, error) {
	config, ok := ctx.Value(configKey).(Config)
	if !ok {
		return Config{}, ErrMissingConfig
	}
	return config, nil
}

func Must(ctx context.Context) Config {
	config, err := Get(ctx)
	if err != nil {
		panic(err)
	}
	return config
}
