package main

import (
	"fmt"

	"github.com/willis7/Alice/parser"
)

func main() {
	s := `My Clippings.txt`
	clips := parser.Parse(s)

	fmt.Printf("clips = %#v \n", clips)
}
