package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	var s []int
	k := 0
	if n == 0 {
		z01.PrintRune('0')
	}
	if n == -9223372036854775808 {
		z01.PrintRune('-')
		n = n + 1
		n = -n
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			if i == k-1 {
				z01.PrintRune(rune('8'))
			} else {
				z01.PrintRune(rune('0' + s[k-1-i]))
			}
		}
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			z01.PrintRune(rune('0' + s[k-1-i]))
		}
	}
	if n > 0 {
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			z01.PrintRune(rune('0' + s[k-1-i]))
		}
	}
}
