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
	r, g, b := led.Colour.RGB255()
	assert.Equal(t, byte(0xff), r)
	assert.Equal(t, byte(0xee), g)
	assert.Equal(t, byte(0xbb), b)
}

func TestParseColour_htmlLong(t *testing.T) {
	led, err := ParseColour("#a87932")
	assert.NoError(t, err)
	r, g, b := led.Colour.RGB255()
	assert.Equal(t, byte(168), r)
	assert.Equal(t, byte(121), g)
	assert.Equal(t, byte(50), b)
}

func TestParseColour_invalid(t *testing.T) {
	_, err := ParseColour("notacolour")
	assert.Error(t, err, "cannot parse colour \"notacolour\"")
}

func TestParseColour_outsideHexRange(t *testing.T) {
	_, err := ParseColour("#gggggg")
	assert.Error(t, err, "cannot parse colour \"#gggggg\"")
}
