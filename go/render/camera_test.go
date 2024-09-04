package render

import (
	"math"
	"testing"
)

func TestCameraConstructor(t *testing.T) {
	hsize := 160
	vsize := 120
	fov := math.Pi / 2

	c := NewCamera(hsize, vsize, fov)

	if c.hsize != 160 {
		t.Fatal("Horizontal size is wrong")
	}
	if c.vsize != 120 {
		t.Fatal("Vertical size is wrong")
	}
	if c.fov != math.Pi/2 {
		t.Fatal("Vertical size is wrong")
	}
	if !c.transform.Equals(IdentityMatrix()) {
		t.Fatal("Transform is wrong")
	}
}

func TestCameraPixelSizeHorizontal(t *testing.T) {
	c := NewCamera(200, 125, math.Pi/2)

	if c.pixelSize != 0.01 {
		t.Fatal("Pixel size is wrong")
	}
}

func TestCameraPixelSizeVertical(t *testing.T) {
	c := NewCamera(125, 200, math.Pi/2)

	if c.pixelSize != 0.01 {
		t.Fatal("Pixel size is wrong")
	}
}

func TestCameraRayCenter(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)

	r := c.RayForPixel(100, 50)

	if !r.origin.Equals(NewPoint(0, 0, 0)) {
		t.Fatal("Origin is wrong")
	}
	if !r.direction.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Direction is wrong")
	}
}

func TestCameraRayCorner(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)

	r := c.RayForPixel(0, 0)

	if !r.origin.Equals(NewPoint(0, 0, 0)) {
		t.Fatal("Origin is wrong")
	}
	if !r.direction.Equals(NewVector(0.66519, 0.33259, -0.66851)) {
		t.Fatal("Direction is wrong")
	}
}
func TestCameraRayTransformed(t *testing.T) {
	c := NewCamera(201, 101, math.Pi/2)
	c.SetTransform(RotationY(math.Pi / 4).Multiply(Translation(0, -2, 5)))

	r := c.RayForPixel(100, 50)

	if !r.origin.Equals(NewPoint(0, 2, -5)) {
		t.Fatal("Origin is wrong")
	}
	if !r.direction.Equals(NewVector(math.Sqrt2/2, 0, -math.Sqrt2/2)) {
		t.Fatal("Direction is wrong")
	}
}

func TestCameraRender(t *testing.T) {
	w := DefaultWorld()
	c := NewCamera(11, 11, math.Pi/2)
	from := NewPoint(0, 0, -5)
	to := NewPoint(0, 0, 0)
	up := NewVector(0, 1, 0)
	c.SetTransform(View(from, to, up))

	image := c.Render(w, make(chan bool, 11*11))

	if !image.GetAt(5, 5).Equals(NewColor(0.38066, 0.47583, 0.2855)) {
		t.Fatal("Color at center of image is wrong")
	}
}
