package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	s := ConcatParams(a)
	for i := 0; i < len(s); i++ {
		z01.PrintRune(rune(s[i]))
	}
	z01.PrintRune(rune('\n'))
}
