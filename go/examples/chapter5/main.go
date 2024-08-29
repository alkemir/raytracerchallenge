package main

import (
	"os"
	"raytracerchallenge/render"
)

func main() {
	// Field of view size
	fovX := 6.0
	fovY := 6.0

	// Image resolution
	resX := 1000
	resY := 1000

	image := render.NewCanvas(resX, resY)

	// Create a sphere
	sphere := render.NewSphere()

	// Color
	red := render.NewColor(0.5, 0, 0)

	// Set the camera a bit away
	cameraOrigin := render.NewPoint(0, 0, -2)
	cameraUpperLeft := render.NewVector(3, 3, 1)

	angleXDelta := render.NewVector(-fovX/float64(resX), 0, 0)
	angleYDelta := render.NewVector(0, -fovY/float64(resY), 0)

	// Cast all rays
	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			cameraDirection := cameraUpperLeft.Add(angleXDelta.Mul(float64(x))).Add(angleYDelta.Mul(float64(y)))

			r := render.NewRay(cameraOrigin, cameraDirection)

			if render.Hit(sphere.Intersect(r)) != nil {
				image.SetAt(x, y, red)
			}
		}
	}

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
