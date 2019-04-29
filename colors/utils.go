package colors

func toFractional(v uint8) float64 {
	return float64(v) / 255.0
}

func toScaled(v float64) uint8 {
	return uint8(v * 255.0)
}
