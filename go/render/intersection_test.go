package render

import (
	"math"
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
	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
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

func TestPrecompute(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(4, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if comps.t != i.t {
		t.Fatal("T is wrong")
	}
	if comps.object != shape {
		t.Fatal("Object is wrong")
	}
	if !comps.point.Equals(NewPoint(0, 0, -1)) {
		t.Fatal("Point is wrong")
	}
	if !comps.eye.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Eye is wrong")
	}
	if !comps.normal.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Normal is wrong")
	}
}

func TestPrecompute_outside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(4, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if comps.inside {
		t.Fatal("Inside was wrong")
	}
}

func TestPrecompute_inside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := NewSphere()
	i := NewIntersection(1, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if !comps.inside {
		t.Fatal("Inside was wrong")
	}
	if !comps.point.Equals(NewPoint(0, 0, 1)) {
		t.Fatal("Point is wrong")
	}
	if !comps.eye.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Eye is wrong")
	}
	if !comps.normal.Equals(NewVector(0, 0, -1)) {
		t.Fatal("Normal is wrong")
	}
}

func TestPrecompute_overPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	shape.SetTransform(Translation(0, 0, 1))
	i := NewIntersection(5, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if comps.overPoint.z >= EPSILON/2 {
		t.Fatal("Over point is wrong")
	}
	if comps.point.z <= comps.overPoint.z {
		t.Fatal("Over point is not over the point")
	}
}

func TestPrecompute_reflection(t *testing.T) {
	r := NewRay(NewPoint(0, 1, -1), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	shape := NewPlane()
	shape.SetTransform(Translation(0, 0, 1))
	i := NewIntersection(math.Sqrt2, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if !comps.reflectv.Equals(NewVector(0, math.Sqrt2/2, math.Sqrt2/2)) {
		t.Fatal("Reflective vector is wrong")
	}
}

func TestPrecompute_underPoint(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := NewSphere()
	shape.SetTransform(Translation(0, 0, 1))
	i := NewIntersection(5, shape)

	comps := i.Precompute(r, []*Intersection{i})

	if comps.underPoint.z <= EPSILON/2 {
		t.Fatal("Under point is wrong")
	}
	if comps.point.z >= comps.underPoint.z {
		t.Fatal("Under point is not under the point")
	}
}
