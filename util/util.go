package util

import "math"

func DegreesTpRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}
