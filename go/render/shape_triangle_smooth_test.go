package render

import (
	"math"
	"testing"
)

func TestSmoothTriangleConstructor(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)

	if !tri.p1.Equals(p1) {
		t.Fatal("Point is wrong")
	}
	if !tri.p2.Equals(p2) {
		t.Fatal("Point is wrong")
	}
	if !tri.p3.Equals(p3) {
		t.Fatal("Point is wrong")
	}
	if !tri.n1.Equals(n1) {
		t.Fatal("Point is wrong")
	}
	if !tri.n2.Equals(n2) {
		t.Fatal("Point is wrong")
	}
	if !tri.n3.Equals(n3) {
		t.Fatal("Point is wrong")
	}
}

func TestSmoothTriangleIntersection_storesUV(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)
	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVector(0, 0, 1))

	ii := tri.concreteIntersect(r)

	if math.Abs(ii[0].u-0.45) > EPSILON {
		t.Fatal("U is wrong")
	}
	if math.Abs(ii[0].v-0.25) > EPSILON {
		t.Fatal("V is wrong")
	}
}

func TestSmoothTriangleNormal_interpolated(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)
	i := NewIntersectionUV(1, tri, 0.45, 0.25)
	n := tri.Normal(NewPoint(0, 0, 0), i)

	if !n.Equals(NewVector(-0.5547, 0.83205, 0)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSmoothTriangleNormal_computed(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	n1 := NewVector(0, 1, 0)
	n2 := NewVector(-1, 0, 0)
	n3 := NewVector(1, 0, 0)

	tri := NewSmoothTriangle(p1, p2, p3, n1, n2, n3)
	i := NewIntersectionUV(1, tri, 0.45, 0.25)
	r := NewRay(NewPoint(-0.2, 0.3, -2), NewVector(0, 0, 1))
	ii := []*Intersection{i}
	comps := i.Precompute(r, ii)

	if !comps.normal.Equals(NewVector(-0.5547, 0.83205, 0)) {
		t.Fatal("Normal is wrong")
	}
}
