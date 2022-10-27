package piscine

func BasicAtoi2(s string) int {
	r := 0
	d := 1
	if s[0] == '-' {
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] > '9' || s[i] < '0' {
				return 0
			} else {
				a := 0
				for j := s[i]; j > '0'; j-- {
					a++
				}
				r = r + a*d
				d = d * 10
			}
		}
		r = -r
	}
	if s[0] == '+' {
		for i := len(s) - 1; i > 0; i-- {
			if s[i] > '9' || s[i] < '0' {
				return 0
			} else {
				a := 0
				for j := s[i]; j > '0'; j-- {
					a++
				}
				r = r + a*d
				d = d * 10
			}
		}
	}
	return r
}
