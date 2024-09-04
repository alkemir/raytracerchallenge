package render

import (
	"math"
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
	w := DefaultWorld()
	if len(w.lights) != 1 {
		t.Fatal("Default world has the wrong number of lights")
	}
	if !w.lights[0].intensity.Equals(NewColor(1, 1, 1)) {
		t.Fatal("Default world light is the wrong color")
	}
	if !w.lights[0].position.Equals(NewPoint(-10, 10, -10)) {
		t.Fatal("Default world light is at the wrong position")
	}

	if len(w.objs) != 2 {
		t.Fatal("Default world does not have default objects")
	}
	if !w.objs[0].material().color.Equals(NewColor(0.8, 1.0, 0.6)) {
		t.Fatal("Default object 1 has wrong color")
	}
	if w.objs[0].material().ambient != 0.1 {
		t.Fatal("Default object 1 has wrong ambient")
	}
	if w.objs[0].material().diffuse != 0.7 {
		t.Fatal("Default object 1 has wrong diffuse")
	}
	if w.objs[0].material().specular != 0.2 {
		t.Fatal("Default object 1 has wrong specular")
	}
	if w.objs[0].material().shininess != 200 {
		t.Fatal("Default object 1 has wrong shininess")
	}
	if !w.objs[0].transform().Equals(IdentityMatrix()) {
		t.Fatal("Default object 1 has wrong transform")
	}

	if !w.objs[1].material().color.Equals(DefaultMaterial().color) {
		t.Fatal("Default object 2 has wrong color")
	}
	if w.objs[1].material().ambient != 0.1 {
		t.Fatal("Default object 2 has wrong ambient")
	}
	if w.objs[1].material().diffuse != 0.9 {
		t.Fatal("Default object 2 has wrong diffuse")
	}
	if w.objs[1].material().specular != 0.9 {
		t.Fatal("Default object 2 has wrong specular")
	}
	if w.objs[1].material().shininess != 200 {
		t.Fatal("Default object 2 has wrong shininess")
	}
	if !w.objs[1].transform().Equals(Scaling(0.5, 0.5, 0.5)) {
		t.Fatal("Default object 2 has wrong transform")
	}
}

func TestDefaultWorldIntersect(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	xs := w.Intersect(r)

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
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))
	shape := w.objs[0]
	i := NewIntersection(4, shape)

	comps := i.Precompute(r, []*Intersection{i})
	c := w.ShadeHit(comps, 1)

	if !c.Equals(NewColor(0.38066119, 0.47582649, 0.2854958)) {
		t.Fatal("Color is wrong")
	}
}

func TestWorldShadeHit_inside(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	l := NewPointLight(NewPoint(0, 0.25, 0), NewColor(1, 1, 1))
	w.lights[0] = l

	shape := w.objs[1]
	i := NewIntersection(0.5, shape)

	comps := i.Precompute(r, []*Intersection{i})
	c := w.ShadeHit(comps, 1)

	if !c.Equals(NewColor(0.90498447, 0.90498447, 0.90498447)) {
		t.Fatal("Color is wrong")
	}
}

func TestWorldShade_miss(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 1, 0))

	c := w.Shade(r, 1)

	if !c.Equals(NewColor(0, 0, 0)) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShade_shadow(t *testing.T) {
	w := NewWorld()
	l := NewPointLight(NewPoint(0, 0, -10), NewColor(1, 1, 1))
	w.AddLight(l)

	s1 := NewSphere()
	w.AddObject(s1)
	s2 := NewSphere()
	s2.SetTransform(Translation(0, 0, 10))
	w.AddObject(s2)

	r := NewRay(NewPoint(0, 0, 5), NewVector(0, 0, 1))
	i := NewIntersection(4, s2)

	comps := i.Precompute(r, []*Intersection{i})
	c := w.ShadeHit(comps, 1)

	if !c.Equals(NewColor(0.1, 0.1, 0.1)) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShade_hit(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, -5), NewVector(0, 0, 1))

	c := w.Shade(r, 1)

	if !c.Equals(NewColor(0.38066119, 0.47582649, 0.2854958)) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShade_hitInside(t *testing.T) {
	w := DefaultWorld()

	outer := w.objs[0]
	outer.material().ambient = 1
	inner := w.objs[1]
	inner.material().ambient = 1
	r := NewRay(NewPoint(0, 0, 0.75), NewVector(0, 0, -1))

	c := w.Shade(r, 1)

	if !c.Equals(inner.material().color) {
		t.Fatal("Color was wrong")
	}
}

func TestWorldShadow_no(t *testing.T) {
	w := DefaultWorld()
	p := NewPoint(0, 10, 0)

	if w.IsShadowed(w.lights[0], p) {
		t.Fatal("Shadow reported is wrong")
	}
}

func TestWorldShadow_noSide(t *testing.T) {
	w := DefaultWorld()
	p := NewPoint(0, 10, 0)

	if w.IsShadowed(w.lights[0], p) {
		t.Fatal("Shadow reported is wrong")
	}
}

func TestWorldShadow_yes(t *testing.T) {
	w := DefaultWorld()
	p := NewPoint(10, -10, 10)

	if !w.IsShadowed(w.lights[0], p) {
		t.Fatal("Shadow reported is wrong")
	}
}

func TestWorldShadow_noBehind(t *testing.T) {
	w := DefaultWorld()
	p := NewPoint(-20, 20, 20)

	if w.IsShadowed(w.lights[0], p) {
		t.Fatal("Shadow reported is wrong")
	}
}

func TestWorldShadow_noBetween(t *testing.T) {
	w := DefaultWorld()
	p := NewPoint(-2, 2, -2)

	if w.IsShadowed(w.lights[0], p) {
		t.Fatal("Shadow reported is wrong")
	}
}

func TestWorldReflect_nonReflective(t *testing.T) {
	w := DefaultWorld()
	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 0, 1))
	shape := w.objs[1]
	shape.material().ambient = 1
	i := NewIntersection(1, shape)

	comps := i.Precompute(r, []*Intersection{i})

	c := w.ReflectedColor(comps, 1)

	if !c.Equals(NewColor(0, 0, 0)) {
		t.Fatal("Reflection is wrong")
	}
}

func TestWorldReflect_reflective(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.material().reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.AddObject(shape)

	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := NewIntersection(math.Sqrt2, shape)

	comps := i.Precompute(r, []*Intersection{i})

	c := w.ReflectedColor(comps, 1)

	if !c.Equals(NewColor(0.1903322, 0.2379152, 0.1427491)) {
		t.Fatal("Reflection is wrong")
	}
}

func TestWorldShadeHit_reflective(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.material().reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.AddObject(shape)

	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := NewIntersection(math.Sqrt2, shape)

	comps := i.Precompute(r, []*Intersection{i})

	c := w.ShadeHit(comps, 1)

	if !c.Equals(NewColor(0.876757, 0.92434, 0.829174)) {
		t.Fatal("Reflection is wrong")
	}
}

func TestWorldShadeHit_infiniteReflection(t *testing.T) {
	w := NewWorld()
	l := NewPointLight(NewPoint(0, 0, 0), NewColor(1, 1, 1))
	w.AddLight(l)

	lower := NewPlane()
	lower.material().reflective = 1
	lower.SetTransform(Translation(0, -1, 0))

	upper := NewPlane()
	upper.material().reflective = 1
	upper.SetTransform(Translation(0, 1, 0))

	w.AddObject(lower)
	w.AddObject(upper)

	r := NewRay(NewPoint(0, 0, 0), NewVector(0, 1, 0))

	w.Shade(r, 1)
}

func TestWorldReflect_reflectiveAtMaxDepth(t *testing.T) {
	w := DefaultWorld()
	shape := NewPlane()
	shape.material().reflective = 0.5
	shape.SetTransform(Translation(0, -1, 0))
	w.AddObject(shape)

	r := NewRay(NewPoint(0, 0, -3), NewVector(0, -math.Sqrt2/2, math.Sqrt2/2))
	i := NewIntersection(math.Sqrt2, shape)

	comps := i.Precompute(r, []*Intersection{i})

	c := w.ReflectedColor(comps, 0)

	if !c.Equals(NewColor(0, 0, 0)) {
		t.Fatal("Reflection is wrong")
	}
}
