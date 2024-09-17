package render

import (
	"math"
)

var _ Shape = (*Cylinder)(nil)
var _ ConcreteShape = (*Cylinder)(nil)

type Cylinder struct {
	minimum float64
	maximum float64
	closed  bool
	BaseShape
}

func NewCylinder() *Cylinder {
	baseShape := *DefaultBaseShape()
	res := &Cylinder{
		minimum:   math.Inf(-1),
		maximum:   math.Inf(1),
		closed:    false,
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Cylinder) Includes(o Shape) bool {
	return s == o
}

func (s *Cylinder) SetMinimum(m float64) {
	s.minimum = m
}

func (s *Cylinder) SetMaximum(m float64) {
	s.maximum = m
}

func (s *Cylinder) SetClosed(c bool) {
	s.closed = c
}

func (s *Cylinder) concreteNormal(p Tuple, i *Intersection) Tuple {
	if !s.closed {
		return NewVector(p.x, 0, p.z)
	}

	dist := p.x*p.x + p.z*p.z
	if dist < 1 && p.y >= s.maximum-EPSILON {
		return NewVector(0, 1, 0)
	}
	if dist < 1 && p.y <= s.minimum+EPSILON {
		return NewVector(0, -1, 0)
	}

	return NewVector(p.x, 0, p.z)
}

func (s *Cylinder) concreteIntersect(tr *Ray) []*Intersection {
	ii := make([]*Intersection, 0, 2)

	a := tr.direction.x*tr.direction.x + tr.direction.z*tr.direction.z
	if math.Abs(a) >= EPSILON { // Check walls
		b := 2 * (tr.origin.x*tr.direction.x + tr.origin.z*tr.direction.z)
		c := tr.origin.x*tr.origin.x + tr.origin.z*tr.origin.z - 1

		discriminant := b*b - 4*a*c
		if discriminant >= 0 {

			t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
			t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

			if t1 < t0 {
				t0, t1 = t1, t0
			}

			y0 := tr.origin.y + t0*tr.direction.y
			if s.minimum < y0 && y0 < s.maximum {
				ii = append(ii, NewIntersection(t0, s))
			}

			y1 := tr.origin.y + t1*tr.direction.y
			if s.minimum < y1 && y1 < s.maximum {
				ii = append(ii, NewIntersection(t1, s))
			}
		}
	}

	if !s.closed || len(ii) == 2 {
		return ii
	}

	// check caps
	if t := (s.minimum - tr.origin.y) / tr.direction.y; s.checkCap(tr, t) {
		ii = append(ii, NewIntersection(t, s))
	}
	if t := (s.maximum - tr.origin.y) / tr.direction.y; s.checkCap(tr, t) {
		ii = append(ii, NewIntersection(t, s))
	}

	return ii
}

func (s *Cylinder) checkCap(tr *Ray, t float64) bool {
	x := tr.origin.x + t*tr.direction.x
	z := tr.origin.z + t*tr.direction.z

	return x*x+z*z <= 1
}
