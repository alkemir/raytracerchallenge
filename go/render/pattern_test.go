package render

import (
	"testing"
)

var black = NewColor(0, 0, 0)
var white = NewColor(1, 1, 1)

func TestPattern_defaultTransform(t *testing.T) {
	pattern := NewTestPattern()

	if !pattern.transform.Equals(IdentityMatrix()) {
		t.Fatal("Transform is wrong")
	}
}

func TestPattern_setTransform(t *testing.T) {
	pattern := NewTestPattern()
	pattern.SetTransform(Translation(1, 2, 3))

	if !pattern.transform.Equals(Translation(1, 2, 3)) {
		t.Fatal("Transform is wrong")
	}
}

func TestPattern_objectTransform(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	pattern := NewTestPattern()

	c := pattern.AtObject(obj, NewPoint(2, 3, 4))

	if !c.Equals(NewColor(1, 1.5, 2)) {
		t.Fatal("Transform is wrong is wrong")
	}
}

func TestPattern_patternTransform(t *testing.T) {
	obj := NewSphere()
	pattern := NewTestPattern()
	pattern.SetTransform(Scaling(2, 2, 2))

	c := pattern.AtObject(obj, NewPoint(2, 3, 4))

	if !c.Equals(NewColor(1, 1.5, 2)) {
		t.Fatal("Transform is wrong is wrong")
	}
}

func TestPattern_patternAndObjTransform(t *testing.T) {
	obj := NewSphere()
	obj.SetTransform(Scaling(2, 2, 2))
	pattern := NewTestPattern()
	pattern.SetTransform(Translation(0.5, 1, 1.5))

	c := pattern.AtObject(obj, NewPoint(2.5, 3, 3.5))

	if !c.Equals(NewColor(0.75, 0.5, 0.25)) {
		t.Fatal("Stripe is wrong")
	}
}
