package render

import (
	"math"
	"slices"
)

type Intersection struct {
	u   float64
	v   float64
	t   float64
	obj Shape
}

func NewIntersection(t float64, obj Shape) *Intersection {
	return &Intersection{
		t:   t,
		obj: obj,
	}
}

func (i *Intersection) Object() any {
	return i.obj
}

func (i *Intersection) T() float64 {
	return i.t
}

func Hit(ii []*Intersection) *Intersection {
	var currentHit *Intersection = nil

	for _, x := range ii {
		if x.t > 0 && (currentHit == nil || x.t < currentHit.t) {
			currentHit = x
		}
	}

	return currentHit
}

func NewIntersectionUV(t float64, obj Shape, u, v float64) *Intersection {
	return &Intersection{
		u:   u,
		v:   v,
		t:   t,
		obj: obj,
	}
}

type Comps struct {
	t          float64
	object     Shape
	point      Tuple
	overPoint  Tuple
	underPoint Tuple
	eye        Tuple
	normal     Tuple
	reflectv   Tuple
	inside     bool
	n1         float64
	n2         float64
}

func (i *Intersection) Precompute(ray *Ray, xs []*Intersection) *Comps {
	point := ray.Project(i.t)
	eye := ray.direction.Mul(-1)
	normal := i.obj.Normal(point, i)
	inside := false

	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Mul(-1)
	}
	overPoint := point.Add(normal.Mul(EPSILON))
	underPoint := point.Sub(normal.Mul(EPSILON))
	reflectv := ray.direction.Reflect(normal)

	containers := make([]Shape, 0)

	var n1, n2 float64

	for _, x := range xs {
		if x == i {
			if len(containers) == 0 {
				n1 = 1
			} else {
				n1 = containers[len(containers)-1].material().refractiveIndex
			}
		}

		if contIdx := slices.Index(containers, x.obj); contIdx != -1 {
			containers = slices.Delete(containers, contIdx, contIdx+1)
		} else {
			containers = append(containers, x.obj)
		}

		if x == i {
			if len(containers) == 0 {
				n2 = 1.0
			} else {
				n2 = containers[len(containers)-1].material().refractiveIndex
			}
			break
		}
	}

	return &Comps{
		t:          i.t,
		object:     i.obj,
		point:      point,
		overPoint:  overPoint,
		underPoint: underPoint,
		eye:        eye,
		normal:     normal,
		reflectv:   reflectv,
		inside:     inside,
		n1:         n1,
		n2:         n2,
	}
}

func (c *Comps) Schlick() float64 {
	cos := c.eye.Dot(c.normal)

	if c.n1 > c.n2 {
		n := c.n1 / c.n2
		sin2T := n * n * (1 - cos*cos)
		if sin2T > 1 {
			return 1.0
		}

		cosT := math.Sqrt(1 - sin2T)
		cos = cosT
	}

	r0 := ((c.n1 - c.n2) / (c.n1 + c.n2))
	r0 = r0 * r0

	return r0 + (1-r0)*math.Pow(1-cos, 5)
}
