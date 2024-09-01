package render

var _ Shape = (*BaseShape)(nil)

type Shape interface {
	Intersect(r *Ray) []*Intersection
	Normal(p Tuple) Tuple
	material() *Material
	transform() *Matrix
}

type BaseShape struct {
	_transform *Matrix
	_material  *Material
	ConcreteShape
}

type ConcreteShape interface {
	concreteIntersect(r *Ray) []*Intersection
	concreteNormal(p Tuple) Tuple
}

func DefaultBaseShape() *BaseShape {
	return &BaseShape{
		_transform:    IdentityMatrix(),
		_material:     DefaultMaterial(),
		ConcreteShape: nil,
	}
}

func (s *BaseShape) SetTransform(m *Matrix) {
	s._transform = m
}

func (s *BaseShape) transform() *Matrix {
	return s._transform
}

func (s *BaseShape) SetMaterial(m *Material) {
	s._material = m
}

func (s *BaseShape) material() *Material {
	return s._material
}

func (s *BaseShape) Intersect(r *Ray) []*Intersection {
	mInv, err := s._transform.Inverse()
	if err != nil {
		panic(err)
	}

	tr := r.Transform(mInv)
	return s.concreteIntersect(tr)
}

func (s *BaseShape) Normal(p Tuple) Tuple {
	mInv, err := s._transform.Inverse()
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
