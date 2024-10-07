package entities

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vstring"
)

type Solution struct {
	// Header is the Hashcash header satisfying the Challenge requirements.
	Header string
}

func (s Solution) Validate() error {
	return yav.Chain("Header", s.Header, vstring.Required)
}
