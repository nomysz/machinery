package machinery

import (
	"math"
)

const (
	PIDiv180 = math.Pi / 180.0
	PI2      = math.Pi * 2
)

func IsEqualish(a, b float64) bool {
	return math.Abs(b-a) < 10e-5
}

func NormalizeRadians(radians float64) float64 {
	radians = math.Mod(radians, PI2)
	if radians > math.Pi {
		radians -= PI2
	}
	return radians
}

func Clamp(x, min, max float64) float64 {
	if x < min {
		return min
	} else if x > max {
		return max
	}
	return x
}
