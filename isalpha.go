package piscine

func IsAlpha(s string) bool {
	var b bool = true
	for _, value := range s {
		if (value < 65 || value > 90) && (value < 97 || value > 122) && (value < 48 || value > 57) {
			b = false
		}
	}
	return b
}
