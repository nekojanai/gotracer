package ray

import (
	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

type HitRecord struct {
	p          *vec3.Point3
	normal     *vec3.Vec3
	t          float64
	front_face bool
}

func (hitRecord *HitRecord) SetFaceNormal(ray *Ray, outward_normal *vec3.Vec3) {
	// Seets the hit record normal vector
	// NOTE: the parameter `outward_normal` is assumed to have unit length

	hitRecord.front_face = ray.dir.Dot(outward_normal) < 0
	if hitRecord.front_face {
		hitRecord.normal = outward_normal
	} else {
		hitRecord.normal = outward_normal.Neg()
	}
}

func (hitRecord *HitRecord) P() *vec3.Point3 {
	return hitRecord.p
}

func (hitRecord *HitRecord) SetP(p *vec3.Point3) {
	hitRecord.p = p
}

func (hitRecord *HitRecord) Normal() *vec3.Vec3 {
	return hitRecord.normal
}

func (hitRecord *HitRecord) SetNormal(normal *vec3.Vec3) {
	hitRecord.normal = normal
}

func (hitRecord *HitRecord) T() float64 {
	return hitRecord.t
}

func (hitRecord *HitRecord) SetT(t float64) {
	hitRecord.t = t
}

func (hitRecord *HitRecord) FrontFace() bool {
	return hitRecord.front_face
}

func (hitRecord *HitRecord) SetFrontFace(front_face bool) {
	hitRecord.front_face = front_face
}

func (hitRecord *HitRecord) Override(hitRecord0 *HitRecord) {
	hitRecord.SetP(hitRecord0.p)
	hitRecord.SetNormal(hitRecord0.normal)
	hitRecord.SetT(hitRecord0.t)
	hitRecord.SetFrontFace(hitRecord.front_face)
}

type IHittable interface {
	Hit(ray *Ray, ray_t *util.Interval, rec *HitRecord) bool
}

type HittableList struct {
	objects []IHittable
}

func (hittableList *HittableList) Add(object IHittable) {
	hittableList.objects = append(hittableList.objects, object)
}

func (hittableList *HittableList) Clear() {
	hittableList.objects = nil
}

func (hittableList HittableList) Hit(ray *Ray, ray_t *util.Interval, rec *HitRecord) bool {
	temp_rec := &HitRecord{}
	hit_anything := false
	closest_so_far := ray_t.Max()

	for _, object := range hittableList.objects {
		if object.Hit(ray, util.NewInterval(ray_t.Min(), closest_so_far), temp_rec) {
			hit_anything = true
			closest_so_far = temp_rec.T()
			rec.Override(temp_rec)
		}
	}

	return hit_anything
}
