package shape

type Intersection struct {
	t   float64
	obj any
}

func NewIntersection(t float64, obj any) *Intersection {
	return &Intersection{t, obj}
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
