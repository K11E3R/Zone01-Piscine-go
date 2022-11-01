package piscine

func IsLower(s string) bool {
	var b bool = true
	for _, value := range s {
		if value < 97 || value > 122 {
			b = false
		}
	}
	return b
}
