package shape

import (
	"raytracerchallenge/tuple"
	"testing"
)

func TestMaterialDefault(t *testing.T) {
	m := DefaultMaterial

	if !m.color.Equals(tuple.NewColor(1, 1, 1)) {
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
