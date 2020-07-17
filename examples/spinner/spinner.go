// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package main

import (
	"github.com/aggronerd/rpio-apa102"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

const pauseMilliseconds = 50

/*
 * A very basic example that will create a spinner effect based on a set number of LEDs in a ring.
 * The base colour and brightness is set and Spinner will generate the various states in sequence
 * to give a smooth spinning effect.
 */
func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	controller := rpioapa102.NewLEDController(rpio.Spi0)
	defer controller.Finish()

	led := rpioapa102.LED{
		Red:        178,
		Green:      255,
		Blue:       0,
		Brightness: 15,
	}

	spinner := rpioapa102.NewSpinner(18, 4, led)
	for {
		// Write the new spinner state
		controller.Write(spinner.CurrentState())
		time.Sleep(pauseMilliseconds * time.Millisecond)
		spinner.Next()
	}
}
