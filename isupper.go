package piscine

func IsUpper(s string) bool {
	var b bool = true
	for _, value := range s {
		if value < 65 || value > 90 {
			b = false
		}
	}
	return b
}
