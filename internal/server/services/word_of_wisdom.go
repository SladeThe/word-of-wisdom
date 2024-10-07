package services

import (
	"context"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

type WordOfWisdom interface {
	OneRandom(ctx context.Context) (entities.WordOfWisdom, error)
}
