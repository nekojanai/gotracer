package camera

import (
	"bytes"
	"fmt"
	"os"

	"cat7.sh/raytracer/color"
	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

type Camera struct {
	aspect_ratio      float64
	image_width       float64
	image_height      float64
	samples_per_pixel int
	center            *vec3.Point3
	pixel00_loc       *vec3.Point3
	pixel_delta_u     *vec3.Vec3
	pixel_delta_v     *vec3.Vec3
}

func NewCamera() *Camera {
	return &Camera{}
}

func (camera *Camera) SetAspectRatio(aspect_ratio float64) {
	camera.aspect_ratio = aspect_ratio
}

func (camera *Camera) SetImageWidth(image_width float64) {
	camera.image_width = image_width
}

func (camera *Camera) Initialize() {
	camera.image_height = camera.image_width / camera.aspect_ratio
	if camera.image_height < 1 {
		camera.image_height = 1
	}

	camera.center = vec3.NewPoint3(0, 0, 0)

	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * camera.image_width / camera.image_height

	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -viewport_height, 0)

	camera.pixel_delta_u = viewport_u.Div(camera.image_width)
	camera.pixel_delta_v = viewport_v.Div(camera.image_height)

	viewport_upper_left := camera.center.Sub(vec3.NewVec3(0, 0, focal_length)).Sub(viewport_u.Div(2.0)).Sub(viewport_v.Div(2.0))
	camera.pixel00_loc = viewport_upper_left.Add(camera.pixel_delta_u.Add(camera.pixel_delta_v).Scale(0.5))

}

func (camera *Camera) Render(world ray.IHittable) {
	camera.Initialize()

	output00 := bytes.Buffer{}
	output00.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", camera.image_width, camera.image_height))

	for j := 0; j < int(camera.image_height); j++ {
		fmt.Printf("lines remaining: %v \n", int(camera.image_height)-j)
		for i := 0; i < int(camera.image_width); i++ {
			pixel_color := color.NewColor(0, 0, 0)
			for sample := 0; sample < camera.samples_per_pixel; sample++ {
				r := camera.getRay(i, j)
				pixel_color = pixel_color.Add(r.Color(world))
			}

			output00.WriteString(color.WriteColor(pixel_color, camera.samples_per_pixel))
		}
		output00.WriteString("\n")
	}

	os.WriteFile("output.ppm", output00.Bytes(), 0644)
}

func (camera *Camera) getRay(i, j int) *ray.Ray {
	// Get a randomly sampled camera ray for the pixel

	pixel_center := camera.pixel00_loc.Add(camera.pixel_delta_u.Scale(float64(i))).Add(camera.pixel_delta_v.Scale(float64(j)))
	pixel_sample := pixel_center.Add(camera.pixelSampleSquare())

	ray_origin := camera.center
	ray_direction := pixel_sample.Sub(ray_origin)

	return ray.NewRay(ray_origin, ray_direction)
}

func (camera *Camera) pixelSampleSquare() *vec3.Vec3 {
	px := -0.5 + util.RandomFloat64(0, 1)
	py := -0.5 + util.RandomFloat64(0, 1)
	return camera.pixel_delta_u.Scale(px).Add(camera.pixel_delta_v.Scale(py))
}
