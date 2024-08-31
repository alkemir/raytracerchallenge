package render

type Intersection struct {
	t   float64
	obj any
}

func NewIntersection(t float64, obj any) *Intersection {
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
	object    any
	point     Tuple
	overPoint Tuple
	eye       Tuple
	normal    Tuple
	inside    bool
}

func (i *Intersection) Precompute(ray *Ray) *Comps {
	point := ray.Project(i.t)
	eye := ray.direction.Mul(-1)
	normal := i.obj.(*Sphere).Normal(point)
	inside := false

	if normal.Dot(eye) < 0 {
		inside = true
		normal = normal.Mul(-1)
	}
	overPoint := point.Add(normal.Mul(EPSILON))

	return &Comps{
		t:         i.t,
		object:    i.obj,
		point:     point,
		overPoint: overPoint,
		eye:       eye,
		normal:    normal,
		inside:    inside,
	}
}
