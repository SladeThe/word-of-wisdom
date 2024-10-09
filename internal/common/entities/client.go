package entities

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vnumber"
	"github.com/google/uuid"
)

type ClientID uuid.UUID

func NewClientID() ClientID {
	return ClientID(uuid.New())
}

func (id ClientID) Validate() error {
	return yav.Chain[[16]byte]("", id, yav.Required[[16]byte])
}

func (id ClientID) String() string {
	return uuid.UUID(id).String()
}

type Client struct {
	ID           ClientID
	ZeroBitCount uint16
}

func (c Client) Validate() error {
	return yav.Join(
		yav.Nested("ID", c.ID.Validate()),
		yav.Chain(
			"ZeroBitCount", c.ZeroBitCount,
			vnumber.BetweenUint16(ChallengeMinZeroBitCount, ChallengeMaxZeroBitCount),
		),
	)
}
