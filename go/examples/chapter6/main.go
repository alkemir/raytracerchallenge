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
	purple := render.NewColor(1, 0.2, 1)
	material := render.NewMaterial(purple, 0.1, 0.9, 0.9, 200)

	sphere := render.NewSphere()
	sphere.SetMaterial(material)

	// Create a light source
	lPos := render.NewPoint(-10, 10, -10)
	lColor := render.NewColor(1, 1, 1)
	light := render.NewPointLight(lPos, lColor)

	// Set the camera a bit away
	cameraOrigin := render.NewPoint(0, 0, -2)
	cameraUpperLeft := render.NewVector(3, 3, 1)

	angleXDelta := render.NewVector(-fovX/float64(resX), 0, 0)
	angleYDelta := render.NewVector(0, -fovY/float64(resY), 0)

	// Cast all rays
	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			cameraDirection := cameraUpperLeft.Add(angleXDelta.Mul(float64(x))).Add(angleYDelta.Mul(float64(y))).Norm()

			r := render.NewRay(cameraOrigin, cameraDirection)

			if hit := render.Hit(sphere.Intersect(r)); hit != nil {
				p := r.Project(hit.T())
				n := hit.Object().(*render.Sphere).Normal(p)
				eye := cameraDirection.Mul(-1)
				image.SetAt(x, y, material.Lightning(light, p, eye, n, false))
			}
		}
	}

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
