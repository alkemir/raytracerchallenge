package render

import (
	"math"
	"testing"
)

func TestTriangleConstructor(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)

	if !s.p1.Equals(p1) {
		t.Fatal("P1 is wrong")
	}
	if !s.p2.Equals(p2) {
		t.Fatal("P2 is wrong")
	}
	if !s.p3.Equals(p3) {
		t.Fatal("P3 is wrong")
	}
	if !s.e1.Equals(NewVector(-1, -1, 0)) {
		t.Fatal("E1 is wrong")
	}
	if !s.e2.Equals(NewVector(1, -1, 0)) {
		t.Fatal("E2 is wrong")
	}
	if !s.normal.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Normal is wrong")
	}
}

func TestTriangleNormal(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)

	n1 := s.concreteNormal(NewPoint(0, 0.5, 0), nil)
	n2 := s.concreteNormal(NewPoint(-0.5, 0.75, 0), nil)
	n3 := s.concreteNormal(NewPoint(0.5, 0.25, 0), nil)

	if !s.normal.Equals(n1) {
		t.Fatal("Normal is wrong")
	}
	if !s.normal.Equals(n2) {
		t.Fatal("Normal is wrong")
	}
	if !s.normal.Equals(n3) {
		t.Fatal("Normal is wrong")
	}
}

func TestTriangleIntersect_parallel(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 1, 0))

	ii := s.concreteIntersect(r)

	if len(ii) != 0 {
		t.Fatal("Intersect is wrong")
	}
}

func TestTriangleIntersect_missA(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)
	r := NewRay(NewPoint(1, 1, -2), NewVector(0, 0, 1))

	ii := s.concreteIntersect(r)

	if len(ii) != 0 {
		t.Fatal("Intersect is wrong")
	}
}

func TestTriangleIntersect_missB(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)
	r := NewRay(NewPoint(-1, 1, -2), NewVector(0, 0, 1))

	ii := s.concreteIntersect(r)

	if len(ii) != 0 {
		t.Fatal("Intersect is wrong")
	}
}

func TestTriangleIntersect_missC(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)
	r := NewRay(NewPoint(0, -1, -2), NewVector(0, 0, 1))

	ii := s.concreteIntersect(r)

	if len(ii) != 0 {
		t.Fatal("Intersect is wrong")
	}
}

func TestTriangleIntersect_hit(t *testing.T) {
	p1 := NewPoint(0, 1, 0)
	p2 := NewPoint(-1, 0, 0)
	p3 := NewPoint(1, 0, 0)
	s := NewTriangle(p1, p2, p3)
	r := NewRay(NewPoint(0, 0.5, -2), NewVector(0, 0, 1))

	ii := s.concreteIntersect(r)

	if len(ii) != 1 {
		t.Fatal("Intersect is wrong")
	}
	if math.Abs(ii[0].t-2) > EPSILON {
		t.Fatal("Intersect is wrong")
	}
}
