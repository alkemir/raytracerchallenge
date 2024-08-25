package shape

import (
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0] != 4 {
		t.Fatal("Wrong intersection")
	}
	if points[1] != 6 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_tangent(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0] != 5 {
		t.Fatal("Wrong intersection")
	}
	if points[1] != 5 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_miss(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestSphereIntersect_inside(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0] != -1 {
		t.Fatal("Wrong intersection")
	}
	if points[1] != 1 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_behind(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0] != -6 {
		t.Fatal("Wrong intersection")
	}
	if points[1] != -4 {
		t.Fatal("Wrong intersection")
	}
}
