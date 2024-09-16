package render

import (
	"math"
	"testing"
)

func TestCSGShapeConstructor(t *testing.T) {
	s1 := NewSphere()
	s2 := NewCube()

	c := NewCSGShape(CSG_UNION, s1, s2)

	if c.op != CSG_UNION {
		t.Fatal("Operation is wrong")
	}
	if c.left != s1 {
		t.Fatal("Left is wrong")
	}
	if c.right != s2 {
		t.Fatal("Right is wrong")
	}
	if s1.parent != c {
		t.Fatal("Left parent is wrong")
	}
	if s2.parent != c {
		t.Fatal("Right parent is wrong")
	}
}

type OperationTestCase struct {
	op     CSGOperation
	lHit   bool
	inL    bool
	inR    bool
	result bool
}

func TestCSGOperations_intersectionAllowed(t *testing.T) {
	testCases := []OperationTestCase{
		{CSG_UNION, true, true, true, false},
		{CSG_UNION, true, true, false, true},
		{CSG_UNION, true, false, true, false},
		{CSG_UNION, true, false, false, true},
		{CSG_UNION, false, true, true, false},
		{CSG_UNION, false, true, false, false},
		{CSG_UNION, false, false, true, true},
		{CSG_UNION, false, false, false, true},
		{CSG_INTERSECTION, true, true, true, true},
		{CSG_INTERSECTION, true, true, false, false},
		{CSG_INTERSECTION, true, false, true, true},
		{CSG_INTERSECTION, true, false, false, false},
		{CSG_INTERSECTION, false, true, true, true},
		{CSG_INTERSECTION, false, true, false, true},
		{CSG_INTERSECTION, false, false, true, false},
		{CSG_INTERSECTION, false, false, false, false},
		{CSG_DIFFERENCE, true, true, true, false},
		{CSG_DIFFERENCE, true, true, false, true},
		{CSG_DIFFERENCE, true, false, true, false},
		{CSG_DIFFERENCE, true, false, false, true},
		{CSG_DIFFERENCE, false, true, true, true},
		{CSG_DIFFERENCE, false, true, false, true},
		{CSG_DIFFERENCE, false, false, true, false},
		{CSG_DIFFERENCE, false, false, false, false},
	}

	for _, testCase := range testCases {
		result := testCase.op.intersectionAllowed(testCase.lHit, testCase.inL, testCase.inR)
		if result != testCase.result {
			t.Fatal("Result is wrong")
		}
	}
}

type FilterTestCase struct {
	op CSGOperation
	x0 int
	x1 int
}

func TestCSGShapeFilterIntersections(t *testing.T) {
	s1 := NewSphere()
	s2 := NewCube()

	testCases := []FilterTestCase{
		{CSG_UNION, 0, 3},
		{CSG_INTERSECTION, 1, 2},
		{CSG_DIFFERENCE, 0, 1},
	}

	for _, testCase := range testCases {
		c := NewCSGShape(testCase.op, s1, s2)
		ii := []*Intersection{
			NewIntersection(1, s1),
			NewIntersection(2, s2),
			NewIntersection(3, s1),
			NewIntersection(4, s2),
		}

		result := c.filterIntersection(ii)

		if len(result) != 2 {
			t.Fatal("The number of results is wrong")
		}
		if result[0] != ii[testCase.x0] {
			t.Fatal("The first result is wrong")
		}
		if result[1] != ii[testCase.x1] {
			t.Fatal("The second result is wrong")
		}
	}
}

func TestCSGShapeIntersect_miss(t *testing.T) {
	c := NewCSGShape(CSG_UNION, NewSphere(), NewCube())
	r := NewRay(NewPoint(0, 2, -5), NewVector(0, 0, 1))

	ii := c.concreteIntersect(r)

	if len(ii) != 0 {
		t.Fatal("The number of results is wrong")
	}
}

func TestCSGShapeIntersect_hit(t *testing.T) {
	s1 := NewSphere()
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, 0.5))
	c := NewCSGShape(CSG_UNION, s1, s2)
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	ii := c.concreteIntersect(r)

	if len(ii) != 2 {
		t.Fatal("The number of results is wrong")
	}
	if math.Abs(ii[0].t-4) > EPSILON {
		t.Fatal("T is wrong")
	}
	if ii[0].obj != s1 {
		t.Fatal("Object is wrong")
	}
	if math.Abs(ii[1].t-6.5) > EPSILON {
		t.Fatal("T is wrong")
	}
	if ii[1].obj != s2 {
		t.Fatal("Object is wrong")
	}
}
