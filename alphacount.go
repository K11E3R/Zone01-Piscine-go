package piscine

func AlphaCount(s string) int {
	count := 0
	for i := 1; i <= len(s); i++ {
		condition := NRune(s, i)
		if condition >= 65 && condition <= 90 || condition >= 97 && condition <= 122 {
			count++
		}
	}
	return count
}
