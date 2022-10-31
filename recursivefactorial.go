package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 2147483647 {
		return 0
	} else {
		if nb == 0 {
			return 1
		} else {
			nb = nb * RecursiveFactorial(nb-1)
		}
	}
	return nb
}
