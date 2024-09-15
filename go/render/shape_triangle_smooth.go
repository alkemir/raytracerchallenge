package render

import "math"

var _ Shape = (*SmoothTriangle)(nil)
var _ ConcreteShape = (*SmoothTriangle)(nil)

type SmoothTriangle struct {
	p1 Tuple
	p2 Tuple
	p3 Tuple
	n1 Tuple
	n2 Tuple
	n3 Tuple
	e1 Tuple
	e2 Tuple
	BaseShape
}

func NewSmoothTriangle(p1, p2, p3, n1, n2, n3 Tuple) *SmoothTriangle {
	e1 := p2.Sub(p1)
	e2 := p3.Sub(p1)

	baseShape := *DefaultBaseShape()
	res := &SmoothTriangle{
		p1:        p1,
		p2:        p2,
		p3:        p3,
		n1:        n1,
		n2:        n2,
		n3:        n3,
		e1:        e1,
		e2:        e2,
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *SmoothTriangle) concreteNormal(p Tuple, i *Intersection) Tuple {
	return s.n2.Mul(i.u).Add(s.n3.Mul(i.v)).Add(s.n1.Mul(1 - i.u - i.v))
}

func (s *SmoothTriangle) concreteIntersect(tr *Ray) []*Intersection {
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
	return []*Intersection{NewIntersectionUV(t, s, u, v)}
}
