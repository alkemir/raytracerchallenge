package render

import "math"

type StripePattern struct {
	a         Tuple
	b         Tuple
	transform *Matrix
}

func NewStripePattern(a, b Tuple) *StripePattern {
	return &StripePattern{
		a:         a,
		b:         b,
		transform: IdentityMatrix(),
	}
}

func (p *StripePattern) SetTransform(t *Matrix) {
	p.transform = t
}

func (p *StripePattern) At(point Tuple) Tuple {
	if int(math.Floor(point.x))%2 == 0 {
		return p.a
	}
	return p.b
}

func (p *StripePattern) AtObject(obj Shape, point Tuple) Tuple {
	objTInv, err := obj.transform().Inverse()
	if err != nil {
		panic(err)
	}
	patTInv, err2 := p.transform.Inverse()
	if err2 != nil {
		panic(err)
	}

	wPoint := objTInv.MultiplyTuple(point)
	pPoint := patTInv.MultiplyTuple(wPoint)

	return p.At(pPoint)
}
