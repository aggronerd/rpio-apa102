package rpioapa102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func validLED() LED {
	return LEDFromIntRGBB(34, 223, 101, 20)
}

func TestLEDFromIntRGBB(t *testing.T) {
	led := validLED()
	r, g, b := led.Colour.RGB255()
	assert.Equal(t, byte(34), r)
	assert.Equal(t, byte(223), g)
	assert.Equal(t, byte(101), b)
	assert.Equal(t, byte(20), led.Brightness)
}

func TestLEDS_AsFrame(t *testing.T) {
	assert.Equal(t, []byte{0xF4, 0x65, 0xDF, 0x22}, validLED().AsFrame())
}

func TestLEDs_AsFrameInvalid(t *testing.T) {
	led := LEDFromIntRGBB(34, 223, 101, 32)
	assert.PanicsWithError(t, "brightness of 32 is invalid must be 0-31", func() {
		led.AsFrame()
	})
}

func TestLEDs_AsHTML(t *testing.T) {
	assert.Equal(t, "#22df65", validLED().AsHTML())
}

func TestLEDs_AsHTMLInvalid(t *testing.T) {
	led := LEDFromIntRGBB(34, 223, 101, 32)
	assert.PanicsWithError(t, "brightness of 32 is invalid must be 0-31", func() {
		led.AsHTML()
	})
}
