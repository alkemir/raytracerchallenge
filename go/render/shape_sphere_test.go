package render

import (
	"math"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
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
	r := NewRay(NewPoint(0, 1, -5), NewVector(0, 0, 1))
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
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestSphereIntersect_inside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
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
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
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

func TestSphereIntersect_scaled(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(Scaling(2, 2, 2))

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
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(Translation(5, 0, 0))

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestSphereNormal_xaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(NewPoint(1, 0, 0))

	if !n.Equals(NewVector(1, 0, 0)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSphereNormal_yaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(NewPoint(0, 1, 0))

	if !n.Equals(NewVector(0, 1, 0)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_zaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(NewPoint(0, 0, 1))

	if !n.Equals(NewVector(0, 0, 1)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_offAxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSphereNormal_normalized(t *testing.T) {
	s := NewSphere()

	n := s.Normal(NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(n.Norm()) {
		t.Fatal("Normal is not normal")
	}
}

func TestSphereNormal_translated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Translation(0, 1, 0))

	n := s.Normal(NewPoint(0, 1.70710678, -0.70710678))

	if !n.Equals(NewVector(0, 0.70710678, -0.70710678)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_scaledRotated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(Scaling(1, 0.5, 1).Multiply(RotationZ(math.Pi / 5)))

	n := s.Normal(NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))

	if !n.Equals(NewVector(0, 0.970142500, -0.24253562)) {
		t.Fatal("Normal is wrong")
	}
}

func TestGlassSphere(t *testing.T) {
	s := NewGlassSphere()

	if !s.transform().Equals(IdentityMatrix()) {
		t.Fatal("Transform is wrong")
	}
	if s.material().transparency != 1 {
		t.Fatal("Transparency is wrong")
	}
	if s.material().refractiveIndex != 1.5 {
		t.Fatal("Refractive index is wrong")
	}
}

func TestGlassSphere_indexes(t *testing.T) {
	a := NewGlassSphere()
	a.SetTransform(Scaling(2, 2, 2))
	a.material().refractiveIndex = 1.5

	b := NewGlassSphere()
	b.SetTransform(Translation(0, 0, -0.25))
	b.material().refractiveIndex = 2

	c := NewGlassSphere()
	c.SetTransform(Translation(0, 0, 0.25))
	c.material().refractiveIndex = 2.5

	r := NewRay(NewPoint(0, 0, -4), NewVector(0, 0, 1))
	ii := []*Intersection{
		NewIntersection(2, a),
		NewIntersection(2.75, b),
		NewIntersection(3.25, c),
		NewIntersection(4.75, b),
		NewIntersection(5.25, c),
		NewIntersection(6, a),
	}

	var comps *Comps

	comps = ii[0].Precompute(r, ii)
	if comps.n1 != 1.0 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 1.5 {
		t.Fatal("n2 is wrong")
	}

	comps = ii[1].Precompute(r, ii)
	if comps.n1 != 1.5 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 2.0 {
		t.Fatal("n2 is wrong")
	}

	comps = ii[2].Precompute(r, ii)
	if comps.n1 != 2.0 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 2.5 {
		t.Fatal("n2 is wrong")
	}

	comps = ii[3].Precompute(r, ii)
	if comps.n1 != 2.5 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 2.5 {
		t.Fatal("n2 is wrong")
	}

	comps = ii[4].Precompute(r, ii)
	if comps.n1 != 2.5 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 1.5 {
		t.Fatal("n2 is wrong")
	}

	comps = ii[5].Precompute(r, ii)
	if comps.n1 != 1.5 {
		t.Fatal("n1 is wrong")
	}
	if comps.n2 != 1.0 {
		t.Fatal("n2 is wrong")
	}
}
