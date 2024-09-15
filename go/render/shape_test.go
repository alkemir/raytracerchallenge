package render

import (
	"math"
	"testing"
)

func TestShapeDefaultTransform(t *testing.T) {
	s := NewTestShape()

	if !s.transform().Equals(IdentityMatrix()) {
		t.Fatal("Default transform is wrong")
	}
}

func TestShapeSetTransform(t *testing.T) {
	s := NewTestShape()

	s.SetTransform(Translation(2, 3, 4))

	if !s.transform().Equals(Translation(2, 3, 4)) {
		t.Fatal("Default transform is wrong")
	}
}

func TestShapeDefaultMaterial(t *testing.T) {
	s := NewTestShape()

	if !s.material().Equals(DefaultMaterial()) {
		t.Fatal("Default transform is wrong")
	}
}

func TestShapeSetMaterial(t *testing.T) {
	s := NewTestShape()
	m := DefaultMaterial()
	m.ambient = 1

	s.SetMaterial(m)

	if !s.material().Equals(m) {
		t.Fatal("Default transform is wrong")
	}
}

func TestShapeIntersectScaled(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()

	s.SetTransform(Scaling(2, 2, 2))
	s.Intersect(r)

	if !s.savedRay.origin.Equals(NewPoint(0, 0, -2.5)) {
		t.Fatal("Ray origin was wrong")
	}
	if !s.savedRay.direction.Equals(NewVector(0, 0, 0.5)) {
		t.Fatal("Ray origin was wrong")
	}
}

func TestShapeIntersectTranslated(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewTestShape()

	s.SetTransform(Translation(5, 0, 0))
	s.Intersect(r)

	if !s.savedRay.origin.Equals(NewPoint(-5, 0, -5)) {
		t.Fatal("Ray origin was wrong")
	}
	if !s.savedRay.direction.Equals(NewVector(0, 0, 1)) {
		t.Fatal("Ray origin was wrong")
	}
}

func TestShapeNormalTranslated(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(Translation(0, 1, 0))

	n := s.Normal(NewPoint(0, 1.70711, -0.70711), nil)

	if !n.Equals(NewVector(0, 0.70711, -0.70711)) {
		t.Fatal("Normal is wrong")
	}
}

func TestShapeNormalScaled(t *testing.T) {
	s := NewTestShape()
	s.SetTransform(Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5)))

	n := s.Normal(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2), nil)

	if !n.Equals(NewVector(0, 0.97014, -0.24254)) {
		t.Fatal("Normal is wrong")
	}
}
