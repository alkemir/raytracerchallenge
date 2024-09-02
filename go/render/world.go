package render

import (
	"slices"
)

type World struct {
	lights []*Light
	objs   []Shape
}

func NewWorld() *World {
	return &World{}
}

func DefaultWorld() *World {
	w := NewWorld()

	w.AddLight(NewPointLight(NewPoint(-10, 10, -10), NewColor(1, 1, 1)))

	s1 := NewSphere()
	s1.SetMaterial(NewMaterial(NewColor(0.8, 1.0, 0.6), 0.1, 0.7, 0.2, 200, nil))
	w.AddObject(s1)

	s2 := NewSphere()
	s2.SetTransform(Scaling(0.5, 0.5, 0.5))
	w.AddObject(s2)

	return w
}

func (w *World) AddObject(obj Shape) {
	w.objs = append(w.objs, obj)
}

func (w *World) AddLight(light *Light) {
	w.lights = append(w.lights, light)
}

func (w *World) Intersect(ray *Ray) []*Intersection {
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

func (w *World) ShadeHit(comps *Comps) Tuple {
	res := NewColor(0, 0, 0)
	for l := range w.lights {
		res = res.Add(comps.object.material().Lightning(w.lights[l], comps.point, comps.eye, comps.normal, w.IsShadowed(w.lights[l], comps.overPoint)))
	}
	return res
}

func (w *World) Shade(ray *Ray) Tuple {
	i := Hit(w.Intersect(ray)) // Can optimize since w.Intersect is sorted
	if i == nil {
		return NewColor(0, 0, 0)
	}

	comps := i.Precompute(ray)
	return w.ShadeHit(comps)
}

func (w *World) IsShadowed(l *Light, p Tuple) bool {
	v := l.position.Sub(p)
	distance := v.Mag()
	direction := v.Norm()

	r := NewRay(p, direction)
	ii := w.Intersect(r)

	h := Hit(ii)
	return h != nil && h.t < distance
}
