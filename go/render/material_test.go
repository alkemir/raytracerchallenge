package render

import (
	"math"
	"testing"
)

func TestMaterialDefault(t *testing.T) {
	m := DefaultMaterial()

	if !m.color.Equals(NewColor(1, 1, 1)) {
		t.Fatal("Color is wrong")
	}
	if m.ambient != 0.1 {
		t.Fatal("ambient is wrong")
	}
	if m.diffuse != 0.9 {
		t.Fatal("diffuse is wrong")
	}
	if m.specular != 0.9 {
		t.Fatal("specular is wrong")
	}
	if m.shininess != 200 {
		t.Fatal("shininess is wrong")
	}
}

func TestMaterialLightning_eyeBetweenLightSurface(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	shadowed := false

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(1.9, 1.9, 1.9)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_eyeBetweenLightSurface_EyeOffset(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, math.Sqrt2/2, math.Sqrt2/2)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	shadowed := false

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(1.0, 1.0, 1.0)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_eyeBetweenLightSurface_LightOffset(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	shadowed := false

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(0.73639610, 0.73639610, 0.73639610)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_eyeInReflectionPath(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, -math.Sqrt2/2, -math.Sqrt2/2)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 10, -10), NewColor(1, 1, 1))
	shadowed := false

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(1.63639610, 1.63639610, 1.63639610)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_eyeBehindSurface(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, 10), NewColor(1, 1, 1))
	shadowed := false

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(0.1, 0.1, 0.1)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_inShadow(t *testing.T) {
	m := DefaultMaterial()
	p := NewPoint(0, 0, 0)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	shadowed := true

	res := m.Lightning(NewSphere(), light, p, eye, normal, shadowed)

	if !res.Equals(NewColor(0.1, 0.1, 0.1)) {
		t.Fatal("Lightning is wrong")
	}
}

func TestMaterialLightning_pattern(t *testing.T) {
	pattern := NewStripesPattern(NewColor(1, 1, 1), NewColor(0, 0, 0))
	m := NewMaterial(NewColor(0, 0, 0), 1, 0, 0, 200, pattern)
	eye := NewVector(0, 0, -1)
	normal := NewVector(0, 0, -1)
	light := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	shadowed := false

	c1 := m.Lightning(NewSphere(), light, NewPoint(0.9, 0, 0), eye, normal, shadowed)
	c2 := m.Lightning(NewSphere(), light, NewPoint(1.1, 0, 0), eye, normal, shadowed)

	if !c1.Equals(NewColor(1, 1, 1)) {
		t.Fatal("Lightning is wrong")
	}
	if !c2.Equals(NewColor(0, 0, 0)) {
		t.Fatal("Lightning is wrong")
	}
}
