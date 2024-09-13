package render

import (
	"strings"
	"testing"
)

func TestParserGibberish(t *testing.T) {
	gibberish := strings.NewReader(
		`There was a young lady name Bright
		who traveled much faster than light.
		She set out one day
		in a relative way,
		and came back the previous night.`)
	p := NewParser(gibberish)

	ignored := p.Parse()

	if ignored != 5 {
		t.Fatal("Parse is wrong")
	}
}
