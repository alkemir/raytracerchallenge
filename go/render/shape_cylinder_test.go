package render

import (
	"math"
	"testing"
)

type CylinderTestCase struct {
	origin    Tuple // or point
	direction Tuple // or normal
	t1        float64
	t2        float64
}

func TestCylinderIntersect_hit(t *testing.T) {
	tt := []CylinderTestCase{
		{NewPoint(1, 0, -5), NewVector(0, 0, 1), 5, 5},
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 4, 6},
		{NewPoint(0.5, 0, -5), NewVector(0.1, 1, 1), 6.80798, 7.08872},
	}

	c := NewCylinder()

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

func TestCylinderIntersect_miss(t *testing.T) {
	tt := []CylinderTestCase{
		{NewPoint(1, 0, 0), NewVector(0, 1, 0), 0, 0},
		{NewPoint(0, 0, 0), NewVector(0, 1, 0), 0, 0},
		{NewPoint(0, 0, -5), NewVector(1, 1, 1), 0, 0},
	}

	c := NewCylinder()

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction)

		ii := c.concreteIntersect(r)
		if len(ii) != 0 {
			t.Fatal("Number of intersections is wrong")
		}
	}
}

func TestCylinderNormal(t *testing.T) {
	tt := []CylinderTestCase{
		{NewPoint(1, 0, 0), NewVector(1, 0, 0), 0, 0},
		{NewPoint(0, 5, -1), NewVector(0, 0, -1), 0, 0},
		{NewPoint(0, -2, 1), NewVector(0, 0, 1), 0, 0},
		{NewPoint(-1, 1, 0), NewVector(-1, 0, 0), 0, 0},
	}

	c := NewCylinder()

	for _, testCase := range tt {
		n := c.concreteNormal(testCase.origin)

		if !n.Equals(testCase.direction) {
			t.Fatal("Normal is wrong")
		}
	}
}

func TestCylinderDefaultLength(t *testing.T) {
	c := NewCylinder()

	if c.minimum != math.Inf(-1) {
		t.Fatal("Minimum is wrong")
	}
	if c.maximum != math.Inf(1) {
		t.Fatal("Minimum is wrong")
	}
}

func TestCylinderIntersect_constrained(t *testing.T) {
	tt := []CylinderTestCase{
		{NewPoint(0, 1.5, 0), NewVector(0.1, 1, 0), 0, 0},
		{NewPoint(0, 3, -5), NewVector(0, 0, 1), 0, 0},
		{NewPoint(0, 0, -5), NewVector(0, 0, 1), 0, 0},
		{NewPoint(0, 2, -5), NewVector(0, 0, 1), 0, 0},
		{NewPoint(0, 1, -5), NewVector(0, 0, 1), 0, 0},
		{NewPoint(0, 1.5, -2), NewVector(0, 0, 1), 2, 0},
	}

	c := NewCylinder()
	c.SetMinimum(1)
	c.SetMaximum(2)

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction)

		ii := c.concreteIntersect(r)
		if len(ii) != int(testCase.t1) {
			t.Fatal("Number of intersections is wrong")
		}
	}
}
