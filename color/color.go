package color

import (
	"fmt"
	"math"

	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/vec3"
)

type Color = *vec3.Vec3

func NewColor(r, g, b float64) Color {
	return Color(vec3.NewVec3(r, g, b))
}

func WriteColor(color Color) string {
	return fmt.Sprintf("%v %v %v ", math.Floor(math.Abs(color.X()*255)), math.Floor(math.Abs(color.Y()*255)), math.Floor(math.Abs(color.Z()*255)))
}

func RayColor(ray *ray.Ray) Color {
	unit_direction := ray.Direction().Vec3Unit()
	a := (unit_direction.Y() + 1.0) * 0.5
	return NewColor(1.0, 1.0, 1.0).Vec3MultWFloat(1.0 - a).Vec3Add(NewColor(0.5, 0.7, 1.0).Vec3MultWFloat(a))
}
