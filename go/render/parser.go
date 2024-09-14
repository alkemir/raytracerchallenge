package render

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	r        io.Reader
	g        *Group
	vertices []Tuple
}

func NewParser(r io.Reader) *Parser {
	return &Parser{
		r: r,
		g: NewGroup(),
	}
}

func (p *Parser) DefaultGroup() *Group {
	return p.g
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
	if len(idxs) != 3 {
		return fmt.Errorf("wrong number of vertices %v", idxs)
	}

	numIdxs := make([]int, 3)
	for i, idx := range idxs {
		num, err := strconv.ParseInt(idx, 0, 0)
		if err != nil {
			return err
		}
		numIdxs[i] = int(num) - 1
	}

	s := NewTriangle(p.vertices[numIdxs[0]], p.vertices[numIdxs[1]], p.vertices[numIdxs[2]])
	p.g.Add(s)

	return nil
}
