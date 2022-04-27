package rpioapa102

import (
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stianeikeland/go-rpio/v4"
)

// MaxBrightness is the highest brightness value supported by the LEDs
const MaxBrightness = 31

// LEDController retains state and methods to control the APA102 on open SPI
type LEDController struct {
	pin rpio.SpiDev
}

// Finish should be called to close the SPI
func (c *LEDController) Finish() {
	rpio.SpiEnd(c.pin)
}

// Write takes a slice of LEDs in the order they are connected.
func (c *LEDController) Write(ledSlice []LED) {
	rpio.SpiTransmit([]byte{0x00, 0x00, 0x00, 0x00}...)
	for _, led := range ledSlice {
		rpio.SpiTransmit(led.AsFrame()...)
	}
	rpio.SpiTransmit([]byte{0xFF, 0xFF, 0xFF, 0xFF}...)
}

// NewLEDController will return controller for the LEDs
func NewLEDController(pin rpio.SpiDev) LEDController {
	controller := LEDController{pin: pin}

	if err := rpio.SpiBegin(pin); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0) // Select CE0 slave
	return controller
}

// LEDFromIntRGB returns LED objects whom color is based off RGB values 0-255, with max brightness
func LEDFromIntRGB(r uint, g uint, b uint) LED {
	return LEDFromIntRGBB(r, g, b, MaxBrightness)
}

// LEDFromIntRGBB returns LED objects whom color is based off RGB values 0-255
func LEDFromIntRGBB(r uint, g uint, b uint, brightness byte) LED {
	return LED{
		Colour: colorful.Color{
			R: float64(r) / 255.0,
			G: float64(g) / 255.0,
			B: float64(b) / 255.0,
		},
		Brightness: brightness,
	}
}

// LED is a representation of an RGB LED
type LED struct {
	//Colour for the LED
	Colour colorful.Color

	// Brightness 0-31 value for the LED. It is recommended NOT to use this due to do poor quality
	Brightness byte
}

func (l LED) validate() {
	if l.Brightness > MaxBrightness {
		panic(fmt.Errorf(
			"brightness of %d is invalid must be 0-%d", l.Brightness, MaxBrightness))
	}
}

// AsFrame returns the LED as a data frame to send to the APA102 via SPI
func (l LED) AsFrame() []byte {
	l.validate()
	r, g, b := l.Colour.RGB255()
	return []byte{0xE0 | l.Brightness, b, g, r}
}

// AsHTML returns the HTML hex code value for the colour of the LED
func (l LED) AsHTML() string {
	l.validate()
	return l.Colour.Hex()
}
