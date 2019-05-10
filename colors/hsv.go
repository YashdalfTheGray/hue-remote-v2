package colors

import (
	"errors"
	"fmt"
)

// HSV represents a color in the HSV color space. This also
// implements the fmt.Stringer interface
type HSV struct {
	H, S, V float64
}

// NewHSV takes values for H, S and V and returns an HSV type
func NewHSV(h, s, v float64) (HSV, error) {
	if h < 0 || h > 360 || s < 0 || s > 100 || v < 0 || v > 100 {
		return HSV{}, errors.New("Invalid argument error")
	}
	return HSV{h, s, v}, nil
}

// String returns a string representation of the HSV object
func (c HSV) String() string {
	return fmt.Sprintf("hsv(%.0f, %.0f%%, %.0f%%)", c.H, c.S, c.V)
}
