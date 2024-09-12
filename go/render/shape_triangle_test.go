package render

import (
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

	n1 := s.concreteNormal(NewPoint(0, 0.5, 0))
	n2 := s.concreteNormal(NewPoint(-0.5, 0.75, 0))
	n3 := s.concreteNormal(NewPoint(0.5, 0.25, 0))

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
