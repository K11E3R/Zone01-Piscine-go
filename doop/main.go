package main

import (
	"os"
)

func IsNumeric(s string) bool {
	var b bool = true
	for _, value := range s {
		if value < 48 || value > 57 {
			b = false
		}
	}
	return b
}

func Atoi(s string) int {
	r := 0
	d := 1
	if len(s) == 0 {
		r = 0
	} else {
		for i := len(s) - 1; i >= 0; i-- {
			if rune(s[i]) > '9' || rune(s[i]) < '0' {
				if i != 0 {
					r = 0
					break
				}
				continue
			} else {
				a := 0
				for j := rune(s[i]); j > '0'; j-- {
					a++
				}
				r = r + a*d
				d = d * 10
			}
		}
		if s[0] == '-' {
			r = -r
		}
	}
	return r
}

func inttostring(n int) string {
	var s []int
	var r []string
	var result string
	k := 0
	if n == 0 {
		r = append(r, "0")
	}
	if n == -9223372036854775808 {
		r = append(r, "-")
		n = n + 1
		n = -n
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			if i == k-1 {
				r = append(r, "8")
			} else {
				r = append(r, string('0'+s[k-1-i]))
			}
		}
	}
	if n < 0 {
		r = append(r, "-")
		n = -n
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			r = append(r, string('0'+s[k-1-i]))
		}
	}
	if n > 0 {
		for i := 1; n > 0; i++ {
			s = append(s, n%10)
			n /= 10
			k++
		}
		for i := range s {
			r = append(r, string('0'+s[k-1-i]))
		}
	}
	for i := 0; i < len(r); i++ {
		result = result + r[i]
	}
	result = result + "\n"
	return result
}

func main() {
	args := os.Args
	if len(args) == 4 && (args[2] == "/" || args[2] == "*" || args[2] == "-" || args[2] == "+" || args[2] == "%") && IsNumeric(args[1][1:]) && IsNumeric(args[3][1:]) {
		op := args[2]
		num1 := int64(Atoi(args[1]))
		num2 := int64(Atoi(args[3]))
		var r int64
		if op == "+" {
			r = num1 + num2
			if r <= 2147483647 && r >= -2147483648 {
				os.Stderr.WriteString(inttostring(int(r)))
			}
		}
		if op == "-" {
			r = num1 - num2
			if r <= 2147483647 && r >= -2147483648 {
				os.Stderr.WriteString(inttostring(int(r)))
			}
		}
		if op == "*" {
			r = num1 * num2
			if r <= 2147483647 && r >= -2147483648 {
				os.Stderr.WriteString(inttostring(int(r)))
			}
		}
		if op == "/" {
			if num2 != 0 {
				r = num1 / num2
				if r <= 2147483647 && r >= -2147483648 {
					os.Stderr.WriteString(inttostring(int(r)))
				}
			} else {
				os.Stderr.WriteString("No division by 0\n")
			}
		}
		if op == "%" {
			if num2 != 0 {
				r = num1 % num2
				if r <= 2147483647 && r >= -2147483648 {
					os.Stderr.WriteString(inttostring(int(r)))
				}
			} else {
				os.Stderr.WriteString("No modulo by 0\n")
			}
		}
	}
}
