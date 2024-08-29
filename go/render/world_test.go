package render

import (
	"testing"
)

func TestWorldConstructor(t *testing.T) {
	w := NewWorld()

	if len(w.lights) != 0 {
		t.Fatal("Initial world is not empty")
	}
	if len(w.objs) != 0 {
		t.Fatal("Initial world is not empty")
	}
}

func TestDefaultWorld(t *testing.T) {
	if len(DefaultWorld.lights) != 1 {
		t.Fatal("Default world has the wrong number of lights")
	}
	if !DefaultWorld.lights[0].intensity.Equals(NewColor(1, 1, 1)) {
		t.Fatal("Default world light is the wrong color")
	}
	if !DefaultWorld.lights[0].position.Equals(NewPoint(-10, 10, -10)) {
		t.Fatal("Default world light is at the wrong position")
	}

	if len(DefaultWorld.objs) != 2 {
		t.Fatal("Default world does not have default objects")
	}
	if !DefaultWorld.objs[0].material.color.Equals(NewColor(0.8, 1.0, 0.6)) {
		t.Fatal("Default object 1 has wrong color")
	}
	if DefaultWorld.objs[0].material.ambient != DefaultMaterial.ambient {
		t.Fatal("Default object 1 has wrong ambient")
	}
	if DefaultWorld.objs[0].material.diffuse != 0.7 {
		t.Fatal("Default object 1 has wrong diffuse")
	}
	if DefaultWorld.objs[0].material.specular != 0.2 {
		t.Fatal("Default object 1 has wrong specular")
	}
	if DefaultWorld.objs[0].material.shininess != DefaultMaterial.shininess {
		t.Fatal("Default object 1 has wrong shininess")
	}
	if !DefaultWorld.objs[0].transform.Equals(IdentityMatrix()) {
		t.Fatal("Default object 1 has wrong transform")
	}

	if !DefaultWorld.objs[1].material.color.Equals(DefaultMaterial.color) {
		t.Fatal("Default object 2 has wrong color")
	}
	if DefaultWorld.objs[1].material.ambient != DefaultMaterial.ambient {
		t.Fatal("Default object 2 has wrong ambient")
	}
	if DefaultWorld.objs[1].material.diffuse != DefaultMaterial.diffuse {
		t.Fatal("Default object 2 has wrong diffuse")
	}
	if DefaultWorld.objs[1].material.specular != DefaultMaterial.specular {
		t.Fatal("Default object 2 has wrong specular")
	}
	if DefaultWorld.objs[1].material.shininess != DefaultMaterial.shininess {
		t.Fatal("Default object 2 has wrong shininess")
	}
	if !DefaultWorld.objs[1].transform.Equals(Scaling(0.5, 0.5, 0.5)) {
		t.Fatal("Default object 2 has wrong transform")
	}
}

func TestDefaultWorldIntersect(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	xs := DefaultWorld.Intersect(r)

	if len(xs) != 4 {
		t.Fatal("Number of intersections is wrong")
	}
	if xs[0].t != 4 {
		t.Fatal("First intersection is wrong")
	}
	if xs[1].t != 4.5 {
		t.Fatal("Second intersection is wrong")
	}
	if xs[2].t != 5.5 {
		t.Fatal("Third intersection is wrong")
	}
	if xs[3].t != 6 {
		t.Fatal("Fourth intersection is wrong")
	}
}

func TestWorldShadeHit(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := DefaultWorld.objs[0]
	i := NewIntersection(4, shape)

	comps := i.Precompute(r)
	c := DefaultWorld.ShadeHit(comps)

	if !c.Equals(NewColor(0.38066119, 0.47582649, 0.2854958)) {
		t.Fatal("Color is wrong")
	}
}

func TestWorldShadeHit_inside(t *testing.T) {
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	l := NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))
	ol := DefaultWorld.lights[0]
	defer func() { DefaultWorld.lights[0] = ol }() // restore original light
	DefaultWorld.lights[0] = l                     // Grrr

	shape := DefaultWorld.objs[1]
	i := NewIntersection(0.5, shape)

	comps := i.Precompute(r)
	c := DefaultWorld.ShadeHit(comps)

	if !c.Equals(NewColor(0.90498447, 0.90498447, 0.90498447)) {
		t.Fatal("Color is wrong")
	}
}

func TestWorldShade_miss(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))

	c := DefaultWorld.Shade(r)

	if !c.Equals(NewColor(0, 0, 0)) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShade_hit(t *testing.T) {
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	c := DefaultWorld.Shade(r)

	if !c.Equals(NewColor(0.38066119, 0.47582649, 0.2854958)) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShade_hitInside(t *testing.T) {
	outer := DefaultWorld.objs[0]
	outer.material.ambient = 1
	inner := DefaultWorld.objs[1]
	inner.material.ambient = 1
	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))

	c := DefaultWorld.Shade(r)

	if !c.Equals(inner.material.color) {
		t.Fatal("Color was wrong")
	}
}
