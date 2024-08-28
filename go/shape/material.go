package shape

import (
	"math"
	"raytracerchallenge/tuple"
)

type Material struct {
	color     tuple.Tuple
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

var DefaultMaterial *Material = &Material{
	color:     tuple.NewColor(1, 1, 1),
	ambient:   0.1,
	diffuse:   0.9,
	specular:  0.9,
	shininess: 200,
}

func NewMaterial(color tuple.Tuple, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{color, ambient, diffuse, specular, shininess}
}

func (m *Material) Lightning(l *Light, p, eye, normal tuple.Tuple) tuple.Tuple {
	rCol := m.color.Hadamard(l.intensity)
	lVec := l.position.Sub(p).Norm()

	ambientContrib := rCol.Mul(m.ambient)
	diffuseContrib := tuple.NewColor(0, 0, 0)
	specularContrib := tuple.NewColor(0, 0, 0)

	if lightDotNormal := lVec.Dot(normal); lightDotNormal > 0 {
		diffuseContrib = rCol.Mul(m.diffuse * lightDotNormal)

		if reflectDotEye := lVec.Mul(-1).Reflect(normal).Dot(eye); reflectDotEye > 0 {
			specularContrib = l.intensity.Mul(m.specular * math.Pow(reflectDotEye, m.shininess))
		}
	}

	return ambientContrib.Add(diffuseContrib).Add(specularContrib)
}
