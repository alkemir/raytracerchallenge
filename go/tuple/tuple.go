package tuple

const EPSILON = 0.0000001

type Tuple struct {
	x, y, z, w float64
}

func NewTuple(x, y, z, w float64) Tuple {
	return Tuple{x, y, z, w}
}

func NewPoint(x, y, z float64) Tuple {
	return Tuple{x, y, z, 1}
}

func NewVector(x, y, z float64) Tuple {
	return Tuple{x, y, z, 0}
}

func (a Tuple) Equals(b Tuple) bool {
	return abs(a.x-b.x) < EPSILON &&
		abs(a.y-b.y) < EPSILON &&
		abs(a.z-b.z) < EPSILON &&
		abs(a.w-b.w) < EPSILON
}

func (a Tuple) IsPoint() bool {
	return abs(a.w-1) < EPSILON
}

func (a Tuple) IsVector() bool {
	return abs(a.w) < EPSILON
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
