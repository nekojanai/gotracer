package vec3

import (
	"math"

	"cat7.sh/raytracer/util"
)

type Vec3 struct {
	x, y, z float64
}

type Point3 = Vec3

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x: x, y: y, z: z}
}

func NewPoint3(x, y, z float64) *Point3 {
	return &Point3{x: x, y: y, z: z}
}

func (vec *Vec3) X() float64 {
	return vec.x
}

func (vec *Vec3) Y() float64 {
	return vec.y
}

func (vec *Vec3) Z() float64 {
	return vec.z
}

func (vec *Vec3) Neg() *Vec3 {
	return &Vec3{x: -vec.x, y: -vec.y, z: -vec.z}
}

func (vec *Vec3) Sub(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.x - vec0.x, y: vec.y - vec0.y, z: vec.z - vec0.z}
}

func (vec *Vec3) Add(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.x + vec0.x, y: vec.y + vec0.y, z: vec.z + vec0.z}
}

func (vec *Vec3) Scale(t float64) *Vec3 {
	return &Vec3{x: vec.x * t, y: vec.y * t, z: vec.z * t}
}

func (vec *Vec3) Length() float64 {
	return math.Sqrt(vec.LengthSquared())
}

func (vec *Vec3) LengthSquared() float64 {
	return vec.Dot(vec)
}

func (vec *Vec3) Dot(vec0 *Vec3) float64 {
	return vec.x*vec0.x + vec.y*vec0.y + vec.z*vec0.z
}

func (vec *Vec3) Cross(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.y*vec0.z - vec.z*vec0.y,
		y: vec.z*vec0.x - vec.x*vec.z,
		z: vec.x*vec.y - vec.y*vec.x}
}

func (vec *Vec3) Div(t float64) *Vec3 {
	return vec.Scale((1.0 / t))
}

func (vec *Vec3) Unit() *Vec3 {
	return vec.Div(vec.Length())
}

func (vec *Vec3) RandomOnHemisphere() *Vec3 {
	on_unit_sphere := RandomUnit()
	if on_unit_sphere.Dot(vec) > 0 {
		return on_unit_sphere
	} else {
		return on_unit_sphere.Neg()
	}
}

func RandomInUnitSphere() *Vec3 {
	for {
		p := Random(-1, 1)
		if p.LengthSquared() < 1 {
			return p
		}
	}
}

func RandomUnit() *Vec3 {
	return RandomInUnitSphere().Unit()
}

func Random(min, max float64) *Vec3 {
	return &Vec3{x: util.RandomFloat64(min, max), y: util.RandomFloat64(min, max), z: util.RandomFloat64(min, max)}
}
