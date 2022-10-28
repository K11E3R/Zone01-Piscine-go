package piscine

func BasicAtoi2(s string) int {
	for i := 0; i < len(s); i++ {
		if '0' > s[i] || s[i] > '9' {
			return 0
		}
	}
	r := 0
	d := 1
	for i := len(s) - 1; i >= 0; i-- {
		a := 0
		for j := s[i]; j > '0'; j-- {
			a++
		}
		r = r + a*d
		d = d * 10
	}
	return r
}
