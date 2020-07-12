// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

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
		rpio.SpiTransmit(led.asFrame()...)
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

// LED is a representation of an RGB LED
type LED struct {
	Red   byte
	Green byte
	Blue  byte

	// Brightness 0-31 value for the LED. It is recommended NOT to use this due to do poor quality
	Brightness byte
}

func (l LED) asFrame() []byte {
	if l.Brightness > 31 {
		panic(fmt.Errorf("brightness of %d is invalid must be 0-31", l.Brightness))
	}

	return []byte{0xE0 | l.Brightness, l.Blue, l.Green, l.Red}
}
