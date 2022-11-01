package piscine

func FirstRune(s string) rune {
	for index, value := range s {
		if index == 0 {
			return value
		} else {
			return 0
		}
	}
	return 0
}
