package shape

import (
	"raytracerchallenge/matrix"
	"raytracerchallenge/ray"
	"raytracerchallenge/tuple"
	"slices"
)

type World struct {
	lights []*Light
	objs   []*Sphere
}

var DefaultWorld = NewWorld()

func init() {
	DefaultWorld.AddLight(NewPointLight(tuple.NewPoint(-10, 10, -10), tuple.NewColor(1, 1, 1)))

	s1 := NewSphere()
	s1.SetMaterial(NewMaterial(tuple.NewColor(0.8, 1.0, 0.6), DefaultMaterial.ambient, 0.7, 0.2, DefaultMaterial.shininess))
	DefaultWorld.AddObject(s1)

	s2 := NewSphere()
	s2.SetTransform(matrix.Scaling(0.5, 0.5, 0.5))
	DefaultWorld.AddObject(s2)
}

func NewWorld() *World {
	return &World{}
}

func (w *World) Objects() []*Sphere {
	return w.objs
}

func (w *World) Lights() []*Light {
	return w.lights
}

func (w *World) AddObject(obj *Sphere) {
	w.objs = append(w.objs, obj)
}

func (w *World) AddLight(light *Light) {
	w.lights = append(w.lights, light)
}

func (w *World) Intersect(ray *ray.Ray) []*Intersection {
	allIntersections := make([]*Intersection, 0)

	for _, obj := range w.objs {
		allIntersections = append(allIntersections, obj.Intersect(ray)...)
	}

	slices.SortFunc(allIntersections, func(a, b *Intersection) int {
		if a.t-b.t < 0 {
			return -1
		} else {
			return 1
		}
	})
	return allIntersections
}

func (w *World) ShadeHit(comps *Comps) tuple.Tuple {
	res := tuple.NewColor(0, 0, 0)
	for l := range w.lights {
		res = res.Add(comps.object.(*Sphere).material.Lightning(w.lights[l], comps.point, comps.eye, comps.normal))
	}
	return res
}
