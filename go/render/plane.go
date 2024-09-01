package render

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

func (s *Plane) concreteNormal(p Tuple) Tuple {
	return NewVector(0, 1, 0)
}

func (s *Plane) concreteIntersect(ray *Ray) []*Intersection {
	if abs(ray.direction.y) < EPSILON {
		return nil
	}

	t := -ray.origin.y / ray.direction.y
	return []*Intersection{NewIntersection(t, s)}
}
