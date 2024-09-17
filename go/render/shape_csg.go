package render

import "slices"

var _ Shape = (*CSGShape)(nil)
var _ ConcreteShape = (*CSGShape)(nil)

type CSGShape struct {
	op    CSGOperation
	left  Shape
	right Shape
	BaseShape
}

func NewCSGShape(op CSGOperation, sl, sr Shape) *CSGShape {
	baseShape := *DefaultBaseShape()

	res := &CSGShape{
		op:        op,
		left:      sl,
		right:     sr,
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res
	sl.setParent(res)
	sr.setParent(res)

	return res
}

func (s *CSGShape) Includes(o Shape) bool {
	return s.left.Includes(o) || s.right.Includes(o)
}

func (s *CSGShape) filterIntersection(xs []*Intersection) []*Intersection {
	inL := false
	inR := false
	result := make([]*Intersection, 0, len(xs))

	for _, x := range xs {
		lHit := s.left.Includes(x.obj)
		if s.op.intersectionAllowed(lHit, inL, inR) {
			result = append(result, x)
		}

		if lHit {
			inL = !inL
		} else {
			inR = !inR
		}
	}

	return result
}

func (s *CSGShape) concreteIntersect(tr *Ray) []*Intersection {
	iiL := s.left.Intersect(tr)
	iiR := s.right.Intersect(tr)

	ii := make([]*Intersection, 0, len(iiL)+len(iiR))
	ii = append(ii, iiL...)
	ii = append(ii, iiR...)

	slices.SortFunc(ii, func(a *Intersection, b *Intersection) int {
		if a.t > b.t {
			return 1
		}
		return -1
	})

	return s.filterIntersection(ii)
}
