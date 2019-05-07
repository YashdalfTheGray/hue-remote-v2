package colors

import (
	"errors"
	"fmt"
)

// HSL represents a color in the HSL coordinate system,
// having 8 bits for each of hue, saturation, and luminosity.
// Implements fmt.Stringer
type HSL struct {
	H, S, L float64
}

// NewHSL returns a new HSL object constructured out of the
// values given to the constructor function
func NewHSL(h, s, l float64) (HSL, error) {
	if h < 0 || h > 360 || s < 0 || s > 100 || l < 0 || l > 100 {
		return HSL{}, errors.New("Invalid argument error")
	}
	return HSL{h, s, l}, nil
}

// String returns a string representation of the HSL object
func (c HSL) String() string {
	return fmt.Sprintf("hsl(%.0f, %.0f%%, %.0f%%)", c.H, c.S, c.L)
}

// ToRGB converts an HSL color to its RGB representation
func (c HSL) ToRGB() RGB {
	h := c.H / 360.0
	s := c.S / 100.0
	l := c.L / 100.0

	if s == 0 {
		// it's gray
		return RGB{uint8(l * 255.0), uint8(l * 255.0), uint8(l * 255.0)}
	}

	var v1, v2 float64
	if l < 0.5 {
		v2 = l * (1 + s)
	} else {
		v2 = (l + s) - (s * l)
	}

	v1 = 2*l - v2

	r := hueToRGB(v1, v2, h+(1.0/3.0))
	g := hueToRGB(v1, v2, h)
	b := hueToRGB(v1, v2, h-(1.0/3.0))

	return RGB{uint8(r * 255.0), uint8(g * 255.0), uint8(b * 255.0)}
}

// ToHexCode returns a hex color string equivalent of the
// HSL color
func (c HSL) ToHexCode() HexCode {
	return c.ToRGB().ToHexCode()
}

func hueToRGB(v1, v2, h float64) float64 {
	if h < 0 {
		h++
	}
	if h > 1 {
		h--
	}
	switch {
	case 6*h < 1:
		return (v1 + (v2-v1)*6*h)
	case 2*h < 1:
		return v2
	case 3*h < 2:
		return v1 + (v2-v1)*((2.0/3.0)-h)*6
	}
	return v1
}
