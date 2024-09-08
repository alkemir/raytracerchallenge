package render

import (
	"math"
)

var _ Shape = (*Cone)(nil)
var _ ConcreteShape = (*Cone)(nil)

type Cone struct {
	minimum float64
	maximum float64
	closed  bool
	BaseShape
}

func NewCone() *Cone {
	baseShape := *DefaultBaseShape()
	res := &Cone{
		minimum:   math.Inf(-1),
		maximum:   math.Inf(1),
		closed:    false,
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Cone) SetMinimum(m float64) {
	s.minimum = m
}

func (s *Cone) SetMaximum(m float64) {
	s.maximum = m
}

func (s *Cone) SetClosed(c bool) {
	s.closed = c
}

func (s *Cone) concreteNormal(p Tuple) Tuple {
	y := math.Sqrt(p.x*p.x + p.z*p.z)
	if p.y > 0 {
		y = -y
	}

	if !s.closed {
		return NewVector(p.x, y, p.z)
	}

	dist := p.x*p.x + p.z*p.z
	if dist < y && p.y >= s.maximum-EPSILON {
		return NewVector(0, 1, 0)
	}
	if dist < y && p.y <= s.minimum+EPSILON {
		return NewVector(0, -1, 0)
	}

	return NewVector(p.x, y, p.z)
}

func (s *Cone) concreteIntersect(tr *Ray) []*Intersection {
	ii := make([]*Intersection, 0)

	a := tr.direction.x*tr.direction.x - tr.direction.y*tr.direction.y + tr.direction.z*tr.direction.z
	b := 2 * (tr.origin.x*tr.direction.x - tr.origin.y*tr.direction.y + tr.origin.z*tr.direction.z)
	c := tr.origin.x*tr.origin.x + -tr.origin.y*tr.origin.y + tr.origin.z*tr.origin.z
	if math.Abs(a) >= EPSILON { // Check walls
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
	} else if math.Abs(b) >= EPSILON {
		ii = append(ii, NewIntersection(-c/2/b, s))
	}

	if !s.closed {
		return ii
	}

	// check caps
	if t := (s.minimum - tr.origin.y) / tr.direction.y; s.checkCap(tr, t, s.minimum) {
		ii = append(ii, NewIntersection(t, s))
	}
	if t := (s.maximum - tr.origin.y) / tr.direction.y; s.checkCap(tr, t, s.maximum) {
		ii = append(ii, NewIntersection(t, s))
	}

	return ii
}

func (s *Cone) checkCap(tr *Ray, t, radius float64) bool {
	x := tr.origin.x + t*tr.direction.x
	z := tr.origin.z + t*tr.direction.z

	return x*x+z*z <= radius*radius
}
