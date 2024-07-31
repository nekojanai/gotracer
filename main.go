package main

import (
	"bytes"
	"fmt"
	"os"

	"cat7.sh/raytracer/color"
	"cat7.sh/raytracer/geometry"
	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/vec3"
)

func main() {
	aspect_ratio := 16.0 / 9.0
	width := 1920.0

	// Calculate the image height and ensure that it's at least 1
	height := (width / aspect_ratio)
	if height < 1 {
		height = 1
	}

	// World

	world := ray.HittableList{}

	world.Add(geometry.NewSphere(vec3.NewPoint3(0, -100.5, -1), 100))
	world.Add(geometry.NewSphere(vec3.NewPoint3(0, 0, -1), 0.5))

	// Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * (width / height)
	camera_center := vec3.NewVec3(0, 0, 0)

	// calculate the vectors acrpss the horizontal and down the vertical viewport edges
	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -viewport_height, 0)

	// calculate the horizontal and vertical delta vectors from pixel to pixel
	pixel_delta_u := viewport_u.Div(width)
	pixel_delta_v := viewport_v.Div(height)

	// Calculate the location of the upper left pixel
	viewport_upper_left := camera_center.Sub(vec3.NewVec3(0, 0, focal_length)).Sub(viewport_u.Div(2)).Sub(viewport_v.Div(2))
	pixel00_loc := viewport_upper_left.Add(pixel_delta_u.Add(pixel_delta_v).Scale(0.5))

	// Render
	output00 := bytes.Buffer{}
	output00.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", width, height))

	for j := 0; j < int(height); j++ {
		fmt.Printf("lines remaining: %v \n", int(height)-j)
		for i := 0; i < int(width); i++ {
			pixel_center := pixel00_loc.Add(pixel_delta_u.Scale(float64(i)).Add(pixel_delta_v.Scale(float64(j))))
			ray_direction := pixel_center.Sub(camera_center)
			r := ray.NewRay(camera_center, ray_direction)

			pixel_color := r.Color(world)
			output00.WriteString(color.WriteColor(pixel_color))
		}
		output00.WriteString("\n")
	}

	os.WriteFile("output.ppm", output00.Bytes(), 0644)
}
