package render

import (
	"testing"
)

func TestLightConstructor(t *testing.T) {
	i := NewColor(1, 1, 1)
	p := NewPoint(0, 0, 0)

	l := NewPointLight(p, i)

	if !l.position.Equals(p) {
		t.Fatal("Light position is wrong")
	}
	if !l.intensity.Equals(i) {
		t.Fatal("Light position is wrong")
	}
}
