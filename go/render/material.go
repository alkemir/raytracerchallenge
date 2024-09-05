package render

import (
	"math"
)

type Material struct {
	color           Tuple
	ambient         float64
	diffuse         float64
	specular        float64
	shininess       float64
	reflective      float64
	transparency    float64
	refractiveIndex float64
	pattern         Pattern
}

func NewMaterial(color Tuple, ambient, diffuse, specular, reflective, transparency, refractiveIndex, shininess float64, pattern Pattern) *Material {
	return &Material{
		color:           color,
		ambient:         ambient,
		diffuse:         diffuse,
		specular:        specular,
		shininess:       shininess,
		reflective:      reflective,
		transparency:    transparency,
		refractiveIndex: refractiveIndex,
		pattern:         pattern,
	}
}

func DefaultMaterial() *Material {
	return &Material{
		color:           NewColor(1, 1, 1),
		ambient:         0.1,
		diffuse:         0.9,
		specular:        0.9,
		reflective:      0,
		transparency:    0,
		refractiveIndex: 1,
		shininess:       200,
		pattern:         nil, // TODO: Implement solid pattern
	}
}

func (m *Material) Equals(o *Material) bool {
	return m.color.Equals(o.color) &&
		math.Abs(m.ambient-o.ambient) < EPSILON &&
		math.Abs(m.diffuse-o.diffuse) < EPSILON &&
		math.Abs(m.specular-o.specular) < EPSILON &&
		math.Abs(m.shininess-o.shininess) < EPSILON &&
		math.Abs(m.reflective-o.reflective) < EPSILON &&
		m.pattern == o.pattern
}

func (m *Material) Lightning(obj Shape, l *Light, p, eye, normal Tuple, shadowed bool) Tuple {
	pColor := m.color
	if m.pattern != nil {
		pColor = m.pattern.AtObject(obj, p)
	}

	rCol := pColor.Hadamard(l.intensity)
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
