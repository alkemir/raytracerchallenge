package render

type Intersection struct {
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

type Comps struct {
	t         float64
	object    Shape
	point     Tuple
	overPoint Tuple
	eye       Tuple
	normal    Tuple
	reflectv  Tuple
	inside    bool
}

func (i *Intersection) Precompute(ray *Ray, xs []*Intersection) *Comps {
	point := ray.Project(i.t)
	eye := ray.direction.Mul(-1)
	normal := i.obj.Normal(point)
	inside := false

	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Mul(-1)
	}
	overPoint := point.Add(normal.Mul(EPSILON))
	reflectv := ray.direction.Reflect(normal)

	return &Comps{
		t:         i.t,
		object:    i.obj,
		point:     point,
		overPoint: overPoint,
		eye:       eye,
		normal:    normal,
		reflectv:  reflectv,
		inside:    inside,
	}
}
