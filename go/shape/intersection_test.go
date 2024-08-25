package shape

import (
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
	"testing"
)

func TestIntersectionConstructor(t *testing.T) {
	s := NewSphere()
	i := NewIntersection(3.5, s)

	if i.t != 3.5 {
		t.Fatal("Intersection t is wrong")
	}
	if i.obj != s {
		t.Fatal("Intersection object is wrong")
	}
}

func TestIntersectionCollection(t *testing.T) {
	s := NewSphere()
	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)

	ii := []*Intersection{i1, i2}

	if len(ii) != 2 {
		t.Fatal("Intersections length is wrong")
	}
	if ii[0].t != 1 {
		t.Fatal("Intersection t is wrong")
	}
	if ii[1].t != 2 {
		t.Fatal("Intersection t is wrong")
	}
}

func TestSphereIntersectObject(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].obj != s {
		t.Fatal("Wrong intersection")
	}
	if points[1].obj != s {
		t.Fatal("Wrong intersection")
	}
}

func TestHit_positive(t *testing.T) {
	s := NewSphere()

	i1 := NewIntersection(1, s)
	i2 := NewIntersection(2, s)
	ii := []*Intersection{i2, i1}

	h := Hit(ii)

	if h != i1 {
		t.Fatal("Wrong intersection selected as hit")
	}
}

func TestHit_mixed(t *testing.T) {
	s := NewSphere()

	i1 := NewIntersection(-1, s)
	i2 := NewIntersection(1, s)
	ii := []*Intersection{i2, i1}

	h := Hit(ii)

	if h != i2 {
		t.Fatal("Wrong intersection selected as hit")
	}
}

func TestHit_negative(t *testing.T) {
	s := NewSphere()

	i1 := NewIntersection(-2, s)
	i2 := NewIntersection(-1, s)
	ii := []*Intersection{i2, i1}

	h := Hit(ii)

	if h != nil {
		t.Fatal("Wrong intersection selected as hit")
	}
}

func TestHit_many(t *testing.T) {
	s := NewSphere()

	i1 := NewIntersection(5, s)
	i2 := NewIntersection(7, s)
	i3 := NewIntersection(-3, s)
	i4 := NewIntersection(2, s)
	ii := []*Intersection{i1, i2, i3, i4}

	h := Hit(ii)

	if h != i4 {
		t.Fatal("Wrong intersection selected as hit")
	}
}
