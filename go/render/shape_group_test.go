package render

import (
	"math"
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

func TestGroupWorldToObject(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(2, 2, 2))
	g1.Add(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.Add(s)

	p := s.worldToObject(NewPoint(-2, 0, -10))

	if !p.Equals(NewPoint(0, 0, -1)) {
		t.Fatal("Point is wrong")
	}
}

func TestGroupNormalToWorld(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.Add(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.Add(s)

	n := s.normalToWorld(NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(NewVector(0.285714, 0.42857, -0.85714)) {
		t.Fatal("Normal is wrong")
	}
}

func TestGroupNormal(t *testing.T) {
	g1 := NewGroup()
	g1.SetTransform(RotationY(math.Pi / 2))
	g2 := NewGroup()
	g2.SetTransform(Scaling(1, 2, 3))
	g1.Add(g2)
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))
	g2.Add(s)

	n := s.Normal(NewPoint(1.7321, 1.1547, -5.5774), nil)

	if !n.Equals(NewVector(0.2857036, 0.4285431, -0.85716)) {
		t.Fatal("Normal is wrong")
	}
}
