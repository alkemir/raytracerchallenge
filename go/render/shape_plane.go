package render

import "math"

var _ Shape = (*Plane)(nil)
var _ ConcreteShape = (*Plane)(nil)

type Plane struct {
	BaseShape
}

func NewPlane() *Plane {
	baseShape := *DefaultBaseShape()
	res := &Plane{
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *Plane) Includes(o Shape) bool {
	return s == o
}

func (s *Plane) concreteNormal(p Tuple, i *Intersection) Tuple {
	return NewVector(0, 1, 0)
}

func (s *Plane) concreteIntersect(ray *Ray) []*Intersection {
	if math.Abs(ray.direction.y) < EPSILON {
		return nil
	}

	t := -ray.origin.y / ray.direction.y
	return []*Intersection{NewIntersection(t, s)}
}
