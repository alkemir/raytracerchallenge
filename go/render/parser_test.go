package render

import (
	"strings"
	"testing"
)

func TestParserGibberish(t *testing.T) {
	gibberish := strings.NewReader(
		`There was a young lady name Bright` + "\n" +
			`who traveled much faster than light.` + "\n" +
			`She set out one day` + "\n" +
			`in a relative way,` + "\n" +
			`and came back the previous night.`)
	p := NewParser(gibberish)

	ignored := p.Parse()

	if ignored != 5 {
		t.Fatal("Parse is wrong")
	}
}

func TestParserVertices(t *testing.T) {
	gibberish := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1.000 0.5000 0.0000` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0`)
	p := NewParser(gibberish)

	p.Parse()

	if !p.vertices[0].Equals(NewPoint(-1, 1, 0)) {
		t.Fatal("Parse vertices is wrong")
	}
	if !p.vertices[1].Equals(NewPoint(-1, 0.5, 0)) {
		t.Fatal("Parse vertices is wrong")
	}
	if !p.vertices[2].Equals(NewPoint(1, 0, 0)) {
		t.Fatal("Parse vertices is wrong")
	}
	if !p.vertices[3].Equals(NewPoint(1, 1, 0)) {
		t.Fatal("Parse vertices is wrong")
	}
}

func TestParserTriangle(t *testing.T) {
	gibberish := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`` + "\n" +
			`f 1 2 3` + "\n" +
			`f 1 3 4`)
	p := NewParser(gibberish)

	p.Parse()
	t1 := p.DefaultGroup().children[0]
	t2 := p.DefaultGroup().children[1]

	if !t1.(*Triangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Parse triangle is wrong")
	}
	if !t1.(*Triangle).p2.Equals(p.vertices[1]) {
		t.Fatal("Parse triangle is wrong")
	}
	if !t1.(*Triangle).p3.Equals(p.vertices[2]) {
		t.Fatal("Parse triangle is wrong")
	}
	if !t2.(*Triangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Parse triangle is wrong")
	}
	if !t2.(*Triangle).p2.Equals(p.vertices[2]) {
		t.Fatal("Parse triangle is wrong")
	}
	if !t2.(*Triangle).p3.Equals(p.vertices[3]) {
		t.Fatal("Parse triangle is wrong")
	}
}

func TestParserPolygon(t *testing.T) {
	gibberish := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`v 0 2 0` + "\n" +
			`` + "\n" +
			`f 1 2 3 4 5`)
	p := NewParser(gibberish)

	p.Parse()
	t1 := p.DefaultGroup().children[0]
	t2 := p.DefaultGroup().children[1]
	t3 := p.DefaultGroup().children[2]

	if !t1.(*Triangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t1.(*Triangle).p2.Equals(p.vertices[1]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t1.(*Triangle).p3.Equals(p.vertices[2]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t2.(*Triangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t2.(*Triangle).p2.Equals(p.vertices[2]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t2.(*Triangle).p3.Equals(p.vertices[3]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t3.(*Triangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t3.(*Triangle).p2.Equals(p.vertices[3]) {
		t.Fatal("Parse polygon is wrong")
	}
	if !t3.(*Triangle).p3.Equals(p.vertices[4]) {
		t.Fatal("Parse polygon is wrong")
	}
}
