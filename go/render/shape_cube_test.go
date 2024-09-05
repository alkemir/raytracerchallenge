package render

import (
	"math"
	"testing"
)

type CubeTestCase struct {
	origin    Tuple // or point
	direction Tuple // or normal
	t1        float64
	t2        float64
}

func TestCubeIntersect_hit(t *testing.T) {
	tt := []CubeTestCase{
		{NewPoint(5, 0.5, 0), NewVector(-1, 0, 0), 4, 6},
		{NewPoint(-5, 0.5, 0), NewVector(1, 0, 0), 4, 6},
		{NewPoint(0.5, 5, 0), NewVector(0, -1, 0), 4, 6},
		{NewPoint(0.5, -5, 0), NewVector(0, 1, 0), 4, 6},
		{NewPoint(0.5, 0, 5), NewVector(0, 0, -1), 4, 6},
		{NewPoint(0.5, 0, -5), NewVector(0, 0, 1), 4, 6},
		{NewPoint(0, 0.5, 0), NewVector(0, 0, 1), -1, 1},
	}

	c := NewCube()

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction)

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

func TestCubeIntersect_miss(t *testing.T) {
	tt := []CubeTestCase{
		{NewPoint(-2, 0, 0), NewVector(0.2673, 0.5345, 0.8018), 0, 0},
		{NewPoint(0, -2, 0), NewVector(0.8018, 0.2673, 0.5345), 0, 0},
		{NewPoint(0, 0, -2), NewVector(0.5345, 0.8018, 0.2673), 0, 0},
		{NewPoint(2, 0, 2), NewVector(0, 0, -1), 0, 0},
		{NewPoint(0, 2, 2), NewVector(0, -1, 0), 0, 0},
		{NewPoint(2, 2, 0), NewVector(-1, 0, 0), 0, 0},
	}

	c := NewCube()

	for _, testCase := range tt {
		r := NewRay(testCase.origin, testCase.direction)

		ii := c.concreteIntersect(r)
		if len(ii) != 0 {
			t.Fatal("Number of intersections is wrong")
		}
	}
}

func TestCubeNormal(t *testing.T) {
	tt := []CubeTestCase{
		{NewPoint(1, 0.5, -0.8), NewVector(1, 0, 0), 0, 0},
		{NewPoint(-1, -0.2, 0.9), NewVector(-1, 0, 0), 0, 0},
		{NewPoint(-0.4, 1, -0.1), NewVector(0, 1, 0), 0, 0},
		{NewPoint(0.3, -1, -0.7), NewVector(0, -1, 0), 0, 0},
		{NewPoint(-0.6, 0.3, 1), NewVector(0, 0, 1), 0, 0},
		{NewPoint(0.4, 0.4, -1), NewVector(0, 0, -1), 0, 0},
		{NewPoint(1, 1, 1), NewVector(1, 0, 0), 0, 0},
		{NewPoint(-1, -1, -1), NewVector(-1, 0, 0), 0, 0},
	}

	c := NewCube()

	for _, testCase := range tt {
		n := c.concreteNormal(testCase.origin)

		if !n.Equals(testCase.direction) {
			t.Fatal("Normal is wrong")
		}
	}
}
