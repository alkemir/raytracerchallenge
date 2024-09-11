package main

import (
	"fmt"
	"math"
	"os"
	"raytracerchallenge/render"
	"time"
)

func hexagonCorner() render.Shape {
	corner := render.NewSphere()
	corner.SetTransform(render.Translation(0, 0, -1).Multiply(render.Scaling(0.25, 0.25, 0.25)))
	corner.SetMaterial(render.NewMaterial(render.NewColor(1, 1.8, 0.1), 0, 0, 0, 1.0, 0, 1, 200, nil))
	return corner
}

func hexagonEdge() render.Shape {
	edge := render.NewCylinder()
	edge.SetMinimum(0)
	edge.SetMaximum(1)
	edge.SetTransform(render.Translation(0, 0, -1).Multiply(render.RotationY(-math.Pi / 6)).Multiply(render.RotationZ(-math.Pi / 2)).Multiply(render.Scaling(0.25, 1, 0.25)))
	edge.SetMaterial(render.NewMaterial(render.NewColor(1, 1.8, 0.1), 0, 0, 0, 0, 1.0, 1.5, 200, nil))
	return edge
}

func hexagonSide() render.Shape {
	side := render.NewGroup()
	side.Add(hexagonCorner())
	side.Add(hexagonEdge())

	return side
}

func hexagon() render.Shape {
	hex := render.NewGroup()

	for i := 0; i < 6; i++ {
		side := hexagonSide()
		side.SetTransform(render.RotationY(float64(i) * math.Pi / 3))
		hex.Add(side)
	}

	return hex
}

func main() {
	from := render.NewPoint(0, 2.5, -5)
	to := render.NewPoint(0, 0.5, 0)
	up := render.NewVector(0, 1, 0)
	camera := render.NewCamera(1000, 500, math.Pi/3)
	camera.SetTransform(render.View(from, to, up))
	camera.SetParallelism(32)
	camera.SetAntialiasing(8)

	world := render.NewWorld()
	l := render.NewPointLight(render.NewPoint(-10, 10, -10), render.NewColor(3, 3, 3))
	world.AddLight(l)

	patternCheckers := render.NewCheckersPattern(render.NewColor(1, 1, 1), render.NewColor(0, 0, 0))

	floor := render.NewPlane()
	floor.SetTransform(render.Translation(0, -0.1, 0))
	floorMaterial := render.NewMaterial(render.NewColor(1, 0.9, 0.9), 0.1, 0.9, 0, 0, 0, 1, 200, patternCheckers)
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)

	hex := hexagon()
	hex.SetTransform(render.RotationX(-math.Pi / 6).Multiply(render.Translation(0, 1.5, 0)))
	world.AddObject(hex)

	progress := make(chan bool)
	go func(progress chan bool) {
		count := 0
		for <-progress {
			count++
			fmt.Println(count, "/ 10")
		}
	}(progress)

	startTime := time.Now()
	image := camera.Render(world, progress)
	duration := time.Since(startTime)

	fmt.Println("Rendered in", duration)

	f, err := os.Create("example.png")
	if err != nil {
		panic(err)
	}
	image.Encode(f)
}
