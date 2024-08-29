package render

import (
	"math"
	"testing"
)

func TestTranslatePoint(t *testing.T) {
	mat := Translation(5, -3, 2)
	p := NewPoint(-3, 4, 5)

	r := mat.MultiplyTuple(p)

	if !r.Equals(NewPoint(2, 1, 7)) {
		t.Fatal("Translation is wrong")
	}
}

func TestTranslatePoint_reverse(t *testing.T) {
	mat := Translation(5, -3, 2)
	tInv, err := mat.Inverse()
	if err != nil {
		t.Fatal("Inverse was wrong")
	}

	p := NewPoint(-3, 4, 5)

	r := tInv.MultiplyTuple(p)

	if !r.Equals(NewPoint(-8, 7, 3)) {
		t.Fatal("Translation is wrong")
	}
}

func TestTranslateVector(t *testing.T) {
	mat := Translation(5, -3, 2)
	v := NewVector(-3, 4, 5)

	r := mat.MultiplyTuple(v)

	if !r.Equals(v) {
		t.Fatal("Translation is wrong")
	}
}

func TestScalePoint(t *testing.T) {
	s := Scaling(2, 3, 4)
	p := NewPoint(-4, 6, 8)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(-8, 18, 32)) {
		t.Fatal("Translation is wrong")
	}
}

func TestScaleVector(t *testing.T) {
	s := Scaling(2, 3, 4)
	v := NewVector(-4, 6, 8)

	r := s.MultiplyTuple(v)

	if !r.Equals(NewVector(-8, 18, 32)) {
		t.Fatal("Scaling a vector is wrong")
	}
}

func TestScaleVector_reverse(t *testing.T) {
	s := Scaling(2, 3, 4)
	sInv, err := s.Inverse()
	if err != nil {
		t.Fatal("Inverse was wrong")
	}
	v := NewVector(-4, 6, 8)

	r := sInv.MultiplyTuple(v)

	if !r.Equals(NewVector(-2, 2, 2)) {
		t.Fatal("Scaling a vector is wrong")
	}
}

func TestReflectionIsNegativeScaling(t *testing.T) {
	s := Scaling(-1, 1, 1)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(-2, 3, 4)) {
		t.Fatal("Scaling a vector is wrong")
	}
}

func TestRotationX(t *testing.T) {
	halfQuarter := RotationX(math.Pi / 4)
	fullQuarter := RotationX(math.Pi / 2)
	p := NewPoint(0, 1, 0)

	r1 := halfQuarter.MultiplyTuple(p)
	r2 := fullQuarter.MultiplyTuple(p)

	if !r1.Equals(NewPoint(0, math.Sqrt2/2, math.Sqrt2/2)) {
		t.Fatal("Half a quarter rotation around X is wrong")
	}
	if !r2.Equals(NewPoint(0, 0, 1)) {
		t.Fatal("Full quarter rotation around X is wrong")
	}
}

func TestRotationX_reverse(t *testing.T) {
	halfQuarter := RotationX(math.Pi / 4)
	rInv, err := halfQuarter.Inverse()
	if err != nil {
		t.Fatal("Inverse was wrong")
	}

	p := NewPoint(0, 1, 0)

	r := rInv.MultiplyTuple(p)

	if !r.Equals(NewPoint(0, math.Sqrt2/2, -1*math.Sqrt2/2)) {
		t.Fatal("Rotating around X is wrong")
	}
}

func TestRotationY(t *testing.T) {
	halfQuarter := RotationY(math.Pi / 4)
	fullQuarter := RotationY(math.Pi / 2)
	p := NewPoint(0, 0, 1)

	r1 := halfQuarter.MultiplyTuple(p)
	r2 := fullQuarter.MultiplyTuple(p)

	if !r1.Equals(NewPoint(math.Sqrt2/2, 0, math.Sqrt2/2)) {
		t.Fatal("Half a quarter rotation around Y is wrong")
	}
	if !r2.Equals(NewPoint(1, 0, 0)) {
		t.Fatal("Full quarter rotation around Y is wrong")
	}
}

func TestRotationZ(t *testing.T) {
	halfQuarter := RotationZ(math.Pi / 4)
	fullQuarter := RotationZ(math.Pi / 2)
	p := NewPoint(0, 1, 0)

	r1 := halfQuarter.MultiplyTuple(p)
	r2 := fullQuarter.MultiplyTuple(p)

	if !r1.Equals(NewPoint(-1*math.Sqrt2/2, math.Sqrt2/2, 0)) {
		t.Fatal("Half a quarter rotation around Z is wrong")
	}
	if !r2.Equals(NewPoint(-1, 0, 0)) {
		t.Fatal("Full quarter rotation around Z is wrong")
	}
}

func TestShearing_xtoy(t *testing.T) {
	s := Shearing(1, 0, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(5, 3, 4)) {
		t.Fatal("Shearing X to Y is wrong")
	}
}

func TestShearing_xtoz(t *testing.T) {
	s := Shearing(0, 1, 0, 0, 0, 0)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(6, 3, 4)) {
		t.Fatal("Shearing X to Z is wrong")
	}
}

func TestShearing_ytox(t *testing.T) {
	s := Shearing(0, 0, 1, 0, 0, 0)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(2, 5, 4)) {
		t.Fatal("Shearing Y to X is wrong")
	}
}

func TestShearing_ytoz(t *testing.T) {
	s := Shearing(0, 0, 0, 1, 0, 0)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(2, 7, 4)) {
		t.Fatal("Shearing Y to Z is wrong")
	}
}

func TestShearing_ztox(t *testing.T) {
	s := Shearing(0, 0, 0, 0, 1, 0)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(2, 3, 6)) {
		t.Fatal("Shearing Z to X is wrong")
	}
}

func TestShearing_ztoy(t *testing.T) {
	s := Shearing(0, 0, 0, 0, 0, 1)
	p := NewPoint(2, 3, 4)

	r := s.MultiplyTuple(p)

	if !r.Equals(NewPoint(2, 3, 7)) {
		t.Fatal("Shearing Z to Y is wrong")
	}
}

func TestIndividualTransformations(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)

	p2 := A.MultiplyTuple(p)
	p3 := B.MultiplyTuple(p2)
	p4 := C.MultiplyTuple(p3)

	if !p2.Equals(NewPoint(1, -1, 0)) {
		t.Fatal("Rotation around X is wrong")
	}
	if !p3.Equals(NewPoint(5, -5, 0)) {
		t.Fatal("Scaling around X is wrong")
	}
	if !p4.Equals(NewPoint(15, 0, 7)) {
		t.Fatal("Translation around X is wrong")
	}
}

func TestChainedTransformations(t *testing.T) {
	p := NewPoint(1, 0, 1)
	A := RotationX(math.Pi / 2)
	B := Scaling(5, 5, 5)
	C := Translation(10, 5, 7)

	T := C.Multiply(B).Multiply(A)
	r := T.MultiplyTuple(p)

	if !r.Equals(NewPoint(15, 0, 7)) {
		t.Fatal("Chained transformation is wrong")
	}
}

func TestFluentTransformations(t *testing.T) {
	p := NewPoint(1, 0, 1)

	T := IdentityMatrix().RotateX(math.Pi/2).Scale(5, 5, 5).Translate(10, 5, 7)
	r := T.MultiplyTuple(p)

	if !r.Equals(NewPoint(15, 0, 7)) {
		t.Fatal("Fluent transformation is wrong")
	}
}
