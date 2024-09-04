package main

import (
	"math/rand"
)

// Provides int from range [min, max]
func GetRandIntInRange(min, max int) int {
	return rand.Intn(max-min) + min
}

// Provides random radians value from range [0,360] degrees
func GetRandRadians() float64 {
	return rand.Float64() * PI2
}

// Provides float64 from range [-1,1]
func GetRandSignedFloat64() float64 {
	val := rand.Float64()
	if bool := rand.Intn(2); bool == 1 {
		return -val
	}
	return val
}
