package render

import (
	"bufio"
	"io"
)

type Parser struct {
	r io.Reader
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
	return 1
}
