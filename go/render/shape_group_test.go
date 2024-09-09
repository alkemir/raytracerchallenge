package render

import (
	"testing"
)

func TestEmptyGroup(t *testing.T) {
	s := NewGroup()

	if !s.transform().Equals(IdentityMatrix()) {
		t.Fatal("Transform is wrong")
	}
	if len(s.Children()) != 0 {
		t.Fatal("Number of children is wrong")
	}
}

func TestRootShape(t *testing.T) {
	s := NewTestShape()

	if s.parent != nil {
		t.Fatal("Parent is wrong")
	}
}

func TestGroupAdd(t *testing.T) {
	g := NewGroup()
	s := NewTestShape()

	g.Add(s)

	if g.Children()[0] != s {
		t.Fatal("Children is wrong")
	}
	if s.parent == nil {
		t.Fatal("Parent is wrong")
	}
}
