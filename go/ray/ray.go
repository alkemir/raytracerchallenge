package ray

import (
	"raytracerchallenge/matrix"
	"raytracerchallenge/tuple"
)

type Ray struct {
	origin    tuple.Tuple
	direction tuple.Tuple
}

func NewRay(origin, direction tuple.Tuple) *Ray {
	return &Ray{origin, direction}
}

func (r *Ray) Origin() tuple.Tuple {
	return r.origin
}

func (r *Ray) Direction() tuple.Tuple {
	return r.direction
}

func (r *Ray) Project(distance float64) tuple.Tuple {
	return r.direction.Mul(distance).Add(r.origin)
}

func (r *Ray) Transform(m *matrix.Matrix) *Ray {
	return &Ray{
		origin:    m.MultiplyTuple(r.origin),
		direction: m.MultiplyTuple(r.direction),
	}
}
