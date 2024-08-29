package render

type Ray struct {
	origin    Tuple
	direction Tuple
}

func NewRay(origin, direction Tuple) *Ray {
	return &Ray{
		origin:    origin,
		direction: direction,
	}
}

func (r *Ray) Project(distance float64) Tuple {
	return r.direction.Mul(distance).Add(r.origin)
}

func (r *Ray) Transform(m *Matrix) *Ray {
	return &Ray{
		origin:    m.MultiplyTuple(r.origin),
		direction: m.MultiplyTuple(r.direction),
	}
}
