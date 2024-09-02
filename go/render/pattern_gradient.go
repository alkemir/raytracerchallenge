package render

import "math"

var _ Pattern = (*GradientPattern)(nil)
var _ ConcretePattern = (*GradientPattern)(nil)

type GradientPattern struct {
	a Tuple
	b Tuple
	BasePattern
}

func NewGradientPattern(a, b Tuple) *GradientPattern {
	basePattern := *DefaultBasePattern()
	res := &GradientPattern{
		a:           a,
		b:           b,
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *GradientPattern) At(point Tuple) Tuple {
	delta := p.b.Sub(p.a)
	return p.a.Add(delta.Mul(point.x - math.Floor(point.x)))
}
