package entities

import (
	"github.com/SladeThe/yav"
	"github.com/SladeThe/yav/vstring"
)

type WordOfWisdom struct {
	Text string
}

func (w WordOfWisdom) Validate() error {
	return yav.Chain("Text", w.Text, vstring.Required)
}
