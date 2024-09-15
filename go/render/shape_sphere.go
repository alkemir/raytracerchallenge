package render

import (
	"math"
)

var _ Shape = (*Sphere)(nil)
var _ ConcreteShape = (*Sphere)(nil)

type Sphere struct {
	BaseShape
}

func NewSphere() *Sphere {
	baseShape := *DefaultBaseShape()
	res := &Sphere{
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func NewGlassSphere() *Sphere {
	res := NewSphere()
	res.material().transparency = 1
	res.material().refractiveIndex = 1.5
	return res
}

func (s *Sphere) concreteNormal(p Tuple, i *Intersection) Tuple {
	return p.Sub(NewPoint(0, 0, 0))
}

func (s *Sphere) concreteIntersect(tr *Ray) []*Intersection {
	sphereToRay := tr.origin.Sub(NewPoint(0, 0, 0))

	a := tr.direction.Dot(tr.direction)
	b := 2 * tr.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	return []*Intersection{
		NewIntersection((-b-math.Sqrt(discriminant))/(2*a), s),
		NewIntersection((-b+math.Sqrt(discriminant))/(2*a), s),
	}
}
