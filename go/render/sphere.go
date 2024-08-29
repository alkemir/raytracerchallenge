package render

import (
	"math"
)

type Sphere struct {
	transform *Matrix
	material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{
		transform: IdentityMatrix(),
		material:  DefaultMaterial(),
	}
}

func (s *Sphere) SetTransform(m *Matrix) {
	s.transform = m
}

func (s *Sphere) SetMaterial(m *Material) {
	s.material = m
}

func (s *Sphere) Normal(p Tuple) Tuple {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	objP := mInv.MultiplyTuple(p)
	objN := objP.Sub(NewPoint(0, 0, 0))
	worldN := mInv.Transpose().MultiplyTuple(objN).ZeroW()
	return worldN.Norm()
}

func (s *Sphere) Intersect(r *Ray) []*Intersection {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	tr := r.Transform(mInv)
	sphereToRay := tr.origin.Sub(NewPoint(0, 0, 0))

	a := tr.direction.Dot(tr.direction)
	b := 2 * tr.direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1

	discriminant := b*b - 4*a*c

	if discriminant < 0 {
		return nil
	}

	res := make([]*Intersection, 2)

	res[0] = NewIntersection((-1*b-math.Sqrt(discriminant))/(2*a), s)
	res[1] = NewIntersection((-1*b+math.Sqrt(discriminant))/(2*a), s)
	return res
}
