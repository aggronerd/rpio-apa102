// Copyright 2019 Gregory Doran <greg@gregorydoran.co.uk>.
// All rights reserved.

package rpioapa102

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var htmlColourRegex *regexp.Regexp

func init() {
	htmlColourRegex = regexp.MustCompile(
		"^#([0-9a-fA-Z]{1,2})([0-9a-fA-Z]{1,2})([0-9a-fA-Z]{1,2})$")
}

// ParseColour attempts to parse string hexadecimal representations of colours. eg. "#fff" or
//"#FEFEFE"
func ParseColour(colourStr string) (*LED, error) {
	matches := htmlColourRegex.FindAllStringSubmatch(colourStr, -1)
	if matches != nil {
		var parts [3]uint64
		for i := 1; i < 4; i++ {
			if len(matches[0][i]) == 1 {
				matches[0][i] = matches[0][i] + matches[0][i]
			}
			parts[i-1], _ = strconv.ParseUint(strings.ToLower(matches[0][i]), 16, 16)
		}
		return &LED{
			Red:        byte(parts[0]),
			Green:      byte(parts[1]),
			Blue:       byte(parts[2]),
			Brightness: MaxBrightness,
		}, nil
	}
	return nil, fmt.Errorf("cannot parse colour \"%s\"", colourStr)
}
