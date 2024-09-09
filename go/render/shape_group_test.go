package render

import (
	"testing"
)

func TestEmptyGroup(t *testing.T) {
	s := NewGroup()

	if !s.transform().Equals(IdentityMatrix()) {
		t.Fatal("Transform is wrong")
	}
	if len(s.Children()) != 0 {
		t.Fatal("Number of children is wrong")
	}
}

func TestRootShape(t *testing.T) {
	s := NewTestShape()

	if s.parent != nil {
		t.Fatal("Parent is wrong")
	}
}

func TestGroupAdd(t *testing.T) {
	g := NewGroup()
	s := NewTestShape()

	g.Add(s)

	if g.Children()[0] != s {
		t.Fatal("Children is wrong")
	}
	if s.parent == nil {
		t.Fatal("Parent is wrong")
	}
}

func TestGroupIntersect_empty(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	s := NewGroup()

	points := s.concreteIntersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestGroupIntersect_hit(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	g := NewGroup()

	s1 := NewSphere()
	g.Add(s1)
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, -3))
	g.Add(s2)
	s3 := NewSphere()
	s3.SetTransform(Translation(5, 0, 0))
	g.Add(s3)

	points := g.concreteIntersect(r)

	if len(points) != 4 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].obj != s2 {
		t.Fatal("Intersected object is wrong")
	}
	if points[1].obj != s2 {
		t.Fatal("Intersected object is wrong")
	}
	if points[2].obj != s1 {
		t.Fatal("Intersected object is wrong")
	}
	if points[3].obj != s1 {
		t.Fatal("Intersected object is wrong")
	}
}

func TestGroupIntersect_transformed(t *testing.T) {
	r := NewRay(NewPoint(10, 0, -10), NewVector(0, 0, 1))
	g := NewGroup()
	g.SetTransform(Scaling(2, 2, 2))
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g.Add(s)

	points := g.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
}
