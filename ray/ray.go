package ray

import (
	"math"

	"cat7.sh/raytracer/color"
	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

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

func NewRay(orig *vec3.Vec3, dir *vec3.Vec3) *Ray {
	return &Ray{orig: orig, dir: dir}
}

func (ray *Ray) Origin() *vec3.Vec3 {
	return ray.orig
}

func (ray *Ray) Direction() *vec3.Vec3 {
	return ray.dir
}

func (ray *Ray) At(t float64) *vec3.Vec3 {
	return ray.orig.Add(ray.dir.Scale(t))
}

func (ray *Ray) HitSphere(center *vec3.Vec3, radius float64) float64 {
	oc := ray.orig.Sub(center)
	a := ray.dir.LengthSquared()
	half_b := oc.Dot(ray.dir)
	c := oc.LengthSquared() - radius*radius
	discriminant := half_b*half_b - a*c

	if discriminant < 0 {
		return -1.0
	} else {
		return (-half_b - math.Sqrt(discriminant)) / a
	}
}

func (ray *Ray) Color(world IHittable) *color.Color {
	rec := &HitRecord{}
	if world.Hit(ray, util.NewInterval(0, math.Inf(0)), rec) {
		return rec.normal.Add(color.NewColor(0, 0, 0)).Scale(0.5)
	}

	unit_direction := ray.Direction().Unit()
	a := (unit_direction.Y() + 1.0) * 0.5
	return color.NewColor(1.0, 1.0, 1.0).Scale(1.0 - a).Add(color.NewColor(0.5, 0.7, 1.0).Scale(a))
}
