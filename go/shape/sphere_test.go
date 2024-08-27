package shape

import (
	"raytracerchallenge/matrix"
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
	if points[0].t != 4 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 6 {
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
	if points[0].t != 5 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 5 {
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
	if points[0].t != -1 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 1 {
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
	if points[0].t != -6 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != -4 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereTransform_default(t *testing.T) {
	s := NewSphere()

	if !s.transform.Equals(matrix.Identity) {
		t.Fatal("Default sphere transform is not the identity")
	}
}

func TestSphereTransform_changed(t *testing.T) {
	s := NewSphere()
	m := matrix.Translation(2, 3, 4)

	s.SetTransform(m)

	if !s.transform.Equals(m) {
		t.Fatal("Set sphere transform is wrong")
	}
}

func TestSphereIntersect_scaled(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(matrix.Scaling(2, 2, 2))

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != 3 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 7 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_translated(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}
