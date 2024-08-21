package canvas

import (
	"os"
	"raytracerchallenge/tuple"
	"testing"
)

func TestCanvasIsBlack(t *testing.T) {
	c := NewCanvas(10, 20)

	for x := 0; x < 10; x++ {
		for y := 0; y < 20; y++ {
			t.Log(x, y)
			if !c.GetAt(x, y).Equals(tuple.NewColor(0, 0, 0)) {
				t.Fatal("Pixel initialized to non-black")
			}
		}
	}
}

func TestWriteToCanvas(t *testing.T) {
	c := NewCanvas(10, 20)
	red := tuple.NewColor(1, 0, 0)

	c.SetAt(2, 3, red)

	if !red.Equals(c.GetAt(2, 3)) {
		t.Fatal("Pixel written does not match")
	}
}

func TestCanvasToFile(t *testing.T) {
	c := NewCanvas(5, 3)
	color1 := tuple.NewColor(1.5, 0, 0)
	color2 := tuple.NewColor(0, 0.5, 0)
	color3 := tuple.NewColor(-0.5, 0, 1)

	c.SetAt(0, 0, color1)
	c.SetAt(2, 1, color2)
	c.SetAt(4, 2, color3)

	f, err := os.CreateTemp("./", "canvasTest*.png")
	if err != nil {
		t.Fatalf("Could not create temporary file: %v", err)
	}
	c.Encode(f)
}
