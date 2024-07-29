package main

import (
	"bytes"
	"fmt"
	"os"

	"cat7.sh/raytracer/color"
	"cat7.sh/raytracer/ray"
	"cat7.sh/raytracer/vec3"
)

func main() {
	aspect_ratio := 16.0 / 9.0
	width := 256.0

	// Calculate the image height and ensure that it's at least 1
	height := (width / aspect_ratio)
	if height < 1 {
		height = 1
	}

	// Camera
	focal_length := 1.0
	viewport_height := 2.0
	viewport_width := viewport_height * width / height
	camera_center := vec3.NewVec3(0, 0, 0)

	// calculate the vectors acrpss the horizontal and down the vertical viewport edges
	viewport_u := vec3.NewVec3(viewport_width, 0, 0)
	viewport_v := vec3.NewVec3(0, -viewport_height, 0)

	// calculate the horizontal and vertical delta vectors from pixel to pixel
	pixel_delta_u := viewport_u.Vec3DivWFloat(height)
	pixel_delta_v := viewport_v.Vec3DivWFloat(width)

	// Calculate the location of the upper left pixel
	viewport_upper_left := camera_center.Vec3Sub(vec3.NewVec3(0, 0, focal_length)).Vec3Sub(viewport_u.Vec3DivWFloat(2)).Vec3Sub(viewport_v.Vec3DivWFloat(2))
	pixel00_loc := viewport_upper_left.Vec3Add(vec3.NewVec3(0.5, 0.5, 0.5).Vec3Mult((pixel_delta_u.Vec3Add(pixel_delta_v))))

	// Render
	output00 := bytes.Buffer{}
	output00.WriteString(fmt.Sprintf("P3\n%v %v\n255\n", width, height))

	for j := 0; j < int(height); j++ {
		fmt.Printf("lines remaining: %v \n", int(height)-j)
		for i := 0; i < int(width); i++ {
			pixel_center := pixel00_loc.Vec3Add(pixel_delta_u.Vec3MultWFloat(float64(i)).Vec3Add(pixel_delta_v.Vec3MultWFloat(float64(j))))
			ray_direction := pixel_center.Vec3Sub(camera_center)
			r := ray.NewRay(camera_center, ray_direction)

			pixel_color := color.RayColor(&r)
			output00.WriteString(color.WriteColor(pixel_color))
		}
		output00.WriteString("\n")
	}

	os.WriteFile("output.ppm", output00.Bytes(), 0644)
}
