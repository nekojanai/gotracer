package color

import (
	"fmt"
	"math"

	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

type Color = vec3.Vec3

func NewColor(r, g, b float64) *Color {
	return vec3.NewVec3(r, g, b)
}

func LinearToGamma(linear_component float64) float64 {
	if linear_component > 0 {
		return math.Sqrt(linear_component)
	}

	return linear_component
}

func WriteColor(color *Color, samples_per_pixel int) string {
	intensity := util.NewInterval(0.000, 0.999)
	r := math.Floor(math.Abs(255 * intensity.Clamp(LinearToGamma(color.X())*(1.0/float64(samples_per_pixel)))))
	g := math.Floor(math.Abs(255 * intensity.Clamp(LinearToGamma(color.Y())*(1.0/float64(samples_per_pixel)))))
	b := math.Floor(math.Abs(255 * intensity.Clamp(LinearToGamma(color.Z())*(1.0/float64(samples_per_pixel)))))

	return fmt.Sprintf("%v %v %v ", r, g, b)
}
