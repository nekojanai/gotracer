package ray

import "cat7.sh/raytracer/vec3"

// Ray function P(t) = A + tb
// P := 3D position along a line
// A := ray origin
// b := ray direction
// t := real number
// put in a different t and P(t) moves the point along the ray.
// negative t values move it anywhere on the 3D line

type Ray struct {
	orig *vec3.Vec3
	dir  *vec3.Vec3
}

func NewRay(orig *vec3.Vec3, dir *vec3.Vec3) Ray {
	return Ray{orig: orig, dir: dir}
}

func (ray *Ray) Origin() *vec3.Vec3 {
	return ray.orig
}

func (ray *Ray) Direction() *vec3.Vec3 {
	return ray.dir
}

func (ray *Ray) At(t float64) *vec3.Vec3 {
	return ray.orig.Vec3Add(ray.dir.Vec3Mult(vec3.NewVec3(t, t, t)))
}
