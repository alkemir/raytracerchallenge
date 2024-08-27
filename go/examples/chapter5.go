package main

import (
	"os"
	"raytracerchallenge/canvas"
	"raytracerchallenge/ray"
	"raytracerchallenge/shape"
	"raytracerchallenge/tuple"
)

func main() {
	// Field of view size
	fovX := 6.0
	fovY := 6.0

	// Image resolution
	resX := 1000
	resY := 1000

	image := canvas.NewCanvas(resX, resY)

	// Create a sphere
	sphere := shape.NewSphere()

	// Color
	red := tuple.NewColor(0.5, 0, 0)

	// Set the camera a bit away
	cameraOrigin := tuple.NewPoint(0, 0, -2)
	cameraUpperLeft := tuple.NewVector(3, 3, 1)

	angleXDelta := tuple.NewVector(-fovX/float64(resX), 0, 0)
	angleYDelta := tuple.NewVector(0, -fovY/float64(resY), 0)

	// Cast all rays
	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			cameraDirection := cameraUpperLeft.Add(angleXDelta.Mul(float64(x))).Add(angleYDelta.Mul(float64(y)))

			r := ray.NewRay(cameraOrigin, cameraDirection)

			if shape.Hit(sphere.Intersect(r)) != nil {
				image.SetAt(x, y, red)
			}
		}
	}

	f, err := os.CreateTemp("./", "chapter5*.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
