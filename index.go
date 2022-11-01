package piscine

func Index(s string, toFind string) int {
	r := 0
	if len(s) < len(toFind) {
		r = -1
	} else {
		for i := 1; i <= len(s)-len(toFind)+1; i++ {
			if s[i-1] == toFind[0] {
				test := s[i-1 : len(toFind)+i-1]
				if test == toFind {
					r = i - 1
					break
				}
			} else {
				r = -1
			}
		}
	}
	return r
}
