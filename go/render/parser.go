package render

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	r            io.Reader
	currentGroup string
	g            map[string]*Group
	vertices     []Tuple
	normals      []Tuple
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		r:            r,
		currentGroup: "",
		g:            map[string]*Group{"": NewGroup()},
	}
}

func (p *Parser) DefaultGroup() *Group {
	return p.g[""]
}

func (p *Parser) Group(group string) *Group {
	return p.g[group]
}

func (p *Parser) AsGroup() *Group {
	g := NewGroup()
	for _, subG := range p.g {
		g.Add(subG)
	}
	return g
}

func (p *Parser) Parse() int {
	ignored := 0
	scanner := bufio.NewScanner(p.r)
	for scanner.Scan() {
		if err := p.parseLine(scanner.Text()); err != nil {
			ignored += 1
		}
	}

	return ignored
}

func (p *Parser) parseLine(line string) error {
	tokens := strings.Split(line, " ")
	switch cmd := tokens[0]; cmd {
	case "v":
		return p.parseVertex(tokens[1:])
	case "f":
		return p.parseFace(tokens[1:])
	case "g":
		return p.parseGroup(tokens[1:])
	case "vn":
		return p.parseVertexNormal(tokens[1:])
	}
	return fmt.Errorf("command unknown: %v", tokens)
}

func (p *Parser) parseVertex(vxs []string) error {
	if len(vxs) != 3 {
		return fmt.Errorf("wrong number of vertices: %v", vxs)
	}

	x, errX := strconv.ParseFloat(vxs[0], 64)
	if errX != nil {
		return errX
	}
	y, errY := strconv.ParseFloat(vxs[1], 64)
	if errY != nil {
		return errY
	}
	z, errZ := strconv.ParseFloat(vxs[2], 64)
	if errZ != nil {
		return errZ
	}

	p.vertices = append(p.vertices, NewPoint(x, y, z))
	return nil
}

func (p *Parser) parseFace(idxs []string) error {
	vertices := len(idxs)
	numIdxs := make([]int, vertices)
	numNormals := make([]int, 0)
	for i, idx := range idxs {
		idxParts := strings.Split(idx, "/")

		num, err := strconv.ParseInt(idxParts[0], 0, 0)
		if err != nil {
			return err
		}
		numIdxs[i] = int(num) - 1

		if len(idxParts) != 1 {
			normal, err2 := strconv.ParseInt(idxParts[2], 0, 0)
			if err2 != nil {
				return err
			}
			numNormals = append(numNormals, int(normal)-1)
		}
	}

	g := p.g[p.currentGroup]
	for i := 0; i < vertices-2; i++ {
		var s Shape
		if len(numNormals) == 0 {
			s = NewTriangle(p.vertices[numIdxs[0]], p.vertices[numIdxs[i+1]], p.vertices[numIdxs[i+2]])
		} else if len(numNormals) == vertices {
			s = NewSmoothTriangle(p.vertices[numIdxs[0]], p.vertices[numIdxs[i+1]], p.vertices[numIdxs[i+2]],
				p.normals[numNormals[i]], p.normals[numNormals[i+1]], p.normals[numNormals[i+2]])
		} else {
			return fmt.Errorf("number of vertices did not match number of normals: %v %v", numIdxs, numNormals)
		}
		g.Add(s)
	}

	return nil
}

func (p *Parser) parseGroup(groupNames []string) error {
	if len(groupNames) != 1 {
		return fmt.Errorf("invalid group name %v", groupNames)
	}

	p.currentGroup = groupNames[0]
	g := p.g[p.currentGroup]
	if g == nil {
		p.g[p.currentGroup] = NewGroup()
	}
	return nil
}

func (p *Parser) parseVertexNormal(ns []string) error {
	if len(ns) != 3 {
		return fmt.Errorf("wrong number of normals: %v", ns)
	}

	x, errX := strconv.ParseFloat(ns[0], 64)
	if errX != nil {
		return errX
	}
	y, errY := strconv.ParseFloat(ns[1], 64)
	if errY != nil {
		return errY
	}
	z, errZ := strconv.ParseFloat(ns[2], 64)
	if errZ != nil {
		return errZ
	}

	p.normals = append(p.normals, NewVector(x, y, z))
	return nil
}
