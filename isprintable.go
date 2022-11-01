package piscine

func IsPrintable(s string) bool {
	var b bool = true
	for _, value := range s {
		if value < 32 || value > 127 {
			b = false
		}
	}
	return b
}
