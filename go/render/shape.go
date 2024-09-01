package render

type BaseShape struct {
	transform *Matrix
	material  *Material
	ConcreteShape
}

type ConcreteShape interface {
	concreteIntersect(r *Ray) []*Intersection
	concreteNormal(p Tuple) Tuple
}

func DefaultBaseShape() *BaseShape {
	return &BaseShape{
		transform:     IdentityMatrix(),
		material:      DefaultMaterial(),
		ConcreteShape: nil,
	}
}

func (s *BaseShape) SetTransform(m *Matrix) {
	s.transform = m
}

func (s *BaseShape) SetMaterial(m *Material) {
	s.material = m
}

func (s *BaseShape) Intersect(r *Ray) []*Intersection {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	tr := r.Transform(mInv)
	return s.concreteIntersect(tr)
}

func (s *BaseShape) Normal(p Tuple) Tuple {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	objP := mInv.MultiplyTuple(p)
	objN := s.concreteNormal(objP)
	worldN := mInv.Transpose().MultiplyTuple(objN).ZeroW()
	return worldN.Norm()
}

type TestShape struct {
	savedRay *Ray
	BaseShape
}

func NewTestShape() *TestShape {
	baseShape := *DefaultBaseShape()
	res := &TestShape{
		BaseShape: baseShape,
	}
	res.BaseShape.ConcreteShape = res

	return res
}

func (s *TestShape) concreteIntersect(r *Ray) []*Intersection {
	s.savedRay = r
	return nil
}

func (s *TestShape) concreteNormal(p Tuple) Tuple {
	return NewVector(p.x, p.y, p.z)
}
