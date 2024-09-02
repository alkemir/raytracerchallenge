package render

import "math"

var _ Pattern = (*StripePattern)(nil)
var _ ConcretePattern = (*StripePattern)(nil)

type StripePattern struct {
	a Tuple
	b Tuple
	BasePattern
}

func NewStripePattern(a, b Tuple) *StripePattern {
	basePattern := *DefaultBasePattern()
	res := &StripePattern{
		a:           a,
		b:           b,
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *StripePattern) At(point Tuple) Tuple {
	if int(math.Floor(point.x))%2 == 0 {
		return p.a
	}
	return p.b
}
