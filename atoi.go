package piscine

func Atoi(s string) int {
	r := 0
	if len(s) > 0 {
		for i := 1; i < len(s); i++ {
			if '0' > s[i] || s[i] > '9' {
				return r
			}
		}

		d := 1
		for i := len(s) - 1; i >= 0; i-- {
			a := 0
			for j := s[i]; j > '0'; j-- {
				a++
			}
			r = r + a*d
			d = d * 10
		}
		if s[0] == '-' {
			return -r
		}
	}
	return r
}
