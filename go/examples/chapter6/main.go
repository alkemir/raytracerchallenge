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
	purple := tuple.NewColor(1, 0.2, 1)
	material := shape.NewMaterial(
		purple,
		shape.DefaultMaterial.Ambient(),
		shape.DefaultMaterial.Diffuse(),
		shape.DefaultMaterial.Specular(),
		shape.DefaultMaterial.Shininess())

	sphere := shape.NewSphere()
	sphere.SetMaterial(material)

	// Create a light source
	lPos := tuple.NewPoint(-10, 10, -10)
	lColor := tuple.NewColor(1, 1, 1)
	light := shape.NewPointLight(lPos, lColor)

	// Set the camera a bit away
	cameraOrigin := tuple.NewPoint(0, 0, -2)
	cameraUpperLeft := tuple.NewVector(3, 3, 1)

	angleXDelta := tuple.NewVector(-fovX/float64(resX), 0, 0)
	angleYDelta := tuple.NewVector(0, -fovY/float64(resY), 0)

	// Cast all rays
	for x := 0; x < resX; x++ {
		for y := 0; y < resY; y++ {
			cameraDirection := cameraUpperLeft.Add(angleXDelta.Mul(float64(x))).Add(angleYDelta.Mul(float64(y))).Norm()

			r := ray.NewRay(cameraOrigin, cameraDirection)

			if hit := shape.Hit(sphere.Intersect(r)); hit != nil {
				p := r.Project(hit.T())
				n := hit.Object().(*shape.Sphere).Normal(p)
				eye := r.Direction().Mul(-1)
				image.SetAt(x, y, hit.Object().(*shape.Sphere).Material().Lightning(light, p, eye, n))
			}
		}
	}

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
