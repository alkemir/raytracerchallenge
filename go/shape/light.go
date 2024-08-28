package shape

import "raytracerchallenge/tuple"

type Light struct {
	position  tuple.Tuple
	intensity tuple.Tuple
}

func NewPointLight(p tuple.Tuple, i tuple.Tuple) *Light {
	return &Light{p, i}
}
