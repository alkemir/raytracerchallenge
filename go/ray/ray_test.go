package ray

import (
	"raytracerchallenge/matrix"
	"raytracerchallenge/tuple"
	"testing"
)

func TestRayConstructor(t *testing.T) {
	origin := tuple.NewPoint(1, 2, 3)
	direction := tuple.NewVector(4, 5, 6)

	r := NewRay(origin, direction)

	if !origin.Equals(r.origin) {
		t.Fatal("Origin was not created right")
	}
	if !direction.Equals(r.direction) {
		t.Fatal("Direction was not created right")
	}
}

func TestRayProject(t *testing.T) {
	r := NewRay(tuple.NewPoint(2, 3, 4), tuple.NewVector(1, 0, 0))

	p1 := r.Project(0)
	p2 := r.Project(1)
	p3 := r.Project(-1)
	p4 := r.Project(2.5)

	if !p1.Equals(tuple.NewPoint(2, 3, 4)) {
		t.Fatal("Projection is wrong")
	}
	if !p2.Equals(tuple.NewPoint(3, 3, 4)) {
		t.Fatal("Projection is wrong")
	}
	if !p3.Equals(tuple.NewPoint(1, 3, 4)) {
		t.Fatal("Projection is wrong")
	}
	if !p4.Equals(tuple.NewPoint(4.5, 3, 4)) {
		t.Fatal("Projection is wrong")
	}
}

func TestRayTranslate(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	m := matrix.Translation(3, 4, 5)

	r2 := r.Transform(m)

	if !r2.origin.Equals(tuple.NewPoint(4, 6, 8)) {
		t.Fatal("Translated origin is wrong")
	}
	if !r2.direction.Equals(tuple.NewVector(0, 1, 0)) {
		t.Fatal("Translated direction is wrong")
	}
}

func TestRayScaling(t *testing.T) {
	r := NewRay(tuple.NewPoint(1, 2, 3), tuple.NewVector(0, 1, 0))
	m := matrix.Scaling(2, 3, 4)

	r2 := r.Transform(m)

	if !r2.origin.Equals(tuple.NewPoint(2, 6, 12)) {
		t.Fatal("Scaled origin is wrong")
	}
	if !r2.direction.Equals(tuple.NewVector(0, 3, 0)) {
		t.Fatal("Scaled direction is wrong")
	}
}
