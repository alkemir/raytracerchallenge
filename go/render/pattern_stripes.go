package render

import "math"

var _ Pattern = (*StripesPattern)(nil)
var _ ConcretePattern = (*StripesPattern)(nil)

type StripesPattern struct {
	a Tuple
	b Tuple
	BasePattern
}

func NewStripesPattern(a, b Tuple) *StripesPattern {
	basePattern := *DefaultBasePattern()
	res := &StripesPattern{
		a:           a,
		b:           b,
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *StripesPattern) At(point Tuple) Tuple {
	if int(math.Floor(point.x))%2 == 0 {
		return p.a
	}
	return p.b
}
