package main

import (
	"math"
	"os"
	"raytracerchallenge/render"
)

func main() {
	from := render.NewPoint(0, 1.5, -5)
	to := render.NewPoint(0, 1, 0)
	up := render.NewVector(0, 1, 0)
	camera := render.NewCamera(1000, 500, math.Pi/3)
	camera.SetTransform(render.View(from, to, up))

	patternStripes := render.NewStripesPattern(render.NewColor(0.1, 0.1, 1), render.NewColor(1, 0.1, 0.1))
	patternStripes.SetTransform(render.Translation(0.5, 0, 0))
	patternCheckers := render.NewCheckersPattern(render.NewColor(1, 1, 1), render.NewColor(0, 0, 0))
	patternRings := render.NewRingPattern(render.NewColor(0.0, 0.3, 0), render.NewColor(0, 1, 0))
	patternRings.SetTransform(render.Scaling(0.1, 0.1, 0.1).Multiply(render.Translation(0, 0, -0.5)))
	patternGradient := render.NewGradientPattern(render.NewColor(1, 0, 0), render.NewColor(1, 1, 0))
	patternGradient.SetTransform(render.RotationY(math.Pi/2 + 0.15))

	world := render.NewWorld()
	l := render.NewPointLight(render.NewPoint(-10, 10, -10), render.NewColor(1, 1, 1))
	world.AddLight(l)

	floor := render.NewPlane()
	floor.SetTransform(render.Translation(0, -0.1, 0))
	floorMaterial := render.NewMaterial(render.NewColor(1, 0.9, 0.9), 0.1, 0.9, 0, 200, patternCheckers)
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)

	middle := render.NewSphere()
	middle.SetTransform(render.Translation(-0.5, 1, 0.5))
	middleMaterial := render.NewMaterial(render.NewColor(0.1, 1, 0.5), 0.1, 0.7, 0.3, 200, patternRings)
	middle.SetMaterial(middleMaterial)
	world.AddObject(middle)

	right := render.NewSphere()
	right.SetTransform(render.Translation(1.5, 0.5, -0.5).Multiply(render.Scaling(0.5, 0.5, 0.5)))
	rightMaterial := render.NewMaterial(render.NewColor(0.5, 1, 0.1), 0.1, 0.7, 0.3, 200, patternGradient)
	right.SetMaterial(rightMaterial)
	world.AddObject(right)

	left := render.NewSphere()
	left.SetTransform(render.Translation(-1.5, 0.33, -0.75).Multiply(render.Scaling(0.33, 0.33, 0.33)))
	leftMaterial := render.NewMaterial(render.NewColor(1, 0.8, 0.1), 0.1, 0.7, 0.3, 200, patternStripes)
	left.SetMaterial(leftMaterial)
	world.AddObject(left)

	image := camera.Render(world)

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
