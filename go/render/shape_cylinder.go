package render

import (
	"math"
)

var _ Shape = (*Cylinder)(nil)
var _ ConcreteShape = (*Cylinder)(nil)

type Cylinder struct {
	BaseShape
}

func NewCylinder() *Cylinder {
	baseShape := *DefaultBaseShape()
	res := &Cylinder{
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
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

	return []*Intersection{
		NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s),
		NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s),
	}
}
