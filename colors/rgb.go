package colors

import (
	"fmt"
	"math"
)

// RGB represents a traditional 24-bit color, having 8 bits
// for each of red, green, and blue. Implements fmt.Stringer
type RGB struct {
	R, G, B uint8
}

// NewRGB returns a new RGB object constructured out of the
// values given to the constructor function
func NewRGB(r, g, b uint8) RGB {
	return RGB{r, g, b}
}

// String returns a string representation of the RGB object
func (c RGB) String() string {
	return fmt.Sprintf("rgb(%d, %d, %d)", c.R, c.G, c.B)
}

// ToHSL converts an RGB color to the HSL representation
func (c RGB) ToHSL() HSL {
	var h, s, l float64

	fracR := float64(c.R) / 255.0
	fracG := float64(c.G) / 255.0
	fracB := float64(c.B) / 255.0

	max := math.Max(math.Max(fracR, fracG), fracB)
	min := math.Min(math.Min(fracR, fracG), fracB)

	// Luminosity is the average of the max and min rgb color intensities.
	l = (max + min) / 2

	// saturation
	delta := max - min
	if delta == 0 {
		// it's gray
		return HSL{0, 0, l * 100}
	}

	// it's not gray
	if l < 0.5 {
		s = delta / (max + min)
	} else {
		s = delta / (2 - max - min)
	}

	// hue
	r2 := (((max - fracR) / 6) + (delta / 2)) / delta
	g2 := (((max - fracG) / 6) + (delta / 2)) / delta
	b2 := (((max - fracB) / 6) + (delta / 2)) / delta
	switch {
	case fracR == max:
		h = b2 - g2
	case fracG == max:
		h = (1.0 / 3.0) + r2 - b2
	case fracB == max:
		h = (2.0 / 3.0) + g2 - r2
	}

	// fix wraparounds
	switch {
	case h < 0:
		h++
	case h > 1:
		h--
	}

	return HSL{h * 360, s * 100, l * 100}
}

// ToHSV returns the HSV representation of the RGB color
func (c RGB) ToHSV() (HSV, error) {
	var rdif, gdif, bdif, h, s, v float64

	fracR := float64(c.R) / 255.0
	fracG := float64(c.G) / 255.0
	fracB := float64(c.B) / 255.0

	v = math.Max(fracR, math.Max(fracG, fracB))
	diff := v - math.Min(fracR, math.Min(fracG, fracB))

	diffChannel := func(c float64) float64 {
		return (v-c)/6/diff + 1/2
	}

	if diff == 0 {
		h = 0
		s = 0
	} else {
		s = diff / v
		rdif = diffChannel(fracR)
		gdif = diffChannel(fracG)
		bdif = diffChannel(fracB)

		if fracR == v {
			h = bdif - gdif
		} else if fracG == v {
			h = (1 / 3) + rdif - bdif
		} else if fracB == v {
			h = (2 / 3) + gdif - rdif
		}

		switch {
		case h < 0:
			h++
		case h > 1:
			h--
		}
	}

	return NewHSV(h*360.0, s*100.0, v*100.0)
}

// ToHexCode returns a Hex color string equivalent of the the
// RGB color
func (c RGB) ToHexCode() HexCode {
	return HexCode(fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B))
}
