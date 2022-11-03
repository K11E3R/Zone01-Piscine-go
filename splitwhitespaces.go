package piscine

func SplitWhiteSpaces(s string) []string {
	var r []string
	k := 0
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if i == 0 {
				k = i + 1
				continue
			}
			if s[i+1] != ' ' && i != 0 {
				var m string
				for j := k; j < i; j++ {
					m = m + string(s[j])
				}
				r = append(r, m)
				k = i + 1
			}
			if s[i+1] == ' ' && i != 0 {
				var m string
				for j := k; j < i; j++ {
					m = m + string(s[j])
				}
				r = append(r, m)
				k = i + 2
				i = i + 1
			}
		}
	}
	var m string
	for j := k; j < len(s); j++ {
		m = m + string(s[j])
	}
	r = append(r, m)
	return r
}
