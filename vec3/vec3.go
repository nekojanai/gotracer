package vec3

import "math"

type Vec3 struct {
	x, y, z float64
}

func NewVec3(x, y, z float64) *Vec3 {
	return &Vec3{x: x, y: y, z: z}
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

func (vec *Vec3) Vec3Sub(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.x - vec0.x, y: vec.y - vec0.y, z: vec.z - vec0.z}
}

func (vec *Vec3) Vec3Add(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.x + vec0.x, y: vec.y + vec0.y, z: vec.z + vec0.z}
}

func (vec *Vec3) Vec3Mult(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.x * vec0.x, y: vec.y * vec0.y, z: vec.z * vec0.z}
}

func (vec *Vec3) Vec3MultWFloat(t float64) *Vec3 {
	return vec.Vec3Mult(NewVec3(t, t, t))
}

func (vec *Vec3) Vec3Length() float64 {
	return math.Sqrt(vec.Vec3LengthSquared())
}

func (vec *Vec3) Vec3LengthSquared() float64 {
	return vec.x*vec.x + vec.y*vec.y + vec.z*vec.z
}

func (vec *Vec3) Vec3Dot(vec0 *Vec3) float64 {
	return vec.x*vec0.x + vec.y*vec0.y + vec.z*vec0.z
}

func (vec *Vec3) Vec3Cross(vec0 *Vec3) *Vec3 {
	return &Vec3{x: vec.y*vec0.z - vec.z*vec0.y,
		y: vec.z*vec0.x - vec.x*vec.z,
		z: vec.x*vec.y - vec.y*vec.x}
}

func (vec *Vec3) Vec3DivWFloat(t float64) *Vec3 {
	return vec.Vec3MultWFloat((1.0 / t))
}

func (vec *Vec3) Vec3Unit() *Vec3 {
	return vec.Vec3DivWFloat(vec.Vec3Length())
}
