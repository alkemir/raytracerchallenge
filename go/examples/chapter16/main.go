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
	camera.SetAntialiasing(3)

	world := render.NewWorld()
	l1 := render.NewPointLight(render.NewPoint(-10, 10, -10), render.NewColor(1, 1, 1))
	world.AddLight(l1)

	patternCheckers := render.NewCheckersPattern(render.NewColor(1, 1, 1), render.NewColor(0.5, 0.5, 0.5))

	floor := render.NewPlane()
	floor.SetTransform(render.Translation(0, -0.1, 0))
	floorMaterial := render.NewMaterial(render.NewColor(1, 0.9, 0.9), 0.1, 0.9, 0, 0, 0, 1, 200, patternCheckers)
	floor.SetMaterial(floorMaterial)
	world.AddObject(floor)

	greenGlass := render.NewMaterial(render.NewColor(0, 0.3, 0), 0.1, 0.1, 0.8, 0, 0.8, 2, 200, nil)
	c1 := render.NewCube()
	c1.SetMaterial(greenGlass)
	c2 := render.NewCube()
	c2.SetMaterial(greenGlass)
	c2.SetTransform(render.Translation(0.5, 0.5, 0.1))

	g := render.NewGroup()
	g.Add(c1)
	g.Add(c2)
	g.SetTransform(render.Translation(-1.5, 0, 0))
	world.AddObject(g)

	c3 := render.NewCube()
	c3.SetMaterial(greenGlass)
	c4 := render.NewCube()
	c4.SetMaterial(greenGlass)
	c4.SetTransform(render.Translation(0.5, 0.5, 0.1))
	csg := render.NewCSGShape(render.CSG_UNION, c3, c4)
	csg.SetTransform(render.Translation(1.5, 0, 0))
	world.AddObject(csg)

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
