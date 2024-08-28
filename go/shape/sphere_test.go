package shape

import (
	"math"
	"raytracerchallenge/matrix"
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
	"testing"
)

func TestSphereIntersect(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != 4 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 6 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_tangent(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 1, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != 5 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 5 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_miss(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 2, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestSphereIntersect_inside(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 0), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != -1 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 1 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_behind(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, 5), tuple.NewVector(0, 0, 1))
	s := NewSphere()

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != -6 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != -4 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereTransform_default(t *testing.T) {
	s := NewSphere()

	if !s.transform.Equals(matrix.Identity) {
		t.Fatal("Default sphere transform is not the identity")
	}
}

func TestSphereTransform_changed(t *testing.T) {
	s := NewSphere()
	m := matrix.Translation(2, 3, 4)

	s.SetTransform(m)

	if !s.transform.Equals(m) {
		t.Fatal("Set sphere transform is wrong")
	}
}

func TestSphereIntersect_scaled(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(matrix.Scaling(2, 2, 2))

	points := s.Intersect(r)

	if len(points) != 2 {
		t.Fatal("Wrong number of intersections")
	}
	if points[0].t != 3 {
		t.Fatal("Wrong intersection")
	}
	if points[1].t != 7 {
		t.Fatal("Wrong intersection")
	}
}

func TestSphereIntersect_translated(t *testing.T) {
	r := ray.NewRay(tuple.NewPoint(0, 0, -5), tuple.NewVector(0, 0, 1))
	s := NewSphere()
	s.SetTransform(matrix.Translation(5, 0, 0))

	points := s.Intersect(r)

	if len(points) != 0 {
		t.Fatal("Wrong number of intersections")
	}
}

func TestSphereNormal_xaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(tuple.NewPoint(1, 0, 0))

	if !n.Equals(tuple.NewVector(1, 0, 0)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSphereNormal_yaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(tuple.NewPoint(0, 1, 0))

	if !n.Equals(tuple.NewVector(0, 1, 0)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_zaxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(tuple.NewPoint(0, 0, 1))

	if !n.Equals(tuple.NewVector(0, 0, 1)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_offAxis(t *testing.T) {
	s := NewSphere()

	n := s.Normal(tuple.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(tuple.NewVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSphereNormal_normalized(t *testing.T) {
	s := NewSphere()

	n := s.Normal(tuple.NewPoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))

	if !n.Equals(n.Norm()) {
		t.Fatal("Normal is not normal")
	}
}

func TestSphereNormal_translated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(matrix.Translation(0, 1, 0))

	n := s.Normal(tuple.NewPoint(0, 1.70710678, -0.70710678))

	if !n.Equals(tuple.NewVector(0, 0.70710678, -0.70710678)) {
		t.Fatal("Normal is wrong")
	}
}
func TestSphereNormal_scaledRotated(t *testing.T) {
	s := NewSphere()
	s.SetTransform(matrix.Scaling(1, 0.5, 1).Multiply(matrix.RotationZ(math.Pi / 5)))

	n := s.Normal(tuple.NewPoint(0, math.Sqrt2/2, -math.Sqrt2/2))

	if !n.Equals(tuple.NewVector(0, 0.970142500, -0.24253562)) {
		t.Fatal("Normal is wrong")
	}
}

func TestSphereMaterial_default(t *testing.T) {
	s := NewSphere()

	if s.material != DefaultMaterial {
		t.Fatal("Material is wrong")
	}
}

func TestSphereMaterial_set(t *testing.T) {
	s := NewSphere()
	m := NewMaterial(DefaultMaterial.color, DefaultMaterial.ambient, DefaultMaterial.diffuse, DefaultMaterial.specular, DefaultMaterial.shininess)
	m.ambient = 1

	s.SetMaterial(m)

	if !s.material.color.Equals(DefaultMaterial.color) {
		t.Fatal("Material color is wrong")
	}
	if s.material.ambient != 1 {
		t.Fatal("Material ambient is wrong")
	}
	if s.material.diffuse != DefaultMaterial.diffuse {
		t.Fatal("Material difusse is wrong")
	}
	if s.material.specular != DefaultMaterial.specular {
		t.Fatal("Material specular is wrong")
	}
	if s.material.shininess != DefaultMaterial.shininess {
		t.Fatal("Material shininess is wrong")
	}
}
