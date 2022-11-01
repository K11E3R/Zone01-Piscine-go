package piscine

func ToLower(s string) string {
	r := []byte(s)
	for i := 0; i < len(s); i++ {
		if IsUpper(string(s[i])) {
			r[i] = s[i] + 32
		}
	}
	return string(r)
}
