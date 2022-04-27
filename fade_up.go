// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"github.com/lucasb-eyer/go-colorful"
	"math"
)

// FadeUp will give values for a series of LEDs depending on a pos 0-100 in the animation phase. The
// animation itself will go from dim first row to gradually brightening the rows of LEDs to the
// specified LED RGB value
type FadeUp struct {
	// Rows is the number of rows of LEDs
	Rows uint

	// Cols is the number of LEDs horizontally per row
	Cols uint

	// Colour give the colour at the brightest points
	Colour colorful.Color
}

// CurrentLEDs returns the slice of LED values depending on the pos 0-1 which represents the
// position in the animation where 0 is unlit and 1 is fully lit. Any value between is gradually
// getting brighter row by row.
func (f *FadeUp) CurrentLEDs(pos float64) []LED {
	var ledSlice []LED
	h, s, v := f.Colour.Hsv()
	phases := float64(f.Rows) + 1.0
	k := 1.0 / ((1.0 / phases) * 2.0)
	var r float64
	var c float64
	var rowColour colorful.Color
	for r = 0; r < float64(f.Rows); r++ {
		b := -0.5 * r
		m := math.Max(math.Min((k*pos)+b, 1.0), 0.0)
		if m == 1.0 {
			rowColour = f.Colour // Workaround for rounding issues
		} else {
			rowColour = colorful.Hsv(h, s, v*m)
		}
		rowLED := LED{
			Colour:     rowColour,
			Brightness: MaxBrightness,
		}
		for c = 0; c < float64(f.Cols); c++ {
			ledSlice = append(ledSlice, rowLED)
		}
	}
	return ledSlice
}
