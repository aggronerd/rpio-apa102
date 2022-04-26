// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"github.com/lucasb-eyer/go-colorful"
)

// ParseColour attempts to parse string hexadecimal representations of colours. eg. "#fff" or
//"#FEFEFE"
func ParseColour(colourStr string) (*LED, error) {
	colour, err := colorful.Hex(colourStr)
	if err != nil {
		return nil, err
	}
	return &LED{
		Colour:     colour,
		Brightness: MaxBrightness,
	}, nil
}
