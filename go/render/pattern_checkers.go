package render

import "math"

var _ Pattern = (*CheckersPattern)(nil)
var _ ConcretePattern = (*CheckersPattern)(nil)

type CheckersPattern struct {
	a Tuple
	b Tuple
	BasePattern
}

func NewCheckersPattern(a, b Tuple) *CheckersPattern {
	basePattern := *DefaultBasePattern()
	res := &CheckersPattern{
		a:           a,
		b:           b,
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *CheckersPattern) At(point Tuple) Tuple {
	totalCoord := int(math.Floor(point.x) + math.Floor(point.y) + math.Floor(point.z))
	if totalCoord%2 == 0 {
		return p.a
	}
	return p.b
}
