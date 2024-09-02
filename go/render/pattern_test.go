package render

import (
	"testing"
)

var black = NewColor(0, 0, 0)
var white = NewColor(1, 1, 1)

func TestPatternStripe(t *testing.T) {
	pattern := NewStripePattern(white, black)

	if !pattern.a.Equals(white) {
		t.Fatal("First color is wrong")
	}
	if !pattern.b.Equals(black) {
		t.Fatal("Second color is wrong")
	}
}

func TestPatternStrip_constantInY(t *testing.T) {
	pattern := NewStripePattern(white, black)

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

func TestPatternStrip_constantInZ(t *testing.T) {
	pattern := NewStripePattern(white, black)

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

func TestPatternStrip_alternatesInX(t *testing.T) {
	pattern := NewStripePattern(white, black)

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

func TestPatternStrip_objectTransform(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	pattern := NewStripePattern(white, black)

	c := pattern.AtObject(obj, NewPoint(1.50, 0, 0))

	if !c.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStrip_patternTransform(t *testing.T) {
	obj := NewSphere()
	pattern := NewStripePattern(white, black)
	pattern.SetTransform(Scaling(2, 2, 2))

	c := pattern.AtObject(obj, NewPoint(1.50, 0, 0))

	if !c.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}

func TestPatternStrip_patternAndObjTransform(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	pattern := NewStripePattern(white, black)
	pattern.SetTransform(Translation(0.5, 0, 0))

	c := pattern.AtObject(obj, NewPoint(1.50, 0, 0))

	if !c.Equals(white) {
		t.Fatal("Stripe is wrong")
	}
}
