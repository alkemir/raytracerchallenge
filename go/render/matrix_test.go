package render

import (
	"testing"
)

func TestMatrixSubscriptable4x4(t *testing.T) {
	m := NewMatrix(4, 4, []float64{1, 2, 3, 4, 5.5, 6.5, 7.5, 8.5, 9, 10, 11, 12, 13.5, 14.5, 15.5, 16.5})

	if m.get(0, 0) != 1 {
		t.Fatal("Element 0,0 was wrong")
	}
	if m.get(0, 3) != 4 {
		t.Fatal("Element 0,3 was wrong")
	}
	if m.get(1, 0) != 5.5 {
		t.Fatal("Element 1,0 was wrong")
	}
	if m.get(1, 2) != 7.5 {
		t.Fatal("Element 1,2 was wrong")
	}
	if m.get(2, 2) != 11 {
		t.Fatal("Element 2,2 was wrong")
	}
	if m.get(3, 0) != 13.5 {
		t.Fatal("Element 3,0 was wrong")
	}
	if m.get(3, 2) != 15.5 {
		t.Fatal("Element 3,2 was wrong")
	}
}

func TestMatrixSubscriptable2x2(t *testing.T) {
	m := NewMatrix(2, 2, []float64{-3, 5, 1, -2})

	if m.get(0, 0) != -3 {
		t.Fatal("Element 0,0 was wrong")
	}
	if m.get(0, 1) != 5 {
		t.Fatal("Element 0,1 was wrong")
	}
	if m.get(1, 0) != 1 {
		t.Fatal("Element 1,0 was wrong")
	}
	if m.get(1, 1) != -2 {
		t.Fatal("Element 1,1 was wrong")
	}
}

func TestMatrixSubscriptable3x3(t *testing.T) {
	m := NewMatrix(3, 3, []float64{-3, 5, 0, 1, -2, -7, 0, 1, 1})

	if m.get(0, 0) != -3 {
		t.Fatal("Element 0,0 was wrong")
	}
	if m.get(1, 1) != -2 {
		t.Fatal("Element 1,1 was wrong")
	}
	if m.get(2, 2) != 1 {
		t.Fatal("Element 2,2 was wrong")
	}
}

func TestCompareMatricesSame(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := NewMatrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})

	if !m1.Equals(m2) {
		t.Fatal("Matrix should equal a matrix with the same values")
	}
}

func TestCompareMatricesDifferent(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := NewMatrix(4, 4, []float64{2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2, 1})

	if m1.Equals(m2) {
		t.Fatal("Matrix should not equal a different matrix")
	}
}

func TestMultiplyMatrices(t *testing.T) {
	m1 := NewMatrix(4, 4, []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 8, 7, 6, 5, 4, 3, 2})
	m2 := NewMatrix(4, 4, []float64{-2, 1, 2, 3, 3, 2, 1, -1, 4, 3, 6, 5, 1, 2, 7, 8})

	p := NewMatrix(4, 4, []float64{20, 22, 50, 48, 44, 54, 114, 108, 40, 58, 110, 102, 16, 26, 46, 42})

	if !p.Equals(m1.Multiply(m2)) {
		t.Fatal("Matrix should equal its product")
	}
}

func TestMultiplyMatrixByTuple(t *testing.T) {
	m := NewMatrix(4, 4, []float64{1, 2, 3, 4, 2, 4, 4, 2, 8, 6, 4, 1, 0, 0, 0, 1})
	tu := NewTuple(1, 2, 3, 1)

	p := NewTuple(18, 24, 33, 1)

	if !p.Equals(m.MultiplyTuple(tu)) {
		t.Fatal("Matrix should equal its product by a tuple")
	}
}

func TestMultiplyIdentityByMatrix(t *testing.T) {
	m := NewMatrix(4, 4, []float64{0, 1, 2, 4, 1, 2, 4, 8, 2, 4, 8, 16, 4, 8, 16, 32})

	if !m.Equals(m.Multiply(IdentityMatrix())) {
		t.Fatal("Matrix should equal its product by the identity matrix")
	}
}

func TestMultiplyIdentityByTuple(t *testing.T) {
	tu := NewTuple(1, 2, 3, 4)

	if !tu.Equals(IdentityMatrix().MultiplyTuple(tu)) {
		t.Fatal("Matrix should equal its product by the identity tuple")
	}
}

func TestTransposeMatrix(t *testing.T) {
	m := NewMatrix(4, 4, []float64{0, 9, 3, 0, 9, 8, 0, 8, 1, 8, 5, 3, 0, 0, 5, 8})
	tm := NewMatrix(4, 4, []float64{0, 9, 1, 0, 9, 8, 8, 0, 3, 0, 5, 5, 0, 8, 3, 8})

	tr := m.Transpose()

	if !tr.Equals(tm) {
		t.Fatal("Matrix should equal its transpose manually calculated")
	}
}

func TestTransposeIdentity(t *testing.T) {
	tr := IdentityMatrix().Transpose()

	if !tr.Equals(IdentityMatrix()) {
		t.Fatal("The transpose of the identity should be itself")
	}
}

func TestDeterminant2x2(t *testing.T) {
	m := NewMatrix(2, 2, []float64{1, 5, -3, 2})

	det := m.Det()

	if det != 17 {
		t.Fatal("Determinant for 2x2 matrices is wrong")
	}
}

func TestSubmatrixOf3x3Is2x2(t *testing.T) {
	m := NewMatrix(3, 3, []float64{1, 5, 0, -3, 2, 7, 0, 6, -3})
	sm := NewMatrix(2, 2, []float64{-3, 2, 0, 6})

	sub := m.SubMatrix(0, 2)

	if !sm.Equals(sub) {
		t.Fatal("SubMatrix for 3x3 matrices is wrong")
	}
}

func TestSubmatrixOf4x4Is3x3(t *testing.T) {
	m := NewMatrix(4, 4,
		[]float64{-6, 1, 1, 6, -8, 5, 8, 6, -1, 0,
			8, 2, -7, 1, -1, 1})
	sm := NewMatrix(3, 3, []float64{-6, 1, 6, -8, 8, 6, -7, -1, 1})

	sub := m.SubMatrix(2, 1)

	if !sm.Equals(sub) {
		t.Fatal("SubMatrix for 4x4 matrices is wrong")
	}
}

func TestMinor3x3(t *testing.T) {
	m := NewMatrix(3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5})
	sm := m.SubMatrix(1, 0)

	subDet := sm.Det()
	mMinor := m.Minor(1, 0)

	if subDet != 25 {
		t.Fatal("Sub determinant for 3x3 matrices is wrong")
	}
	if mMinor != 25 {
		t.Fatal("Minor for 3x3 matrices is wrong")
	}
}

func TestCofactor3x3(t *testing.T) {
	m := NewMatrix(3, 3, []float64{3, 5, 0, 2, -1, -7, 6, -1, 5})

	mMinor1 := m.Minor(0, 0)
	mMinor2 := m.Minor(1, 0)
	mCofactor1 := m.Cofactor(0, 0)
	mCofactor2 := m.Cofactor(1, 0)

	if mMinor1 != -12 {
		t.Fatal("Minor for 3x3 matrices is wrong")
	}
	if mMinor2 != 25 {
		t.Fatal("Minor for 3x3 matrices is wrong")
	}
	if mCofactor1 != -12 {
		t.Fatal("Cofactor for 3x3 matrices is wrong")
	}
	if mCofactor2 != -25 {
		t.Fatal("Cofactor for 3x3 matrices is wrong")
	}
}

func TestDeterminant3x3(t *testing.T) {
	m := NewMatrix(3, 3, []float64{1, 2, 6, -5, 8, -4, 2, 6, 4})

	mCofactor1 := m.Cofactor(0, 0)
	mCofactor2 := m.Cofactor(0, 1)
	mCofactor3 := m.Cofactor(0, 2)
	det := m.Det()

	if mCofactor1 != 56 {
		t.Fatal("Cofactor is wrong")
	}
	if mCofactor2 != 12 {
		t.Fatal("Cofactor is wrong")
	}
	if mCofactor3 != -46 {
		t.Fatal("Cofactor is wrong")
	}
	if det != -196 {
		t.Fatal("Cofactor is wrong")
	}
}

func TestDeterminant4x4(t *testing.T) {
	m := NewMatrix(4, 4, []float64{-2, -8, 3, 5, -3, 1, 7, 3, 1, 2, -9, 6, -6, 7, 7, -9})

	mCofactor1 := m.Cofactor(0, 0)
	mCofactor2 := m.Cofactor(0, 1)
	mCofactor3 := m.Cofactor(0, 2)
	mCofactor4 := m.Cofactor(0, 3)
	det := m.Det()

	if mCofactor1 != 690 {
		t.Fatal("Cofactor is wrong")
	}
	if mCofactor2 != 447 {
		t.Fatal("Cofactor is wrong")
	}
	if mCofactor3 != 210 {
		t.Fatal("Cofactor is wrong")
	}
	if mCofactor4 != 51 {
		t.Fatal("Cofactor is wrong")
	}
	if det != -4071 {
		t.Fatal("Cofactor is wrong")
	}
}

func TestInvertible_yes(t *testing.T) {
	m := NewMatrix(4, 4, []float64{6, 4, 4, 4, 5, 5, 7, 6, 4, -9, 3, -7, 9, 1, 7, -6})

	if m.Det() != -2120 {
		t.Fatal("Determinant is wrong")
	}
	if !m.Invertible() {
		t.Fatal("Reported invertible matrix as non-invertible")
	}
}

func TestInvertible_no(t *testing.T) {
	m := NewMatrix(4, 4, []float64{-4, 2, -2, -3, 9, 6, 2, 6, 0, -5, 1, -5, 0, 0, 0, 0})

	if m.Det() != 0 {
		t.Fatal("Determinant is wrong")
	}
	if m.Invertible() {
		t.Fatal("Reported non-invertible matrix as invertible")
	}
}

func TestInverse(t *testing.T) {
	m := NewMatrix(4, 4, []float64{-5, 2, 6, -8, 1, -5, 1, 8, 7, 7, -6, -7, 1, -3, 7, 4})
	invExp := NewMatrix(4, 4, []float64{0.21805, 0.45113, 0.24060, -0.04511,
		-0.80827, -1.45677, -0.44361, 0.52068,
		-0.07895, -0.22368, -0.05263, 0.19737,
		-0.52256, -0.81391, -0.30075, 0.30639})

	mInv, err := m.Inverse()
	if err != nil {
		t.Fatal("Could not invert invertible matrix")
	}

	if m.Det() != 532 {
		t.Fatal("Determinant is wrong")
	}
	if m.Cofactor(2, 3) != -160 {
		t.Fatal("Cofactor is wrong")
	}
	if mInv.get(3, 2) != -160.0/532 {
		t.Fatal("Inverse is wrong")
	}
	if m.Cofactor(3, 2) != 105 {
		t.Fatal("Cofactor is wrong")
	}
	if mInv.get(2, 3) != 105.0/532 {
		t.Fatal("Inverse is wrong")
	}
	if !mInv.Equals(invExp) {
		t.Fatal("Inverse is wrong")
	}
}

func TestInverse_second(t *testing.T) {
	m := NewMatrix(4, 4, []float64{8, -5, 9, 2, 7, 5, 6, 1, -6, 0, 9, 6, -3, 0, -9, -4})
	invExp := NewMatrix(4, 4,
		[]float64{
			-0.15385, -0.15385, -0.28205, -0.53846, -0.07692, 0.12308,
			0.02564, 0.03077, 0.35897, 0.35897, 0.43590, 0.92308,
			-0.69231, -0.69231, -0.76923, -1.92308})

	mInv, err := m.Inverse()
	if err != nil {
		t.Fatal("Could not invert invertible matrix")
	}

	if !mInv.Equals(invExp) {
		t.Fatal("Inverse is wrong")
	}
}

func TestInverse_third(t *testing.T) {
	m := NewMatrix(4, 4, []float64{9, 3, 0, 9, -5, -2, -6, -3, -4, 9, 6, 4, -7, 6, 6, 2})
	invExp := NewMatrix(4, 4,
		[]float64{-0.04074, -0.07778, 0.14444, -0.22222,
			-0.07778, 0.03333, 0.36667, -0.33333,
			-0.02901, -0.14630, -0.10926, 0.12963,
			0.17778, 0.06667, -0.26667, 0.33333})

	mInv, err := m.Inverse()
	if err != nil {
		t.Fatal("Could not invert invertible matrix")
	}

	if !mInv.Equals(invExp) {
		t.Fatal("Inverse is wrong")
	}
}

func TestInverse_multiply(t *testing.T) {
	a := NewMatrix(4, 4, []float64{3, -9, 7, 3, 3, -8, 2, -9, -4, 4, 4, 1, -6, 5, -1, 1})
	b := NewMatrix(4, 4, []float64{8, 2, 2, 2, 3, -1, 7, 0, 7, 0, 5, 4, 6, -2, 0, 5})
	c := a.Multiply(b)

	bInv, err := b.Inverse()
	if err != nil {
		t.Fatal("Could not invert invertible matrix")
	}

	if !a.Equals(c.Multiply(bInv)) {
		t.Fatal("Inverse does not invert product")
	}
}
