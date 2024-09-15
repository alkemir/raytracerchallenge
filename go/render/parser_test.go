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
	fileContents := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1.000 0.5000 0.0000` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0`)
	p := NewParser(fileContents)

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
	fileContents := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`` + "\n" +
			`f 1 2 3` + "\n" +
			`f 1 3 4`)
	p := NewParser(fileContents)

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
	fileContents := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`v 0 2 0` + "\n" +
			`` + "\n" +
			`f 1 2 3 4 5`)
	p := NewParser(fileContents)

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

func TestParserNamedGroups(t *testing.T) {
	fileContents := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`` + "\n" +
			`g FirstGroup` + "\n" +
			`f 1 2 3` + "\n" +
			`g SecondGroup` + "\n" +
			`f 1 3 4`)
	p := NewParser(fileContents)

	p.Parse()
	g1 := p.Group("FirstGroup")
	g2 := p.Group("SecondGroup")
	t1 := g1.children[0]
	t2 := g2.children[0]

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
}

func TestParserAsGroup(t *testing.T) {
	fileContents := strings.NewReader(
		`v -1 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`v 1 1 0` + "\n" +
			`` + "\n" +
			`g FirstGroup` + "\n" +
			`f 1 2 3` + "\n" +
			`g SecondGroup` + "\n" +
			`f 1 3 4`)
	p := NewParser(fileContents)

	p.Parse()
	g := p.AsGroup()
	subGs := g.Children()

	if len(subGs) != 3 {
		t.Fatal("Subgroups is wrong")
	}
	for _, subG := range subGs {
		if _, isGroup := subG.(*Group); !isGroup {
			t.Fatal("Subgroups is wrong")
		}
	}
}

func TestParserVertexNormal(t *testing.T) {
	fileContents := strings.NewReader(
		`vn 0 0 1` + "\n" +
			`vn 0.707 0 -0.707` + "\n" +
			`vn 1 2 3`)
	p := NewParser(fileContents)

	p.Parse()

	if !p.normals[0].Equals(NewVector(0, 0, 1)) {
		t.Fatal("Vertex normal is wrong")
	}
	if !p.normals[1].Equals(NewVector(0.707, 0, -0.707)) {
		t.Fatal("Vertex normal is wrong")
	}
	if !p.normals[2].Equals(NewVector(1, 2, 3)) {
		t.Fatal("Vertex normal is wrong")
	}
}

func TestParserFacesWithNormals(t *testing.T) {
	fileContents := strings.NewReader(
		`v 0 1 0` + "\n" +
			`v -1 0 0` + "\n" +
			`v 1 0 0` + "\n" +
			`` + "\n" +
			`vn -1 0 0` + "\n" +
			`vn 1 0 0` + "\n" +
			`vn 0 1 0` + "\n" +
			`` + "\n" +
			`f 1//3 2//1 3//2` + "\n" +
			`f 1/0/3 2/102/1 3/14/2`)
	p := NewParser(fileContents)

	p.Parse()
	g := p.DefaultGroup()
	t1 := g.Children()[0]
	t2 := g.Children()[1]

	if !t1.(*SmoothTriangle).p1.Equals(p.vertices[0]) {
		t.Fatal("Vertex is wrong")
	}
	if !t1.(*SmoothTriangle).p2.Equals(p.vertices[1]) {
		t.Fatal("Vertex is wrong")
	}
	if !t1.(*SmoothTriangle).p3.Equals(p.vertices[2]) {
		t.Fatal("Vertex is wrong")
	}
	if !t1.(*SmoothTriangle).n1.Equals(p.normals[2]) {
		t.Fatal("Normal is wrong")
	}
	if !t1.(*SmoothTriangle).n2.Equals(p.normals[0]) {
		t.Fatal("Normal is wrong")
	}
	if !t1.(*SmoothTriangle).n3.Equals(p.normals[1]) {
		t.Fatal("Normal is wrong")
	}
	if !t1.(*SmoothTriangle).Equals(t2.(*SmoothTriangle)) {
		t.Fatal("Second triangle is wrong")
	}
}
