// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package main

import (
	"github.com/aggronerd/rpio-apa102"
	"github.com/stianeikeland/go-rpio/v4"
	"time"
)

const ledCount = 18

func main() {
	err := rpio.Open()
	if err != nil {
		panic(err)
	}

	controller := rpioapa102.NewLEDController(rpio.Spi0)
	defer controller.Finish()

	rgbLEDs := make([]rpioapa102.LED, ledCount)

	for i := 0; i < ledCount; i++ {
		rgbLEDs[i] = rpioapa102.LEDFromIntRGBB(47, 90, 94, 1)
	}

	controller.Write(rgbLEDs)
	time.Sleep(5 * time.Second)
}
