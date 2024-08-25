package ray

import (
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
