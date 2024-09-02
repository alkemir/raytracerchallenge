package render

import (
	"testing"
)

func TestPatternStripes(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	if !pattern.a.Equals(white) {
		t.Fatal("First color is wrong")
	}
	if !pattern.b.Equals(black) {
		t.Fatal("Second color is wrong")
	}
}

func TestPatternStripes_constantInY(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0, 1, 0))
	c3 := pattern.At(NewPoint(0, 2, 0))

	if !c1.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c3.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStripes_constantInZ(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0, 0, 1))
	c3 := pattern.At(NewPoint(0, 0, 2))

	if !c1.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c3.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStripes_alternatesInX(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	c1 := pattern.At(NewPoint(0, 0, 0))
	c2 := pattern.At(NewPoint(0.9, 0, 0))
	c3 := pattern.At(NewPoint(1, 0, 0))
	c4 := pattern.At(NewPoint(-0.1, 0, 0))
	c5 := pattern.At(NewPoint(-1.1, 0, 0))

	if !c1.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c2.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !c3.Equals(black) {
		t.Fatal("Stripe is wrong")
	}
	if !c4.Equals(black) {
		t.Fatal("Stripe is wrong")
	}
	if !c5.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}
