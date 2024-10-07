package network

import (
	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

type Client interface {
	ReadClientID() (entities.ClientID, error)
	ReadChallenge() (entities.Challenge, error)
	WriteSolution(solution entities.Solution) error
	ReadWordOfWisdom() (entities.WordOfWisdom, error)
}
