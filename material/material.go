package material

import (
	"image/color"

	"cat7.sh/raytracer/ray"
)

type Material struct {
}

func (mat *Material) Scatter(r_in *ray.Ray, rec *ray.HitRecord, attenuation *color.Color, scattered *ray.Ray) bool {
	return false
}
