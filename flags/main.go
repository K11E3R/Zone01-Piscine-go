package main

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := os.Args
	namePro := arguments[0]
	count := 0
	isOrder := false
	isHelp := false
	res := ""
	for range arguments {
		count++
	}
	if count == 1 {
		isHelp = true
	}
	for i := 1; i < count; i++ {
		if arguments[i] == "--order" || arguments[i] == "-o" {
			isOrder = true
		}
		if arguments[i] == "--help" || arguments[i] == "-h" {
			isHelp = true
		}
	}
	for i := 0; i < count; i++ {
		if arguments[i] != namePro {
			ok := false
			tmp := ""
			for j, a := range arguments[i] {
				if a == '=' && (tmp == "--insert" || tmp == "-i") {
					arguments[i] = arguments[i][j+1:]
					res = res + arguments[i]
					ok = true
					continue
				}
				tmp = tmp + string(a)
			}
			if arguments[i] == "--order" || arguments[i] == "-o" {
				continue
			}
			if !ok {
				res = arguments[i] + res
			}
		}
	}
	if isHelp {
		fmt.Print("--insert\n  -i\n\t This flag inserts the string into the string passed as argument.\n--order\n  -o\n\t This flag will behave like a boolean, if it is called it will order the argument.\n")
		return
	}
	if isOrder {
		res = SortStr(res)
	}
	for _, a := range res {
		z01.PrintRune(rune(a))
	}
	z01.PrintRune('\n')
}

func SortStr(s string) string {
	n := StrLen(s)
	rr := []rune(s)
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if rr[i] > rr[j] {
				tmp := rr[i]
				rr[i] = rr[j]
				rr[j] = tmp
			}
		}
	}
	res := ""
	for i := 0; i < n; i++ {
		res = res + string(rr[i])
	}
	return res
}

func StrLen(str string) int {
	newStr := []rune(str)
	count := 0
	for range newStr {
		count++
	}
	return count
}
