package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"github.com/danielgatis/go-utf8"
)

var _ utf8.Performer = (*performer)(nil)

type performer struct {
	output string
}

func (p *performer) CodePoint(r rune) {
	fmt.Println(string(r))
}

func (p *performer) InvalidSequece() {
	fmt.Println("�")
}

func main() {
	performer := &performer{}
	parser := utf8.New(performer)

	reader := bufio.NewReader(os.Stdin)

	for {
		b, err := reader.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			}

			panic(err)
		}

		parser.Advance(b)
	}
}
