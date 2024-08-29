package render

import (
	"image"
	"image/png"
	"io"
)

type Canvas struct {
	width  int
	height int
	data   []Tuple
}

func NewCanvas(width, height int) *Canvas {
	return &Canvas{
		width:  width,
		height: height,
		data:   make([]Tuple, width*height),
	}
}

func (c *Canvas) GetAt(x, y int) Tuple {
	return c.data[y*c.width+x]
}

func (c *Canvas) SetAt(x, y int, t Tuple) {
	c.data[y*c.width+x] = t
}

func (c *Canvas) Encode(w io.Writer) {
	img := image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{c.width, c.height}})

	for x := 0; x < c.width; x++ {
		for y := 0; y < c.height; y++ {
			px := c.GetAt(x, y)
			img.Set(x, y, px)
		}
	}

	png.Encode(w, img)
}
