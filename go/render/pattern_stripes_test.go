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

	if !pattern.At(NewPoint(0, 0, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(0, 1, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(0, 2, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStripes_constantInZ(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	if !pattern.At(NewPoint(0, 0, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(0, 0, 1)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(0, 0, 2)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStripes_alternatesInX(t *testing.T) {
	pattern := NewStripesPattern(white, black)

	if !pattern.At(NewPoint(0, 0, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(0.9, 0, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(1, 0, 0)).Equals(black) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(-0.1, 0, 0)).Equals(black) {
		t.Fatal("Stripe is wrong")
	}
	if !pattern.At(NewPoint(-1.1, 0, 0)).Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}
