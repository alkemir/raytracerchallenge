package tuple

import "math"

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

func NewColor(r, g, b float64) Tuple {
	return Tuple{r, g, b, 0}
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

func (a Tuple) Add(b Tuple) Tuple {
	return Tuple{a.x + b.x, a.y + b.y, a.z + b.z, a.w + b.w}
}

func (a Tuple) Sub(b Tuple) Tuple {
	return Tuple{a.x - b.x, a.y - b.y, a.z - b.z, a.w - b.w}
}

func (a Tuple) Neg() Tuple {
	return Tuple{-a.x, -a.y, -a.z, -a.w}
}

func (a Tuple) Mul(f float64) Tuple {
	return Tuple{a.x * f, a.y * f, a.z * f, a.w * f}
}

func (a Tuple) Div(f float64) Tuple {
	return Tuple{a.x / f, a.y / f, a.z / f, a.w / f}
}

func (a Tuple) Mag() float64 {
	// Taking the magnitude of a point makes no sense, not adding to avoid rounding errors
	return math.Sqrt(a.x*a.x + a.y*a.y + a.z*a.z)
}

func (a Tuple) Norm() Tuple {
	return a.Div(a.Mag())
}

func (a Tuple) Dot(b Tuple) float64 {
	return a.x*b.x + a.y*b.y + a.z*b.z
}

func (a Tuple) Cross(b Tuple) Tuple {
	return Tuple{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z,
		a.x*b.y - a.y*b.x, 0}
}

func (a Tuple) Hadamard(b Tuple) Tuple {
	return Tuple{a.x * b.x, a.y * b.y, a.z * b.z, 0}
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	} else {
		return a
	}
}
