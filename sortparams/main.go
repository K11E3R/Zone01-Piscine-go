package main

import (
	"os"

	"github.com/01-edu/z01"
)

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				s := a[j]
				a[j] = a[i]
				a[i] = s
			}
		}
	}
}

func main() {
	argument := os.Args
	SortWordArr(argument[1:])
	for i := 1; i < len(argument); i++ {
		for j := 0; j < len(argument[i]); j++ {
			z01.PrintRune(rune(argument[i][j]))
		}
		z01.PrintRune('\n')
	}
}
