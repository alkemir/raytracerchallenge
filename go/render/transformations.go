package render

import "math"

func Translation(x, y, z float64) *Matrix {
	return NewMatrix(4, 4, []float64{1, 0, 0, x, 0, 1, 0, y, 0, 0, 1, z, 0, 0, 0, 1})
}

func Scaling(x, y, z float64) *Matrix {
	return NewMatrix(4, 4, []float64{x, 0, 0, 0, 0, y, 0, 0, 0, 0, z, 0, 0, 0, 0, 1})
}

func RotationX(rad float64) *Matrix {
	return NewMatrix(4, 4, []float64{1, 0, 0, 0, 0, math.Cos(rad), -1 * math.Sin(rad), 0, 0, math.Sin(rad), math.Cos(rad), 0, 0, 0, 0, 1})
}

func RotationY(rad float64) *Matrix {
	return NewMatrix(4, 4, []float64{math.Cos(rad), 0, math.Sin(rad), 0, 0, 1, 0, 0, -1 * math.Sin(rad), 0, math.Cos(rad), 0, 0, 0, 0, 1})
}

func RotationZ(rad float64) *Matrix {
	return NewMatrix(4, 4, []float64{math.Cos(rad), -1 * math.Sin(rad), 0, 0, math.Sin(rad), math.Cos(rad), 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
}

func Shearing(xtoy, xtoz, ytox, ytoz, ztox, ztoy float64) *Matrix {
	return NewMatrix(4, 4, []float64{1, xtoy, xtoz, 0, ytox, 1, ytoz, 0, ztox, ztoy, 1, 0, 0, 0, 0, 1})
}

func View(from Tuple, to Tuple, up Tuple) *Matrix {
	forward := to.Sub(from).Norm()
	upn := up.Norm()
	left := forward.Cross(upn)
	trueUp := left.Cross(forward)

	return NewMatrix(4, 4, []float64{
		left.x, left.y, left.z, 0,
		trueUp.x, trueUp.y, trueUp.z, 0,
		-forward.x, -forward.y, -forward.z, 0,
		0, 0, 0, 1,
	}).Multiply(Translation(-from.x, -from.y, -from.z))
}

func (m *Matrix) Translate(x, y, z float64) *Matrix {
	return Translation(x, y, z).Multiply(m)
}

func (m *Matrix) Scale(x, y, z float64) *Matrix {
	return Scaling(x, y, z).Multiply(m)
}

func (m *Matrix) RotateX(rad float64) *Matrix {
	return RotationX(rad).Multiply(m)
}

func (m *Matrix) RotateY(rad float64) *Matrix {
	return RotationY(rad).Multiply(m)
}

func (m *Matrix) RotateZ(rad float64) *Matrix {
	return RotationZ(rad).Multiply(m)
}

func (m *Matrix) Shear(xtoy, xtoz, ytox, ytoz, ztox, ztoy float64) *Matrix {
	return Shearing(xtoy, xtoz, ytox, ytoz, ztox, ztoy).Multiply(m)
}
