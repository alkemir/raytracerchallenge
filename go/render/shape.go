package render

var _ Shape = (*TestShape)(nil)

type Shape interface {
	Intersect(r *Ray) []*Intersection
	Normal(p Tuple, i *Intersection) Tuple
	SetTransform(t *Matrix)
	Includes(s Shape) bool
	material() *Material
	transform() *Matrix
	setParent(p Shape)
	worldToObject(p Tuple) Tuple
	normalToWorld(n Tuple) Tuple
}

type BaseShape struct {
	_transform    *Matrix
	_transformInv *Matrix
	_material     *Material
	parent        Shape
	ConcreteShape
}

type ConcreteShape interface {
	concreteIntersect(r *Ray) []*Intersection
	concreteNormal(p Tuple, i *Intersection) Tuple
}

func DefaultBaseShape() *BaseShape {
	return &BaseShape{
		_transform:    IdentityMatrix(),
		_transformInv: IdentityMatrix(),
		_material:     DefaultMaterial(),
		ConcreteShape: nil,
	}
}

func (s *BaseShape) SetTransform(m *Matrix) {
	s._transform = m
	mInv, err := m.Inverse()
	if err != nil {
		panic(err)
	}
	s._transformInv = mInv
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

	return s._transformInv.MultiplyTuple(p)
}

func (s *BaseShape) normalToWorld(n Tuple) Tuple {
	n = s._transformInv.Transpose().MultiplyTuple(n)
	n = n.ZeroW()
	n = n.Norm()

	if s.parent != nil {
		n = s.parent.normalToWorld(n)
	}

	return n
}

func (s *BaseShape) Intersect(r *Ray) []*Intersection {
	tr := r.Transform(s._transformInv)
	return s.concreteIntersect(tr)
}

func (s *BaseShape) Normal(p Tuple, i *Intersection) Tuple {
	objP := s.worldToObject(p)
	objN := s.concreteNormal(objP, i)
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

func (s *TestShape) Includes(o Shape) bool {
	return s == o
}

func (s *TestShape) concreteIntersect(r *Ray) []*Intersection {
	s.savedRay = r
	return nil
}

func (s *TestShape) concreteNormal(p Tuple, i *Intersection) Tuple {
	return NewVector(p.x, p.y, p.z)
}
