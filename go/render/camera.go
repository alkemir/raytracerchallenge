package render

import (
	"math"
)

type Camera struct {
	hsize      int
	vsize      int
	fov        float64
	transform  *Matrix
	pixelSize  float64
	halfWidth  float64
	halfHeight float64
}

func NewCamera(hsize, vsize int, fov float64) *Camera {
	pixelSize, halfWidth, halfHeight := PixelSize(hsize, vsize, fov)
	return &Camera{
		hsize:      hsize,
		vsize:      vsize,
		fov:        fov,
		transform:  IdentityMatrix(),
		pixelSize:  pixelSize,
		halfWidth:  halfWidth,
		halfHeight: halfHeight,
	}
}

func (c *Camera) SetTransform(transform *Matrix) {
	c.transform = transform
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
	xOffset := (float64(x) + 0.5) * c.pixelSize
	yOffset := (float64(y) + 0.5) * c.pixelSize

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

func (c *Camera) Render(w *World) *Canvas {
	image := NewCanvas(c.hsize, c.vsize)

	for y := 0; y < c.vsize; y++ {
		for x := 0; x < c.hsize; x++ {
			r := c.RayForPixel(x, y)
			c := w.Shade(r)
			image.SetAt(x, y, c)
		}
	}

	return image
}
