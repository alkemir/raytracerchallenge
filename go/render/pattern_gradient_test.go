package render

import (
	"testing"
)

func TestPatternGradient(t *testing.T) {
	pattern := NewGradientPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0.25, 0, 0))
	c3 := pattern.At(NewPoint(0.5, 0, 0))
	c4 := pattern.At(NewPoint(0.75, 0, 0))

	if !c1.Equals(white) {
		t.Fatal("Gradient is wrong")
	}
	if !c2.Equals(NewColor(0.75, 0.75, 0.75)) {
		t.Fatal("Gradientis wrong")
	}
	if !c3.Equals(NewColor(0.5, 0.5, 0.5)) {
		t.Fatal("Gradient is wrong")
	}
	if !c4.Equals(NewColor(0.25, 0.25, 0.25)) {
		t.Fatal("Gradientis wrong")
	}
}
