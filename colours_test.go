// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseColour_htmlShort(t *testing.T) {
	led, err := ParseColour("#feb")
	assert.NoError(t, err)
	assert.Equal(t, byte(0xff), led.Red)
	assert.Equal(t, byte(0xee), led.Green)
	assert.Equal(t, byte(0xbb), led.Blue)
}

func TestParseColour_htmlLong(t *testing.T) {
	led, err := ParseColour("#a87932")
	assert.NoError(t, err)
	assert.Equal(t, byte(168), led.Red)
	assert.Equal(t, byte(121), led.Green)
	assert.Equal(t, byte(50), led.Blue)
}

func TestParseColour_invalid(t *testing.T) {
	_, err := ParseColour("notacolour")
	assert.Error(t, err, "cannot parse colour \"notacolour\"")
}

func TestParseColour_outsideHexRange(t *testing.T) {
	_, err := ParseColour("#gggggg")
	assert.Error(t, err, "cannot parse colour \"#gggggg\"")
}
