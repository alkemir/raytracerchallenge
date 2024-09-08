package render

import (
	"math"
	"testing"
)

type ConeTestCase struct {
	origin    Tuple // or point
	direction Tuple // or normal
	t1        float64
	t2        float64
}

func TestConeIntersect_hit(t *testing.T) {
	tt := []ConeTestCase{
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 5, 5},
		{NewPoint(0, 0, -5), NewVector(1, 1, 1), 8.66025, 8.66025},
		{NewPoint(1, 1, -5), NewVector(-0.5, -1, 1), 4.55006, 49.44994},
	}

	c := NewCone()

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction.Norm())

		ii := c.concreteIntersect(r)
		if len(ii) != 2 {
			t.Fatal("Number of intersections is wrong")
		}
		if math.Abs(ii[0].t-testCase.t1) > EPSILON {
			t.Fatal("First intersection is wrong")
		}
		if math.Abs(ii[1].t-testCase.t2) > EPSILON {
			t.Fatal("Second intersection is wrong")
		}
	}
}

func TestConeIntersect_parallel(t *testing.T) {
	c := NewCone()

	r := NewRay(NewPoint(0, 0, -1), NewVector(0, 1, 1).Norm())

	ii := c.concreteIntersect(r)
	if len(ii) != 1 {
		t.Fatal("Number of intersections is wrong")
	}
	if math.Abs(ii[0].t-0.35355) > EPSILON {
		t.Fatal("t of intersection is wrong")
	}
}

func TestConeNormal(t *testing.T) {
	tt := []ConeTestCase{
		{NewPoint(0, 0, 0), NewVector(0, 0, 0), 0, 0},
		{NewPoint(1, 1, 1), NewVector(1, -math.Sqrt2, 1), 0, 0},
		{NewPoint(-1, -1, 0), NewVector(-1, 1, 0), 0, 0},
	}

	c := NewCone()

	for _, testCase := range tt {
		n := c.concreteNormal(testCase.origin)
		if !n.Equals(testCase.direction) {
			t.Fatal("Normal is wrong")
		}
	}
}

func TestConeNormal_caps(t *testing.T) {
	tt := []CylinderTestCase{
		{NewPoint(0, 1, 0), NewVector(0, -1, 0), 0, 0},
		{NewPoint(0, 2, 0), NewVector(0, 1, 0), 0, 0},
	}

	c := NewCylinder()
	c.SetMinimum(1)
	c.SetMaximum(2)
	c.SetClosed(true)

	for _, testCase := range tt {
		n := c.concreteNormal(testCase.origin)

		if !n.Equals(testCase.direction) {
			t.Fatal("Normal is wrong")
		}
	}
}

func TestConeDefaultLength(t *testing.T) {
	c := NewCone()

	if c.minimum != math.Inf(-1) {
		t.Fatal("Minimum is wrong")
	}
	if c.maximum != math.Inf(1) {
		t.Fatal("Minimum is wrong")
	}
}

func TestConeDefaultClosed(t *testing.T) {
	c := NewCone()

	if c.closed {
		t.Fatal("Default for closed is wrong")
	}
}

func TestConeIntersect_caps(t *testing.T) {
	tt := []ConeTestCase{
		{NewPoint(0, 0, -5), NewVector(0, 1, 0), 0, 0},
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 1), 2, 0},
		{NewPoint(0, 0, -0.25), NewVector(0, 1, 0), 4, 0},
	}

	c := NewCone()
	c.SetMinimum(-0.5)
	c.SetMaximum(0.5)
	c.SetClosed(true)

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction)

		ii := c.concreteIntersect(r)
		if len(ii) != int(testCase.t1) {
			t.Fatal("Number of intersections is wrong")
		}
	}
}
