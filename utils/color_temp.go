package utils

import "math"

// TempToMired converts a color temperature to a mired
// value that Hue will accept
func TempToMired(temp uint16) uint16 {
	return uint16(math.Max(153.0, math.Min(math.Round(1000000.0/float64(temp)), 500.0)))
}

// MiredToTemp converts a mired value to a color temperature
func MiredToTemp(mired uint16) uint16 {
	return uint16(math.Max(2000.0, math.Min(roundToPrecision(1000000.0/float64(mired), -2), 6500.0)))
}

func roundToPrecision(x float64, precision int) float64 {
	p := precision * -1
	return math.Round(x/math.Pow10(p)) * math.Pow10(p)
}
