package utils

import "math"

// TempToMired converts a color temperature to a mired
// value that Hue will accept
func TempToMired(temp uint16) uint16 {
	return uint16(math.Max(153, math.Min(math.Round(1000000/float64(temp)), 500)))
}
