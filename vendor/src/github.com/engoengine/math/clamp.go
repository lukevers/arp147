package math

// Clamp returns f clamped to [low, high]
func Clamp(f, low, high float32) float32 {
	if f < low {
		return low
	}
	if f > high {
		return high
	}
	return f
}
