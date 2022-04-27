package rpioapa102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func baseLED() LED {
	return LEDFromIntRGB(
		45, 21, 92)
}

func TestSpinnerState_CurrentState_Bottom(t *testing.T) {
	state := SpinnerState{
		pos:      2,
		Spread:   3,
		LEDCount: 9,
		BaseLED:  baseLED(),
	}

	returned := state.CurrentState()
	brightnessValues := make([]byte, len(returned))
	for i, led := range returned {
		brightnessValues[i] = led.Brightness
	}

	assert.EqualValues(t,
		[]byte{16, 23, 31, 23, 16, 8, 0, 0, 8},
		brightnessValues)
}

func TestSpinnerState_CurrentState_Top(t *testing.T) {
	state := SpinnerState{
		pos:      14,
		Spread:   4,
		LEDCount: 15,
		BaseLED:  baseLED(),
	}

	returned := state.CurrentState()
	brightnessValues := make([]byte, len(returned))
	for i, led := range returned {
		brightnessValues[i] = led.Brightness
	}

	assert.EqualValues(t,
		[]byte{25, 19, 12, 6, 0, 0, 0, 0, 0, 0, 6, 12, 19, 25, 31},
		brightnessValues)
}

func TestSpinnerState_Next_first(t *testing.T) {
	state := SpinnerState{
		pos:      0,
		LEDCount: 9,
	}

	state.Next()
	assert.EqualValues(t, 1, state.pos)
}

func TestSpinnerState_Next_last(t *testing.T) {
	state := SpinnerState{
		pos:      8,
		LEDCount: 9,
	}

	state.Next()
	assert.EqualValues(t, 0, state.pos)
}
