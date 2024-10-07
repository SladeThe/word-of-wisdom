package impl

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/SladeThe/checked-go/must"
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vnumber"
	"github.com/catalinc/hashcash"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

type ChallengeConfig struct {
	// ZeroBitCount is the number of required zero bits in hash head.
	ZeroBitCount uint16 `default:"20"`
}

func (c ChallengeConfig) Validate() error {
	return yav.Chain(
		"ZeroBitCount", c.ZeroBitCount,
		vnumber.BetweenUint16(entities.ChallengeMinZeroBitCount, entities.ChallengeMaxZeroBitCount),
	)
}

type Challenge struct {
	cfg ChallengeConfig
}

var _ services.Challenge = (*Challenge)(nil)

func NewChallenge(cfg ChallengeConfig) (Challenge, error) {
	if errValidate := cfg.Validate(); errValidate != nil {
		return Challenge{}, errValidate
	}

	return Challenge{cfg: cfg}, nil
}

func (s Challenge) Accept(_ context.Context, id entities.ClientID) (entities.Challenge, error) {
	if errValidate := yav.Nested("id", id.Validate()); errValidate != nil {
		return entities.Challenge{}, errors.Join(services.ErrInvalidArguments, errValidate)
	}

	return entities.Challenge{ZeroBitCount: s.cfg.ZeroBitCount}, nil
}

func (s Challenge) Solve(
	_ context.Context,
	id entities.ClientID,
	challenge entities.Challenge,
	solution entities.Solution,
) error {
	errValidate := yav.Join(
		yav.Nested("id", id.Validate()),
		yav.Nested("challenge", challenge.Validate()),
		yav.Nested("solution", solution.Validate()),
	)

	if errValidate != nil {
		return errors.Join(services.ErrInvalidArguments, errValidate)
	}

	hash := hashcash.New(must.Uint16ToUint(challenge.ZeroBitCount), 8, "")
	if !hash.Check(solution.Header) || !strings.Contains(solution.Header, fmt.Sprintf(":%s:", id.String())) {
		return errors.Join(services.ErrInvalidArguments, services.ErrInvalidSolution)
	}

	return nil
}
