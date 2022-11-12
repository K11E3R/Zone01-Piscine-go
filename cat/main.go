package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	size := len(os.Args)
	if len(os.Args) > 1 {
		for i := 1; i < size; i++ {
			data, err := ioutil.ReadFile(os.Args[i])
			if err != nil {
				for _, aaa := range err.Error() {
					z01.PrintRune(rune(aaa))
				}
				return
			}
			for _, aaa := range data {
				z01.PrintRune(rune(aaa))
			}
		}
		return
	}
	for true {
		for _, aaa := range "Hello" {
			z01.PrintRune(rune(aaa))
		}
		z01.PrintRune('\n')
	}
}
