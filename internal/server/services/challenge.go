package services

import (
	"context"
	"errors"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

var ErrInvalidSolution = errors.New("invalid solution")

type Challenge interface {
	Accept(ctx context.Context, id entities.ClientID) (entities.Challenge, error)
	Solve(ctx context.Context, id entities.ClientID, challenge entities.Challenge, solution entities.Solution) error
}
