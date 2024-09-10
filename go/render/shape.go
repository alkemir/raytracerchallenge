package render

var _ Shape = (*BaseShape)(nil)

type Shape interface {
	Intersect(r *Ray) []*Intersection
	Normal(p Tuple) Tuple
	material() *Material
	transform() *Matrix
	setParent(p Shape)
	worldToObject(p Tuple) Tuple
	normalToWorld(n Tuple) Tuple
}

type BaseShape struct {
	_transform *Matrix
	_material  *Material
	parent     Shape
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

func (s *BaseShape) setParent(p Shape) {
	s.parent = p
}

func (s *BaseShape) worldToObject(p Tuple) Tuple {
	if s.parent != nil {
		p = s.parent.worldToObject(p)
	}

	tInv, err := s.transform().Inverse()
	if err != nil {
		panic(err)
	}

	return tInv.MultiplyTuple(p)
}

func (s *BaseShape) normalToWorld(n Tuple) Tuple {
	tInv, err := s.transform().Inverse()
	if err != nil {
		panic(err)
	}

	n = tInv.Transpose().MultiplyTuple(n)
	n = n.ZeroW()
	n = n.Norm()

	if s.parent != nil {
		n = s.parent.normalToWorld(n)
	}

	return n
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
	objP := s.worldToObject(p)
	objN := s.concreteNormal(objP)
	worldN := s.normalToWorld(objN)
	return worldN
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
