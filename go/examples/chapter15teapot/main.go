package main

import (
	"bytes"
	"fmt"
	"math"
	"os"
	"raytracerchallenge/render"
	"time"
)

func main() {
	from := render.NewPoint(0, 2.5, -5)
	to := render.NewPoint(0, 0.5, 0)
	up := render.NewVector(0, 1, 0)
	camera := render.NewCamera(1000, 500, math.Pi/3)
	camera.SetTransform(render.View(from, to, up))
	camera.SetParallelism(16)
	camera.SetAntialiasing(1)

	world := render.NewWorld()
	l1 := render.NewPointLight(render.NewPoint(-10, 10, -10), render.NewColor(1, 1, 1))
	world.AddLight(l1)

	patternCheckers := render.NewCheckersPattern(render.NewColor(1, 1, 1), render.NewColor(0.5, 0.5, 0.5))

	floor := render.NewPlane()
	floor.SetTransform(render.Translation(0, -0.1, 0))
	floorMaterial := render.NewMaterial(render.NewColor(1, 0.9, 0.9), 0.1, 0.9, 0, 0, 0, 1, 200, patternCheckers)
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)

	teapotData, err := os.ReadFile("teapot.obj")
	if err != nil {
		panic(err)
	}
	teapotParser := render.NewParser(bytes.NewReader(teapotData))
	teapotParser.Parse()
	teapot := teapotParser.AsGroup()
	teapot.SetTransform(render.Scaling(0.3, 0.3, 0.3))

	world.AddObject(teapot)

	progress := make(chan bool)
	go func(progress chan bool) {
		count := 0
		fmt.Println(count, "/ 10")
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
