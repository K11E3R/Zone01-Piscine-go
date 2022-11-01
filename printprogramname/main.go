package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	s := argument[0]
	r := []rune(s)
	for i := range r {
		if rune(r[i]) != '.' && rune(r[i]) != '/' {
			z01.PrintRune(rune(r[i]))
		}
	}
	z01.PrintRune('\n')
}
