package render

var _ Pattern = (*BasePattern)(nil)

type Pattern interface {
	AtObject(obj Shape, point Tuple) Tuple
}

type BasePattern struct {
	transform *Matrix
	ConcretePattern
}

type ConcretePattern interface {
	At(p Tuple) Tuple
}

func DefaultBasePattern() *BasePattern {
	return &BasePattern{
		transform:       IdentityMatrix(),
		ConcretePattern: nil,
	}
}

func (p *BasePattern) SetTransform(t *Matrix) {
	p.transform = t
}

func (p *BasePattern) AtObject(obj Shape, point Tuple) Tuple {
	objTInv, err := obj.transform().Inverse() // TODO: Use cached versions
	if err != nil {
		panic(err)
	}
	patTInv, err2 := p.transform.Inverse()
	if err2 != nil {
		panic(err2)
	}

	wPoint := objTInv.MultiplyTuple(point)
	pPoint := patTInv.MultiplyTuple(wPoint)

	return p.At(pPoint)
}

type TestPattern struct {
	BasePattern
}

func NewTestPattern() *TestPattern {
	basePattern := *DefaultBasePattern()
	res := &TestPattern{
		BasePattern: basePattern,
	}
	res.BasePattern.ConcretePattern = res

	return res
}

func (p *TestPattern) At(point Tuple) Tuple {
	return NewColor(point.x, point.y, point.z)
}
