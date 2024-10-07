package config

import (
	"github.com/SladeThe/yav"

	implServices "github.com/SladeThe/word-of-wisdom/internal/server/services/impl"
)

type Services struct {
	Challenge implServices.ChallengeConfig
}

func (cfg Services) Validate() error {
	return yav.Nested("Challenge", cfg.Challenge.Validate())
}
