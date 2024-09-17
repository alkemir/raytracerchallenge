package render

var _ Pattern = (*BasePattern)(nil)

type Pattern interface {
	AtObject(obj Shape, point Tuple) Tuple
}

type BasePattern struct {
	transform    *Matrix
	transformInv *Matrix
	ConcretePattern
}

type ConcretePattern interface {
	At(p Tuple) Tuple
}

func DefaultBasePattern() *BasePattern {
	return &BasePattern{
		transform:       IdentityMatrix(),
		transformInv:    IdentityMatrix(),
		ConcretePattern: nil,
	}
}

func (p *BasePattern) SetTransform(t *Matrix) {
	p.transform = t
	patTInv, err := t.Inverse()
	if err != nil {
		panic(err)
	}
	p.transformInv = patTInv
}

func (p *BasePattern) AtObject(obj Shape, point Tuple) Tuple {
	wPoint := obj.transformInv().MultiplyTuple(point)
	pPoint := p.transformInv.MultiplyTuple(wPoint)

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
