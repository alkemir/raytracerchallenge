package render

import (
	"testing"
)

func TestPlaneNormal(t *testing.T) {
	p := NewPlane()

	n1 := p.concreteNormal(NewPoint(0, 0, 0), nil)
	n2 := p.concreteNormal(NewPoint(10, 0, -10), nil)
	n3 := p.concreteNormal(NewPoint(-5, 0, 150), nil)

	if !n1.Equals(NewVector(0, 1, 0)) {
		t.Fatal("Normal is wrong")
	}
	if !n2.Equals(NewVector(0, 1, 0)) {
		t.Fatal("Normal is wrong")
	}
	if !n3.Equals(NewVector(0, 1, 0)) {
		t.Fatal("Normal is wrong")
	}
}

func TestPlaneIntersect_parallel(t *testing.T) {
	p := NewPlane()

	ii := p.concreteIntersect(NewRay(NewPoint(0, 10, 0), NewVector(0, 0, 1)))

	if len(ii) != 0 {
		t.Fatal("Intersections are wrong")
	}
}

func TestPlaneIntersect_coplanar(t *testing.T) {
	p := NewPlane()

	ii := p.concreteIntersect(NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1)))

	if len(ii) != 0 {
		t.Fatal("Intersections are wrong")
	}
}

func TestPlaneIntersect_above(t *testing.T) {
	p := NewPlane()

	ii := p.concreteIntersect(NewRay(NewPoint(0, 1, 0), NewVector(0, -1, 0)))

	if len(ii) != 1 {
		t.Fatal("Intersections are wrong")
	}
	if ii[0].t != 1 {
		t.Fatal("Intersected object distance is wrong")
	}
	if ii[0].obj != p {
		t.Fatal("Intersected object is wrong")
	}
}

func TestPlaneIntersect_below(t *testing.T) {
	p := NewPlane()

	ii := p.concreteIntersect(NewRay(NewPoint(0, -1, 0), NewVector(0, 1, 0)))

	if len(ii) != 1 {
		t.Fatal("Intersections are wrong")
	}
	if ii[0].t != 1 {
		t.Fatal("Intersected object distance is wrong")
	}
	if ii[0].obj != p {
		t.Fatal("Intersected object is wrong")
	}
}
