package math

import "math"

// Boilerplate code
func Clamp(min, max, value float64) float64 {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

// Boilerplate code again
func ClampInt(min, max, value int) int {
	if value < min {
		return min
	} else if value > max {
		return max
	} else {
		return value
	}
}

func Round(value float64) int {
	return int(math.Round(value))
}

// 0 < threshold < 1
// Set the decimal where the value is rounded to its superior value
func RoundWithThreshold(value, threshold float64) int {
	return int(math.Round(value + 0.5 - threshold))
}
