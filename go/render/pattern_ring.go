package render

import "math"

var _ Pattern = (*RingPattern)(nil)
var _ ConcretePattern = (*RingPattern)(nil)

type RingPattern struct {
	a Tuple
	b Tuple
	BasePattern
}

func NewRingPattern(a, b Tuple) *RingPattern {
	basePattern := *DefaultBasePattern()
	res := &RingPattern{
		a:           a,
		b:           b,
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *RingPattern) At(point Tuple) Tuple {
	distance := int(math.Floor(math.Sqrt(point.x*point.x + point.z*point.z)))
	if distance%2 == 0 {
		return p.a
	}
	return p.b
}
