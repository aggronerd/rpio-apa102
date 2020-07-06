// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpio_apa102

import (
	"fmt"
	"github.com/stianeikeland/go-rpio/v4"
)

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

// Setup will prepare GPIO
func Setup(pin rpio.SpiDev) {
	if err := rpio.Open(); err != nil {
		panic(err)
	}

	if err := rpio.SpiBegin(pin); err != nil {
		panic(err)
	}

	rpio.SpiChipSelect(0) // Select CE0 slave
}

// SetLEDs takes a slice of LEDs in the order they are connected.
func SetLEDs(leds []LED) {
	rpio.SpiTransmit([]byte{0x00, 0x00, 0x00, 0x00}...)
	for _, led := range leds {
		rpio.SpiTransmit(led.asFrame()...)
	}
	rpio.SpiTransmit([]byte{0xFF, 0xFF, 0xFF, 0xFF}...)
}
