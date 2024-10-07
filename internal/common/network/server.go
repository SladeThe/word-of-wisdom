package network

import (
	"github.com/SladeThe/word-of-wisdom/internal/common/entities"
)

type Server interface {
	WriteClientID(id entities.ClientID) error
	WriteChallenge(challenge entities.Challenge) error
	ReadSolution() (entities.Solution, error)
	WriteWordOfWisdom(word entities.WordOfWisdom) error
}
