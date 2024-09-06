package render

import (
	"math"
	"math/rand"
)

const MAX_REFLECTION_DEPTH = 10

type Camera struct {
	hsize        int
	vsize        int
	fov          float64
	transform    *Matrix
	pixelSize    float64
	halfWidth    float64
	halfHeight   float64
	parallelism  int
	antialiasing int
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	pixelSize, halfWidth, halfHeight := PixelSize(hsize, vsize, fov)
	return &Camera{
		hsize:        hsize,
		vsize:        vsize,
		fov:          fov,
		transform:    IdentityMatrix(),
		pixelSize:    pixelSize,
		halfWidth:    halfWidth,
		halfHeight:   halfHeight,
		parallelism:  1,
		antialiasing: 1,
	}
}

func (c *Camera) SetTransform(transform *Matrix) {
	c.transform = transform
}

func (c *Camera) SetParallelism(p int) {
	c.parallelism = p
}

func (c *Camera) SetAntialiasing(a int) {
	c.antialiasing = a
}

func PixelSize(hsize, vsize int, fov float64) (float64, float64, float64) {
	halfView := math.Tan(fov / 2)
	aspect := float64(hsize) / float64(vsize)

	var halfWidth float64
	var halfHeight float64
	if aspect >= 1 {
		halfWidth = halfView
		halfHeight = halfView / aspect
	} else {
		halfWidth = halfView * aspect
		halfHeight = halfView
	}

	pixelSize := 2 * halfWidth / float64(hsize)
	return pixelSize, halfWidth, halfHeight
}

func (c *Camera) RayForPixel(x, y int) *Ray {
	return c.actualRayForPixel(float64(x)+0.5, float64(y)+0.5)
}

func (c *Camera) actualRayForPixel(xPos, yPos float64) *Ray {
	xOffset := xPos * c.pixelSize
	yOffset := yPos * c.pixelSize

	worldX := c.halfWidth - xOffset
	worldY := c.halfHeight - yOffset

	// TODO: Probably we should store this
	tInv, err := c.transform.Inverse()
	if err != nil {
		panic(err)
	}

	pixel := tInv.MultiplyTuple(NewPoint(worldX, worldY, -1))
	origin := tInv.MultiplyTuple(NewPoint(0, 0, 0))
	direction := pixel.Sub(origin).Norm()

	return NewRay(origin, direction)
}

func (c *Camera) Render(w *World, progress chan bool) *Canvas {
	image := NewCanvas(c.hsize, c.vsize)

	workIn := make(chan workUnit)
	workOut := make(chan bool)
	go c.progressReporter(progress, workOut)
	for i := 0; i < c.parallelism; i++ {
		go c.worker(workIn, workOut)
	}

	for y := 0; y < c.vsize; y++ {
		for x := 0; x < c.hsize; x++ {
			workIn <- workUnit{w, x, y, image}
		}
	}

	return image
}

func (c *Camera) progressReporter(progress chan bool, workOut chan bool) {
	total := c.hsize * c.vsize
	prevReport := 0
	count := 0

	for <-workOut {
		count++
		if 10*count/total > prevReport {
			progress <- true
			prevReport = 10 * count / total
		}

		if count == total {
			close(workOut)
		}
	}
	close(progress)
}

type workUnit struct {
	w     *World
	x     int
	y     int
	image *Canvas
}

func (c *Camera) renderPixel(w *World, x, y int, image *Canvas, workOut chan bool) {
	accC := NewColor(0, 0, 0)
	if c.antialiasing == 1 {
		accC = w.Shade(c.RayForPixel(x, y), MAX_REFLECTION_DEPTH)
	} else {
		for i := 0; i < c.antialiasing; i++ {
			r := c.actualRayForPixel(float64(x)+rand.Float64(), float64(y)+rand.Float64())
			accC = accC.Add(w.Shade(r, MAX_REFLECTION_DEPTH))
		}
	}

	image.SetAt(x, y, accC.Div(float64(c.antialiasing)))
	workOut <- true
}

func (c *Camera) worker(workIn chan workUnit, workOut chan bool) {
	for task := range workIn {
		c.renderPixel(task.w, task.x, task.y, task.image, workOut)
	}
}
