package colors

import (
	"errors"
	"fmt"

	"github.com/YashdalfTheGray/colorcode"
)

// HueHSB represents a color in values that hue can understand
// This type also implements the fmt.Stringer interface
type HueHSB struct {
	H, S, B int
}

// NewHueHSB takes in HSB values and returns a new struct of
// type HueHSB
func NewHueHSB(h, s, b int) (HueHSB, error) {
	if h < 0 || h > 65535 || s < 0 || s > 254 || b < 1 || b > 254 {
		return HueHSB{}, errors.New("Invalid argument error")
	}

	return HueHSB{h, s, b}, nil
}

// String() returns the string representation of the HueHSB color
func (c HueHSB) String() string {
	return fmt.Sprintf("hue(%d, %d, %d)", c.H, c.S, c.B)
}

// ToHSV returns the HSV representation of this HueHSB color
func (c HueHSB) ToHSV() colorcode.HSV {
	return colorcode.HSV{H: (float64(c.H) / 65535.0) * 360.0, S: (float64(c.S) / 254.0) * 100.0, V: (float64(c.B) / 253.0) * 100.0}
}

// ToHSL returns the HSL representation of this HueHSB color
func (c HueHSB) ToHSL() colorcode.HSL {
	return c.ToHSV().ToHSL()
}

// ToRGB returns the RGB representation of this HueHSB color
func (c HueHSB) ToRGB() colorcode.RGB {
	return c.ToHSV().ToRGB()
}

// ToHexCode returns the Hex code representation of this HueHSB color
func (c HueHSB) ToHexCode() colorcode.HexCode {
	return c.ToHSV().ToRGB().ToHexCode()
}
