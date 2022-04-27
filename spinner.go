package rpioapa102

import "math"

// SpinnerState includes the position of the currently lid LED in the ring and parameters
type SpinnerState struct {
	pos uint

	// LEDCount is the number of LEDs in the ring
	LEDCount uint

	// Spread is the number of LEDs that there will be a linear fall-off in brightness
	Spread byte

	// BaseLED give the colour at the brightest point in the ring
	BaseLED LED
}

// NewSpinner creates a SpinnerState with sensible parameters
func NewSpinner(ledCount uint, spread uint8, baseLED LED) SpinnerState {
	return SpinnerState{
		pos:      0,
		LEDCount: ledCount,
		Spread:   spread,
		BaseLED:  baseLED,
	}
}

// CurrentState returns the slice of LED values for the current position in the sequence
func (s *SpinnerState) CurrentState() []LED {
	var i uint
	ledState := make([]LED, s.LEDCount)
	fadeStep := float64(s.BaseLED.Brightness) / float64(s.Spread+1)
	spreadFloat := float64(s.Spread)
	for i = 0; i < s.LEDCount; i++ {
		off1 := math.Abs(float64(i) - float64(s.pos))
		off2 := math.Abs(float64(int(i)-int(s.LEDCount)) - float64(s.pos))
		off3 := math.Abs(float64(int(i)+int(s.LEDCount)) - float64(s.pos))
		offset := math.Min(math.Min(off1, off2), off3)
		led := s.BaseLED
		if offset <= spreadFloat {
			led.Brightness = byte(math.Round(float64(s.BaseLED.Brightness) - (offset * fadeStep)))
		} else {
			led.Brightness = 0
		}
		ledState[i] = led
	}
	return ledState
}

// Next transitions the spinner to the next state
func (s *SpinnerState) Next() {
	if s.pos != (s.LEDCount - 1) {
		s.pos++
	} else {
		s.pos = 0
	}
}
