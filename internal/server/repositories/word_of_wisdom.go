package repositories

import (
	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

type WordOfWisdom interface {
	OneRandom() (entities.WordOfWisdom, error)
}
