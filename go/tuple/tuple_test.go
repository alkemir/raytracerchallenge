package tuple

import "testing"

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
