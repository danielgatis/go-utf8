package utf8

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var _ Performer = (*performer)(nil)

type performer struct {
	output string
}

func (p *performer) CodePoint(r rune) {
	p.output += string(r)
}

func (p *performer) InvalidSequece() {
	p.output += "�"
}

func TestParser(t *testing.T) {
	performer := &performer{}
	parser := New(performer)

	data, err := os.ReadFile("./fixtures/UTF-8-demo.txt")

	if err != nil {
		t.Fatal(err)
	}

	for _, b := range data {
		parser.Advance(b)
	}

	assert.Equal(t, string(data), performer.output)
}
