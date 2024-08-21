package tuple

import (
	"math"
	"testing"
)

func TestTupleIsPoint(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)

	if a.x != 4.3 {
		t.Fatal("Tuple X component did not match")
	}
	if a.y != -4.2 {
		t.Fatal("Tuple Y component did not match")
	}
	if a.z != 3.1 {
		t.Fatal("Tuple Z component did not match")
	}

	if !a.IsPoint() {
		t.Fatal("Tuple should be a point")
	}
	if a.IsVector() {
		t.Fatal("Tuple should not be a vector")
	}
}

func TestTupleIsVector(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 0.0)

	if a.x != 4.3 {
		t.Fatal("Tuple X component did not match")
	}
	if a.y != -4.2 {
		t.Fatal("Tuple Y component did not match")
	}
	if a.z != 3.1 {
		t.Fatal("Tuple Z component did not match")
	}

	if a.IsPoint() {
		t.Fatal("Tuple should not be a point")
	}
	if !a.IsVector() {
		t.Fatal("Tuple should be a vector")
	}
}

func TestTupleEquals(t *testing.T) {
	a := NewTuple(4.3, -4.2, 3.1, 1.0)
	b := NewTuple(4.3, -4.2, 3.1, 1.0)

	if !a.Equals(b) {
		t.Fatal("Equality is wrong")
	}
	if !b.Equals(a) {
		t.Fatal("Equality is wrong")
	}
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4, -4, 3)

	if !p.Equals(NewTuple(4, -4, 3, 1)) {
		t.Fatal("Point creation is wrong")
	}
}

func TestNewVector(t *testing.T) {
	p := NewVector(4, -4, 3)

	if !p.Equals(NewTuple(4, -4, 3, 0)) {
		t.Fatal("Vector creation is wrong")
	}
}

func TestAdd(t *testing.T) {
	p1 := NewTuple(3, -2, 5, 1)
	p2 := NewTuple(-2, 3, 1, 0)

	a1 := p1.Add(p2)
	a2 := p2.Add(p1)

	if !a1.Equals(a2) {
		t.Fatal("Addition is not symmetric")
	}
	if !a1.Equals(NewTuple(1, 1, 6, 1)) {
		t.Fatal("Addition is wrong")
	}
}

func TestSub(t *testing.T) {
	p1 := NewPoint(3, 2, 1)
	p2 := NewPoint(5, 6, 7)

	v := p1.Sub(p2)

	if !v.IsVector() {
		t.Fatal("Point minus point should yield a vector")
	}
	if !v.Equals(NewVector(-2, -4, -6)) {
		t.Fatal("Substraction is wrong")
	}
}

func TestSubPointVector(t *testing.T) {
	p := NewPoint(3, 2, 1)
	v := NewVector(5, 6, 7)

	rp := p.Sub(v)

	if !rp.IsPoint() {
		t.Fatal("Point minus vector should yield a point")
	}
	if !rp.Equals(NewPoint(-2, -4, -6)) {
		t.Fatal("Substraction is wrong")
	}
}

func TestSubVectors(t *testing.T) {
	v1 := NewVector(3, 2, 1)
	v2 := NewVector(5, 6, 7)

	rv := v1.Sub(v2)

	if !rv.IsVector() {
		t.Fatal("Vector minus vector should yield a vector")
	}
	if !rv.Equals(NewVector(-2, -4, -6)) {
		t.Fatal("Substraction is wrong")
	}
}

func TestSubFromZero(t *testing.T) {
	z := NewVector(0, 0, 0)
	v := NewVector(1, -2, 3)

	rv := z.Sub(v)

	if !rv.IsVector() {
		t.Fatal("Vector minus vector should yield a vector")
	}
	if !rv.Equals(NewVector(-1, 2, -3)) {
		t.Fatal("Substraction from zero is wrong")
	}
}

func TestNeg(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	na := a.Neg()

	if !na.Equals(NewTuple(-1, 2, -3, 4)) {
		t.Fatal("Tuple negation is wrong")
	}
}

func TestMulByScalar(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	m := a.Mul(3.5)

	if !m.Equals(NewTuple(3.5, -7, 10.5, -14)) {
		t.Fatal("Multiplication by scalar is wrong")
	}
}

func TestMulByFraction(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	m := a.Mul(0.5)

	if !m.Equals(NewTuple(0.5, -1, 1.5, -2)) {
		t.Fatal("Multiplication by fraction is wrong")
	}
}

func TestDivideByScalar(t *testing.T) {
	a := NewTuple(1, -2, 3, -4)

	m := a.Div(2)

	if !m.Equals(NewTuple(0.5, -1, 1.5, -2)) {
		t.Fatal("Division by scalar is wrong")
	}
}

func TestMag(t *testing.T) {
	v1 := NewVector(1, 0, 0)
	v2 := NewVector(0, 1, 0)
	v3 := NewVector(0, 0, 1)
	v4 := NewVector(1, 2, 3)
	v5 := NewVector(-1, -2, -3)

	if v1.Mag() != 1.0 {
		t.Fatal("Magnitude is wrong")
	}
	if v2.Mag() != 1.0 {
		t.Fatal("Magnitude is wrong")
	}
	if v3.Mag() != 1.0 {
		t.Fatal("Magnitude is wrong")
	}
	if v4.Mag() != math.Sqrt(14) {
		t.Fatal("Magnitude is wrong")
	}
	if v5.Mag() != math.Sqrt(14) {
		t.Fatal("Magnitude is wrong")
	}
}

func TestNorm(t *testing.T) {
	v1 := NewVector(4, 0, 0)
	v2 := NewVector(1, 2, 3)

	if !v1.Norm().Equals(NewVector(1, 0, 0)) {
		t.Fatal("Normalization is wrong")
	}
	if !v2.Norm().Equals(NewVector(1, 2, 3).Div(math.Sqrt(14))) {
		t.Fatal("Normalization is wrong")
	}
	if v2.Norm().Mag() != 1 {
		t.Fatal("Normalization is wrong")
	}
}

func TestDot(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	if v1.Dot(v2) != 20 {
		t.Fatal("Dot product is wrong")
	}
}

func TestCross(t *testing.T) {
	v1 := NewVector(1, 2, 3)
	v2 := NewVector(2, 3, 4)

	if !v1.Cross(v2).Equals(NewVector(-1, 2, -1)) {
		t.Fatal("Cross product is wrong")
	}
	if !v2.Cross(v1).Neg().Equals(NewVector(-1, 2, -1)) {
		t.Fatal("Cross product is not antisymmetric")
	}
}

func TestColorIsTuple(t *testing.T) {
	c := NewColor(-0.5, 0.4, 1.7)

	if c.x != -0.5 {
		t.Fatal("Color red component is wrong")
	}
	if c.y != 0.4 {
		t.Fatal("Color green component is wrong")
	}
	if c.z != 1.7 {
		t.Fatal("Color blue component is wrong")
	}
}

func TestColorAdd(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	if !c1.Add(c2).Equals(NewColor(1.6, 0.7, 1.0)) {
		t.Fatal("Color addition is wrong")
	}
}

func TestColorSub(t *testing.T) {
	c1 := NewColor(0.9, 0.6, 0.75)
	c2 := NewColor(0.7, 0.1, 0.25)

	if !c1.Sub(c2).Equals(NewColor(0.2, 0.5, 0.5)) {
		t.Fatal("Color substraction is wrong")
	}
}

func TestColorMul(t *testing.T) {
	c1 := NewColor(0.2, 0.3, 0.4)

	if !c1.Mul(2).Equals(NewColor(0.4, 0.6, 0.8)) {
		t.Fatal("Color multiplication by scalar is wrong")
	}
}

func TestColorsMultiply(t *testing.T) {
	c1 := NewColor(1, 0.2, 0.4)
	c2 := NewColor(0.9, 1, 0.1)

	if !c1.Hadamard(c2).Equals(NewColor(0.9, 0.2, 0.04)) {
		t.Fatal("Color multiplication is wrong")
	}
}
