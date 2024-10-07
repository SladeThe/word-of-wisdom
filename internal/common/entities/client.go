package entities

import (
	"github.com/SladeThe/yav"
	"github.com/google/uuid"
)

type ClientID uuid.UUID

func (id ClientID) Validate() error {
	return yav.Chain[[16]byte]("", id, yav.Required[[16]byte])
}

func (id ClientID) String() string {
	return uuid.UUID(id).String()
}
