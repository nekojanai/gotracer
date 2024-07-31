package color

import (
	"fmt"
	"math"

	"cat7.sh/raytracer/vec3"
)

type Color = vec3.Vec3

func NewColor(r, g, b float64) *Color {
	return vec3.NewVec3(r, g, b)
}

func WriteColor(color *Color) string {
	return fmt.Sprintf("%v %v %v ", math.Floor(math.Abs(color.X()*255)), math.Floor(math.Abs(color.Y()*255)), math.Floor(math.Abs(color.Z()*255)))
}
