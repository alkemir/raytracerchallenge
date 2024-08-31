package render

import (
	"math"
)

type Material struct {
	color     Tuple
	ambient   float64
	diffuse   float64
	specular  float64
	shininess float64
}

func NewMaterial(color Tuple, ambient, diffuse, specular, shininess float64) *Material {
	return &Material{
		color:     color,
		ambient:   ambient,
		diffuse:   diffuse,
		specular:  specular,
		shininess: shininess,
	}
}

func DefaultMaterial() *Material {
	return &Material{
		color:     NewColor(1, 1, 1),
		ambient:   0.1,
		diffuse:   0.9,
		specular:  0.9,
		shininess: 200,
	}
}

func (m *Material) Equals(o *Material) bool {
	return m.color.Equals(o.color) &&
		abs(m.ambient-o.ambient) < EPSILON &&
		abs(m.diffuse-o.diffuse) < EPSILON &&
		abs(m.specular-o.specular) < EPSILON &&
		abs(m.shininess-o.shininess) < EPSILON
}

func (m *Material) Lightning(l *Light, p, eye, normal Tuple, shadowed bool) Tuple {
	rCol := m.color.Hadamard(l.intensity)
	lVec := l.position.Sub(p).Norm()

	ambientContrib := rCol.Mul(m.ambient)
	if shadowed {
		return ambientContrib
	}

	diffuseContrib := NewColor(0, 0, 0)
	specularContrib := NewColor(0, 0, 0)
	if lightDotNormal := lVec.Dot(normal); lightDotNormal > 0 {
		diffuseContrib = rCol.Mul(m.diffuse * lightDotNormal)

		if reflectDotEye := lVec.Mul(-1).Reflect(normal).Dot(eye); reflectDotEye > 0 {
			specularContrib = l.intensity.Mul(m.specular * math.Pow(reflectDotEye, m.shininess))
		}
	}

	return ambientContrib.Add(diffuseContrib).Add(specularContrib)
}
