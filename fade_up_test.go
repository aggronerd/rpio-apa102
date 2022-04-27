// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stretchr/testify/assert"
	"testing"
)

func blackLED() LED {
	return LEDFromIntRGB(0, 0, 0)
}

func halfFaded(colour colorful.Color) LED {
	h, s, v := colour.Hsv()
	return LED{
		Colour:     colorful.Hsv(h, s, v/2.0),
		Brightness: MaxBrightness,
	}
}

func spanLED(led LED, n uint) []LED {
	var i uint
	var items []LED
	for i = 0; i < n; i++ {
		items = append(items, led)
	}
	return items
}

func TestFadeUp_CurrentLEDs_full(t *testing.T) {
	colour := colorful.Color{
		R: 0.5,
		G: 0.2,
		B: 0.7}
	half := halfFaded(colour)
	full := LED{
		Colour:     colour,
		Brightness: MaxBrightness,
	}
	example := FadeUp{
		Cols:   8,
		Rows:   4,
		Colour: colour,
	}

	// First phase
	assert.Equal(t,
		spanLED(blackLED(), 8*4),
		example.CurrentLEDs(0))

	// Second phase
	assert.Equal(t,
		append(spanLED(half, 8),
			spanLED(blackLED(), 8*3)...),
		example.CurrentLEDs(0.2))

	// 3rd phase
	assert.Equal(t,
		append(
			spanLED(full, 8),
			append(
				spanLED(half, 8),
				spanLED(blackLED(), 8*2)...)...),
		example.CurrentLEDs(0.4))

	// 4th phase
	assert.Equal(t,
		append(
			spanLED(full, 8*2),
			append(
				spanLED(half, 8),
				spanLED(blackLED(), 8*1)...)...),
		example.CurrentLEDs(0.6))

	// 5th phase
	assert.Equal(t,
		append(
			spanLED(full, 8*3),
			spanLED(half, 8)...),
		example.CurrentLEDs(0.8))

	// 6th phase
	assert.Equal(t,
		spanLED(full, 8*4),
		example.CurrentLEDs(1.0))
}

func TestFadeUp_CurrentLEDs_simple(t *testing.T) {
	colour := colorful.Color{
		R: 0.2,
		G: 0.3,
		B: 0.9}
	half := halfFaded(colour)
	full := LED{
		Colour:     colour,
		Brightness: MaxBrightness,
	}
	example := FadeUp{
		Cols:   1,
		Rows:   2,
		Colour: colour,
	}

	// First phase
	assert.Equal(t,
		spanLED(blackLED(), 2),
		example.CurrentLEDs(0))

	// Second phase
	assert.Equal(t,
		append(spanLED(half, 1),
			spanLED(blackLED(), 1)...),
		example.CurrentLEDs(1.0/3.0))

	// 3rd phase
	assert.Equal(t,
		append(
			spanLED(full, 1),
			spanLED(half, 1)...),
		example.CurrentLEDs(2.0/3.0))

	// 4th phase
	assert.Equal(t,
		spanLED(full, 2),
		example.CurrentLEDs(1.0))
}
