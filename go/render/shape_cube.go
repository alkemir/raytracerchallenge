package render

import (
	"math"
)

var _ Shape = (*Cube)(nil)
var _ ConcreteShape = (*Cube)(nil)

type Cube struct {
	BaseShape
}

func NewCube() *Cube {
	baseShape := *DefaultBaseShape()
	res := &Cube{
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Cube) concreteNormal(p Tuple) Tuple {
	return p.Sub(NewPoint(0, 0, 0))
}

func (s *Cube) concreteIntersect(tr *Ray) []*Intersection {
	xtMin, xtMax := checkAxis(tr.origin.x, tr.direction.x)
	ytMin, ytMax := checkAxis(tr.origin.y, tr.direction.y)
	ztMin, ztMax := checkAxis(tr.origin.z, tr.direction.z)

	tMin := math.Max(math.Max(xtMin, ytMin), ztMin)
	tMax := math.Min(math.Min(xtMax, ytMax), ztMax)

	res := make([]*Intersection, 2)

	res[0] = NewIntersection(tMin, s)
	res[1] = NewIntersection(tMax, s)
	return res
}

func checkAxis(origin, direction float64) (float64, float64) {
	tMinNumerator := -1 - origin
	tMaxNumerator := 1 - origin

	var tMin, tMax float64
	if math.Abs(direction) >= EPSILON {
		tMin = tMinNumerator / direction
		tMax = tMaxNumerator / direction
	} else {
		tMin = tMinNumerator * math.Inf(1)
		tMax = tMaxNumerator * math.Inf(1)
	}

	if tMin > tMax {
		return tMax, tMin
	}
	return tMin, tMax

}
