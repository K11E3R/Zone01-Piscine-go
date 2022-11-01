package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for j := 1; j < len(argument); j++ {
		r := []rune(argument[j])
		for i := range r {
			if rune(r[i]) != '.' && rune(r[i]) != '/' {
				z01.PrintRune(rune(r[i]))
			}
		}
		z01.PrintRune('\n')
	}
}
