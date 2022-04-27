// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package main

import (
	rpioapa102 "github.com/aggronerd/rpio-apa102"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

const durationMinutes = 2
const sleepMilliseconds = 100

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

	colour, _ := colorful.Hex("#70a9ff")
	fadeUpState, _ := rpioapa102.NewFadeUp(4, 8, colour, true)

	startTime := time.Now()
	pos := 0.0
	for pos < 1.0 {
		delta := time.Now().Sub(startTime)
		pos = float64(delta) / float64(durationMinutes*time.Minute)
		controller.Write(fadeUpState.CurrentLEDs(pos))
		time.Sleep(sleepMilliseconds * time.Millisecond)
	}
}
