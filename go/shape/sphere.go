package shape

import (
	"math"
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
)

type Sphere struct {
}

func NewSphere() *Sphere {
	return &Sphere{}
}

func (s *Sphere) Intersect(r *ray.Ray) []*Intersection {
	sphereToRay := r.Origin().Sub(tuple.NewPoint(0, 0, 0))

	a := r.Direction().Dot(r.Direction())
	b := 2 * r.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	res := make([]*Intersection, 2)

	res[0] = NewIntersection((-1*b-math.Sqrt(discriminant))/(2*a), s)
	res[1] = NewIntersection((-1*b+math.Sqrt(discriminant))/(2*a), s)
	return res
}
