package util

import (
	"fmt"
	"math"
	"math/rand"
)

func DegreesTpRadians(degrees float64) float64 {
	return degrees * math.Pi / 180.0
}

func RandomFloat64(min, max float64) float64 {
	return min + (max-min)*rand.Float64()
}

func RandomFilename(prefix, suffix string) string {
	return fmt.Sprintf("%v%v%v", prefix, randomString(), suffix)
}

func randomString() string {
	return fmt.Sprintf("%v", rand.Intn(math.MaxInt)+1)
}
