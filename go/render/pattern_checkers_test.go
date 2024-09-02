package render

import (
	"testing"
)

func TestPatternCheckers_repeatInX(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0.99, 0, 0))
	c3 := pattern.At(NewPoint(1.01, 0, 0))

	if !c1.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c3.Equals(black) {
		t.Fatal("Checkers is wrong")
	}
}

func TestPatternCheckers_repeatInY(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0, 0.99, 0))
	c3 := pattern.At(NewPoint(0, 1.01, 0))

	if !c1.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c3.Equals(black) {
		t.Fatal("Checkers is wrong")
	}
}
func TestPatternCheckers_repeatInZ(t *testing.T) {
	pattern := NewCheckersPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0, 0, 0.99))
	c3 := pattern.At(NewPoint(0, 0, 1.01))

	if !c1.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Checkers is wrong")
	}
	if !c3.Equals(black) {
		t.Fatal("Checkers is wrong")
	}
}
