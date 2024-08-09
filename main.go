package main

import (
	"os"

	"cat7.sh/raytracer/camera"
	"cat7.sh/raytracer/geometry"
	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/util"
	"cat7.sh/raytracer/vec3"
)

func main() {
	// World

	world := ray.HittableList{}

	world.Add(geometry.NewSphere(vec3.NewPoint3(0, -100.5, -1), 100))
	world.Add(geometry.NewSphere(vec3.NewPoint3(0, 0, -1), 0.5))

	camera := camera.NewCamera()
	camera.SetAspectRatio(16.0 / 9.0)
	camera.SetImageWidth(640.0)

	switch len(os.Args) {
	case 3:
		camera.SetSamplesPerPixel(util.ParseIntArg(2))
		camera.SetImageWidth(float64(util.ParseIntArg(1)))
	case 2:
		camera.SetImageWidth(float64(util.ParseIntArg(1)))
	default:
	}

	camera.Render(world)
}
