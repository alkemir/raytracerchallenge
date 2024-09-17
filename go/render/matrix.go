package render

import (
	"fmt"
	"math"
)

type Matrix struct {
	cols int
	rows int
	data []float64
}

func NewMatrix(rows, cols int, data []float64) *Matrix {
	return &Matrix{
		rows: rows,
		cols: cols,
		data: data,
	}
}

func IdentityMatrix() *Matrix {
	return NewMatrix(4, 4, []float64{1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1})
}

func (m *Matrix) Equals(o *Matrix) bool {
	if m.rows != o.rows {
		return false
	}

	if m.cols != o.cols {
		return false
	}

	for i := 0; i < m.cols*m.rows; i++ {
		if !(math.Abs(m.data[i]-o.data[i]) < EPSILON) {
			return false
		}
	}

	return true
}

func (m *Matrix) get(row, col int) float64 {
	return m.data[row*m.cols+col]
}

func (m *Matrix) Multiply(o *Matrix) *Matrix {
	rows := m.rows
	cols := o.cols
	resData := make([]float64, rows*cols)

	for y := 0; y < m.rows; y++ {
		for x := 0; x < o.cols; x++ {
			for i := 0; i < m.cols; i++ {
				resData[x+y*m.cols] += m.data[i+y*m.cols] * o.data[x+i*m.cols]
			}
		}
	}

	return NewMatrix(rows, cols, resData)
}

func (m *Matrix) MultiplyTuple(t Tuple) Tuple {
	x := m.data[0+0*m.cols]*t.X() + m.data[1+0*m.cols]*t.Y() +
		m.data[2+0*m.cols]*t.Z() + m.data[3+0*m.cols]*t.W()

	y := m.data[0+1*m.cols]*t.X() + m.data[1+1*m.cols]*t.Y() +
		m.data[2+1*m.cols]*t.Z() + m.data[3+1*m.cols]*t.W()

	z := m.data[0+2*m.cols]*t.X() + m.data[1+2*m.cols]*t.Y() +
		m.data[2+2*m.cols]*t.Z() + m.data[3+2*m.cols]*t.W()

	w := m.data[0+3*m.cols]*t.X() + m.data[1+3*m.cols]*t.Y() +
		m.data[2+3*m.cols]*t.Z() + m.data[3+3*m.cols]*t.W()

	return NewTuple(x, y, z, w)
}

func (m *Matrix) Transpose() *Matrix {
	cols := m.rows
	rows := m.cols

	resData := make([]float64, rows*cols)

	for y := 0; y < m.rows; y++ {
		for x := 0; x < m.cols; x++ {
			resData[y+x*m.cols] = m.data[x+y*m.cols]
		}
	}

	return NewMatrix(rows, cols, resData)
}

func (m *Matrix) SubMatrix(row, col int, scratchMatrix *Matrix) *Matrix {
	cols := m.cols - 1

	yOff := 0
	for y := 0; y < m.rows; y++ {
		if y == row {
			yOff = 1
			continue
		}

		xOff := 0
		for x := 0; x < m.cols; x++ {
			if x == col {
				xOff = 1
				continue
			}

			scratchMatrix.data[x-xOff+(y-yOff)*cols] = m.data[x+y*m.cols]
		}
	}

	return scratchMatrix
}

func (m *Matrix) Det() float64 {
	if m.cols == 2 && m.rows == 2 {
		return m.data[0]*m.data[3] - m.data[1]*m.data[2]
	}

	det := float64(0)
	scratch := NewMatrix(m.rows-1, m.cols-1, make([]float64, (m.rows-1)*(m.cols-1)))
	for y := 0; y < m.cols; y++ {
		det += m.data[y] * m.Cofactor(0, y, scratch)
	}
	return det
}

func (m *Matrix) Minor(row, col int, scratchMatrix *Matrix) float64 {
	return m.SubMatrix(row, col, scratchMatrix).Det()
}

func (m *Matrix) Cofactor(row, col int, scratchMatrix *Matrix) float64 {
	if (row+col)%2 == 1 {
		return -1 * m.Minor(row, col, scratchMatrix)
	}
	return m.Minor(row, col, scratchMatrix)
}

func (m *Matrix) Invertible() bool {
	return !(math.Abs(m.Det()) < EPSILON)
}

func (m *Matrix) Inverse() (*Matrix, error) {
	det := m.Det()

	if math.Abs(det) < EPSILON {
		return nil, fmt.Errorf("Matrix is not invertible")
	}

	resData := make([]float64, m.rows*m.cols)
	scratchMatrix := NewMatrix(3, 3, make([]float64, 9))
	for row := 0; row < m.rows; row++ {
		for col := 0; col < m.cols; col++ {
			resData[row+col*m.cols] = m.Cofactor(row, col, scratchMatrix) / det
		}
	}
	return NewMatrix(m.rows, m.cols, resData), nil
}

func (m *Matrix) Print() {
	for y := 0; y < m.rows; y++ {
		for x := 0; x < m.cols; x++ {
			fmt.Printf("%.10f ", m.data[x+y*m.cols])
		}
		fmt.Println()
	}
}
