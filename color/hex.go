package color

import "fmt"

// HexCode represents an HTML color string
type HexCode string

// ToRGB converts a HexCode string color into its RGB equivalent
func (cs HexCode) ToRGB() (RGB, error) {
	var r, g, b uint8

	n, err := fmt.Sscanf(string(cs), "#%2x%2x%2x", &r, &g, &b)
	if err != nil || n != 3 {
		return RGB{}, err
	}

	return RGB{r, g, b}, nil
}

// ToHSL converts a HexCode string color to its HSL equivalent
func (cs HexCode) ToHSL() (HSL, error) {
	rgb, err := cs.ToRGB()
	if err != nil {
		return HSL{}, err
	}

	return rgb.ToHSL(), nil
}
