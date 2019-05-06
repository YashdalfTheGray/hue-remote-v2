package utils

import (
	"regexp"
)

// ValidateColorString returns true for valid hex color strings
// that start with # and have 6 hex digits
func ValidateColorString(cs string) bool {
	if result, err := regexp.MatchString(`^#[A-Fa-f0-9]{6}$`, cs); err == nil {
		return result
	}
	return false
}

// ValidateColorArray returns true for a valid 3-pair array of
// octets representing RGB values
func ValidateColorArray(ca []int) bool {
	if len(ca) != 3 {
		return false
	}

	result := true

	for _, v := range ca {
		result = result && (v >= 0 && v <= 255)
	}

	return result
}
