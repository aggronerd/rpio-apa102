// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpio_apa102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLEDS_asFrame(t *testing.T) {
	led := LED{
		Red:        34,
		Green:      223,
		Blue:       101,
		Brightness: 20,
	}

	assert.Equal(t, []byte{0xF4, 0x65, 0xDF, 0x22}, led.asFrame())
}

func TestLEDs_asFrameInvalid(t *testing.T) {
	led := LED{
		Red:        34,
		Green:      223,
		Blue:       101,
		Brightness: 32, // Outside of valid range
	}

	assert.PanicsWithError(t, "brightness of 32 is invalid must be 0-31", func() {
		led.asFrame()
	})
}
