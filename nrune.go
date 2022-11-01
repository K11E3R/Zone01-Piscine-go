package piscine

func NRune(s string, n int) rune {
	if n <= len(s) {
		for index, value := range s {
			for i := index; i < len(s); i++ {
				if index == n-1 {
					return value
				}
			}
		}
	}
	return 0
}
