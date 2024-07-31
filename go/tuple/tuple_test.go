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
		t.FailNow()
	}
	if !b.Equals(a) {
		t.FailNow()
	}
}

func TestNewPoint(t *testing.T) {
	p := NewPoint(4, -4, 3)

	if !p.Equals(NewTuple(4, -4, 3, 1)) {
		t.FailNow()
	}
}

func TestNewVector(t *testing.T) {
	p := NewVector(4, -4, 3)

	if !p.Equals(NewTuple(4, -4, 3, 0)) {
		t.FailNow()
	}
}
