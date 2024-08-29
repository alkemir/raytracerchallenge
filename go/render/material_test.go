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

	res := m.Lightning(light, p, eye, normal)

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

	res := m.Lightning(light, p, eye, normal)

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

	res := m.Lightning(light, p, eye, normal)

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

	res := m.Lightning(light, p, eye, normal)

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

	res := m.Lightning(light, p, eye, normal)

	if !res.Equals(NewColor(0.1, 0.1, 0.1)) {
		t.Fatal("Lightning is wrong")
	}
}
