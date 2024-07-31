package geometry

import (
	"math"

	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

type Sphere struct {
	center *vec3.Point3
	radius float64
}

func NewSphere(center *vec3.Vec3, radius float64) *Sphere {
	return &Sphere{center: center, radius: radius}
}

func (sphere *Sphere) Hit(ray *ray.Ray, ray_t *util.Interval, rec *ray.HitRecord) bool {
	oc := ray.Origin().Sub(sphere.center)
	a := ray.Direction().LengthSquared()
	half_b := oc.Dot(ray.Direction())
	c := oc.LengthSquared() - sphere.radius*sphere.radius

	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	// Find the nearest root that lies in the acceptable range
	root := (-half_b - sqrtd) / a
	if ray_t.Surrounds(root) {
		root = (-half_b + sqrtd) / a
		if ray_t.Surrounds(root) {
			return false
		}
	}

	rec.SetT(root)
	rec.SetP(ray.At(rec.T()))
	rec.SetNormal((rec.P().Sub(sphere.center).Div(sphere.radius)))
	outward_normal := rec.P().Sub(sphere.center).Div(sphere.radius)
	rec.SetFaceNormal(ray, outward_normal)

	return true
}
