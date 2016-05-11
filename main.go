package main

import (
	"fmt"

	"github.com/willis7/Alice/parser"
)

func main() {
	s := `My Clippings.txt`
	clips := parser.Parse(s)

	fmt.Println(len(clips))

	for _, clip := range clips {
		fmt.Println(clip.String())
	}
}
