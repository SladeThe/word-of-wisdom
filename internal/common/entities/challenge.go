package entities

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vnumber"
)

const (
	ChallengeMinZeroBitCount = 4
	ChallengeMaxZeroBitCount = 156
)

type Challenge struct {
	// ZeroBitCount is the number of required zero bits in hash head.
	ZeroBitCount uint16
}

func (c Challenge) Validate() error {
	return yav.Chain(
		"ZeroBitCount", c.ZeroBitCount,
		vnumber.BetweenUint16(ChallengeMinZeroBitCount, ChallengeMaxZeroBitCount),
	)
}
