package render

import "math"

type StripePattern struct {
	a Tuple
	b Tuple
}

func NewStripePattern(a, b Tuple) *StripePattern {
	return &StripePattern{
		a: a,
		b: b,
	}
}

func (p *StripePattern) At(point Tuple) Tuple {
	if int(math.Floor(point.x))%2 == 0 {
		return p.a
	}
	return p.b
}
