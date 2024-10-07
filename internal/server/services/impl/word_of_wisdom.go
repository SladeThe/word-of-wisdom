package impl

import (
	"context"

	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
	"github.com/SladeThe/word-of-wisdom/internal/server/repositories"
	"github.com/SladeThe/word-of-wisdom/internal/server/services"
)

type WordOfWisdom struct {
}

var _ services.WordOfWisdom = (*WordOfWisdom)(nil)

func NewWordOfWisdom() WordOfWisdom {
	return WordOfWisdom{}
}

func (s WordOfWisdom) OneRandom(ctx context.Context) (entities.WordOfWisdom, error) {
	rr := repositories.Must(ctx)
	word, errRandom := rr.WordOfWisdom.OneRandom()
	return word, handleError(errRandom, "failed getting random word")
}
