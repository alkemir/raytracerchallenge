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

func (s *Sphere) Intersect(r *ray.Ray) []float64 {
	sphereToRay := r.Origin().Sub(tuple.NewPoint(0, 0, 0))

	a := r.Direction().Dot(r.Direction())
	b := 2 * r.Direction().Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	res := make([]float64, 2)

	res[0] = (-1*b - math.Sqrt(discriminant)) / (2 * a)
	res[1] = (-1*b + math.Sqrt(discriminant)) / (2 * a)
	return res
}
