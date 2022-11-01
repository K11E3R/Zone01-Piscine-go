package piscine

func Capitalize(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if IsLower(string(s[0])) && IsAlpha(string(s[0])) {
				r[0] = s[0] - 32
			} else {
				r[0] = s[0]
			}
		} else {
			if !IsAlpha(string(s[i-1])) && IsLower(string(s[i])) {
				r[i] = s[i] - 32
			} else {
				if IsAlpha(string(s[i-1])) && IsUpper(string(s[i])) {
					r[i] = s[i] + 32
				} else {
					r[i] = s[i]
				}
			}
		}
	}
	return string(r)
}
