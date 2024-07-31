package util

import (
	"math"
	"math/rand"
)

func DegreesTpRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RandomFloat64(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}
