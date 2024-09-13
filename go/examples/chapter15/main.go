package main

import (
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
	camera.SetParallelism(32)
	camera.SetAntialiasing(8)

	world := render.NewWorld()
	l1 := render.NewPointLight(render.NewPoint(-10, 10, -10), render.NewColor(1, 1, 1))
	l2 := render.NewPointLight(render.NewPoint(-10, 1, -10), render.NewColor(1, 1, 1))
	world.AddLight(l1)
	world.AddLight(l2)

	patternCheckers := render.NewCheckersPattern(render.NewColor(1, 1, 1), render.NewColor(0, 0, 0))

	floor := render.NewPlane()
	floor.SetTransform(render.Translation(0, -0.1, 0))
	floorMaterial := render.NewMaterial(render.NewColor(1, 0.9, 0.9), 0.1, 0.9, 0, 0, 0, 1, 200, patternCheckers)
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)

	crystal := render.NewGroup()
	top := render.NewPoint(0, 1, 0)
	a := render.NewPoint(1, 0, 0)
	b := render.NewPoint(0, 0, 1)
	c := render.NewPoint(-1, 0, 0)
	d := render.NewPoint(0, 0, -1)
	bottom := render.NewPoint(0, -1, 0)

	tab := render.NewTriangle(top, a, b)
	tbc := render.NewTriangle(top, b, c)
	tcd := render.NewTriangle(top, c, d)
	tda := render.NewTriangle(top, d, a)
	bab := render.NewTriangle(bottom, a, b)
	bbc := render.NewTriangle(bottom, b, c)
	bcd := render.NewTriangle(bottom, c, d)
	bda := render.NewTriangle(bottom, d, a)

	tt := []*render.Triangle{tab, tbc, tcd, tda, bab, bbc, bcd, bda}
	for _, t := range tt {
		t.SetMaterial(render.NewMaterial(render.NewColor(1, 1, 1), 0, 0, 0, 0, 1.0, 1.5, 200, nil))
		crystal.Add(t)
	}

	crystal.SetTransform(render.Translation(0, 1.5, -2.3).Multiply(render.RotationY(math.Pi / 6).Multiply(render.Scaling(0.5, 0.5, 0.5))))
	world.AddObject(crystal)

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
