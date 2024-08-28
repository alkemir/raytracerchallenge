package shape

import (
	"math"
	"raytracerchallenge/matrix"
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
)

type Sphere struct {
	transform *matrix.Matrix
	material  *Material
}

func NewSphere() *Sphere {
	return &Sphere{matrix.Identity, DefaultMaterial}
}

func (s *Sphere) SetTransform(m *matrix.Matrix) {
	s.transform = m
}

func (s Sphere) Material() *Material {
	return s.material
}

func (s *Sphere) SetMaterial(m *Material) {
	s.material = m
}

func (s *Sphere) Normal(p tuple.Tuple) tuple.Tuple {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	objP := mInv.MultiplyTuple(p)
	objN := objP.Sub(tuple.NewPoint(0, 0, 0))
	worldN := mInv.Transpose().MultiplyTuple(objN).ZeroW()
	return worldN.Norm()
}

func (s *Sphere) Intersect(r *ray.Ray) []*Intersection {
	mInv, err := s.transform.Inverse()
	if err != nil {
		panic(err)
	}

	tr := r.Transform(mInv)
	sphereToRay := tr.Origin().Sub(tuple.NewPoint(0, 0, 0))

	a := tr.Direction().Dot(tr.Direction())
	b := 2 * tr.Direction().Dot(sphereToRay)
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
