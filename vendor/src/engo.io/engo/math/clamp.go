package math

import (
	engomath "github.com/engoengine/math"
)

// Clamp returns f clamped to [low, high]
func Clamp(f, low, high float32) float32 {
	return engomath.Clamp(f, low, high)
}
