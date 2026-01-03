package machinery

import (
	"math"
)

const (
	PIDiv180 = math.Pi / 180.0
	PI2      = math.Pi * 2
)

func IsEqualish(a, b float64) bool {
	return math.Abs(b-a) < .0001
}

func NormalizeRadians(radians float64) float64 {
	radians = math.Mod(radians, PI2)
	if radians > math.Pi {
		radians -= PI2
	}
	return radians
}
