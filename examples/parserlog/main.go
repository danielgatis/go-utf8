package main

import (
	"fmt"
	"os"

	"github.com/danielgatis/go-utf8"
)

var _ utf8.Performer = (*performer)(nil)

type performer struct {
	output string
}

func (p *performer) CodePoint(r rune) {
	p.output += string(r)
}

func (p *performer) InvalidSequece() {
	p.output += "�"
}

func main() {
	performer := &performer{}
	parser := utf8.New(performer)

	data, err := os.ReadFile("./fixtures/UTF-8-demo.txt")

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	for _, b := range data {
		parser.Advance(b)
	}

	fmt.Println(performer.output)
}
