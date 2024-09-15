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

func (s *Cube) concreteNormal(p Tuple, i *Intersection) Tuple {
	maxC := math.Max(math.Max(math.Abs(p.x), math.Abs(p.y)), math.Abs(p.z))
	if math.Abs(p.x) == maxC {
		return NewVector(p.x, 0, 0)
	}
	if math.Abs(p.y) == maxC {
		return NewVector(0, p.y, 0)
	}
	return NewVector(0, 0, p.z)
}

func (s *Cube) concreteIntersect(tr *Ray) []*Intersection {
	xtMin, xtMax := checkAxis(tr.origin.x, tr.direction.x)
	ytMin, ytMax := checkAxis(tr.origin.y, tr.direction.y)
	ztMin, ztMax := checkAxis(tr.origin.z, tr.direction.z)

	tMin := math.Max(math.Max(xtMin, ytMin), ztMin)
	tMax := math.Min(math.Min(xtMax, ytMax), ztMax)

	if tMin > tMax {
		return nil
	}

	return []*Intersection{
		NewIntersection(tMin, s),
		NewIntersection(tMax, s),
	}
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
