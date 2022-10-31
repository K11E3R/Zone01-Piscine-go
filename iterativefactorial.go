package piscine

func IterativeFactorial(nb int) int {
	r := 1
	if nb < 0 || nb > 2147483647 {
		return 0
	} else if nb > 0 && nb <= 2147483647 {
		for i := nb; i > 1; i-- {
			r = r * i
		}
	}
	return r
}
