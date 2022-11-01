package piscine

func IsNumeric(s string) bool {
	var b bool = true
	for _, value := range s {
		if value < 48 || value > 57 {
			b = false
		}
	}
	return b
}
