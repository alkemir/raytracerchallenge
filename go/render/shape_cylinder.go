package render

import (
	"math"
)

var _ Shape = (*Cylinder)(nil)
var _ ConcreteShape = (*Cylinder)(nil)

type Cylinder struct {
	minimum float64
	maximum float64
	BaseShape
}

func NewCylinder() *Cylinder {
	baseShape := *DefaultBaseShape()
	res := &Cylinder{
		minimum:   math.Inf(-1),
		maximum:   math.Inf(1),
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Cylinder) SetMinimum(m float64) {
	s.minimum = m
}

func (s *Cylinder) SetMaximum(m float64) {
	s.maximum = m
}

func (s *Cylinder) concreteNormal(p Tuple) Tuple {
	return NewVector(p.x, 0, p.z)
}

func (s *Cylinder) concreteIntersect(tr *Ray) []*Intersection {
	a := tr.direction.x*tr.direction.x + tr.direction.z*tr.direction.z
	if math.Abs(a) < EPSILON {
		return nil
	}

	b := 2 * (tr.origin.x*tr.direction.x + tr.origin.z*tr.direction.z)
	c := tr.origin.x*tr.origin.x + tr.origin.z*tr.origin.z - 1

	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return nil
	}

	t0 := (-b - math.Sqrt(discriminant)) / (2 * a)
	t1 := (-b + math.Sqrt(discriminant)) / (2 * a)

	if t1 < t0 {
		t0, t1 = t1, t0
	}

	ii := make([]*Intersection, 0)

	y0 := tr.origin.y + t0*tr.direction.y
	if s.minimum < y0 && y0 < s.maximum {
		ii = append(ii, NewIntersection(t0, s))
	}

	y1 := tr.origin.y + t1*tr.direction.y
	if s.minimum < y1 && y1 < s.maximum {
		ii = append(ii, NewIntersection(t1, s))
	}

	return ii
}
