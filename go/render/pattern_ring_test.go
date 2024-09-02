package render

import (
	"testing"
)

func TestPatternRing(t *testing.T) {
	pattern := NewRingPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(1, 0, 0))
	c3 := pattern.At(NewPoint(0, 0, 1))
	c4 := pattern.At(NewPoint(0.708, 0, 0.708))

	if !c1.Equals(white) {
		t.Fatal("First color is wrong")
	}
	if !c2.Equals(black) {
		t.Fatal("First color is wrong")
	}
	if !c3.Equals(black) {
		t.Fatal("First color is wrong")
	}
	if !c4.Equals(black) {
		t.Fatal("First color is wrong")
	}
}
