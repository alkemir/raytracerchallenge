package render

import (
	"bufio"
	"io"
	"strconv"
	"strings"
)

type Parser struct {
	r        io.Reader
	vertices []Tuple
}

func NewParser(r io.Reader) *Parser {
	return &Parser{r: r}
}

func (p *Parser) Parse() int {
	ignored := 0
	scanner := bufio.NewScanner(p.r)
	for scanner.Scan() {
		ignored += p.parseLine(scanner.Text())
	}

	return ignored
}

func (p *Parser) parseLine(line string) int {
	tokens := strings.Split(line, " ")
	switch cmd := tokens[0]; cmd {
	case "v":
		return p.parseVertex(tokens[1:])
	}
	return 1
}

func (p *Parser) parseVertex(vxs []string) int {
	if len(vxs) != 3 {
		return 1
	}

	x, err := strconv.ParseFloat(vxs[0], 64)
	if err != nil {
		return 1
	}
	y, err := strconv.ParseFloat(vxs[1], 64)
	if err != nil {
		return 1
	}
	z, err := strconv.ParseFloat(vxs[2], 64)
	if err != nil {
		return 1
	}

	p.vertices = append(p.vertices, NewPoint(x, y, z))
	return 0
}
