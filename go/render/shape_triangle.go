package render

import (
	"math"
)

var _ Shape = (*Triangle)(nil)
var _ ConcreteShape = (*Triangle)(nil)

type Triangle struct {
	p1     Tuple
	p2     Tuple
	p3     Tuple
	e1     Tuple
	e2     Tuple
	normal Tuple
	BaseShape
}

func NewTriangle(p1, p2, p3 Tuple) *Triangle {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)
	n := e2.Cross(e1).Norm()

	baseShape := *DefaultBaseShape()
	res := &Triangle{
		p1:        p1,
		p2:        p2,
		p3:        p3,
		e1:        e1,
		e2:        e2,
		normal:    n,
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Triangle) concreteNormal(p Tuple) Tuple {
	return s.normal
}

func (s *Triangle) concreteIntersect(tr *Ray) []*Intersection {
	TriangleToRay := tr.origin.Sub(NewPoint(0, 0, 0))

	a := tr.direction.Dot(tr.direction)
	b := 2 * tr.direction.Dot(TriangleToRay)
	c := TriangleToRay.Dot(TriangleToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	return []*Intersection{
		NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s),
		NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s),
	}
}
