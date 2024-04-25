package math

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
