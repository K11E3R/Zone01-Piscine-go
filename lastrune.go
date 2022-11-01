package piscine

func LastRune(s string) rune {
	for index, value := range s {
		for i := index; i < len(s); i++ {
			if index == len(s)-1 {
				return value
			}
		}
	}
	return 0
}
