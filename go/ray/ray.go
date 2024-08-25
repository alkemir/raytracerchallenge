package ray

import "raytracerchallenge/tuple"

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
