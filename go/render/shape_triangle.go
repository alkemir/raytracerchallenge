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

func (s *Triangle) Includes(o Shape) bool {
	return s == o
}

func (s *Triangle) concreteNormal(p Tuple, i *Intersection) Tuple {
	return s.normal
}

func (s *Triangle) concreteIntersect(tr *Ray) []*Intersection {
	directionCrossEdge := tr.direction.Cross(s.e2)
	det := s.e1.Dot(directionCrossEdge)
	if math.Abs(det) < EPSILON {
		return nil
	}

	f := 1.0 / det
	p1ToOrigin := tr.origin.Sub(s.p1)
	u := f * p1ToOrigin.Dot(directionCrossEdge)
	if u < 0 || u > 1 {
		return nil
	}

	originCrossEdge := p1ToOrigin.Cross(s.e1)
	v := f * tr.direction.Dot(originCrossEdge)
	if v < 0 || u+v > 1 {
		return nil
	}

	t := f * s.e2.Dot(originCrossEdge)
	return []*Intersection{NewIntersection(t, s)}
}
