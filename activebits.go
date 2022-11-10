package piscine

func ActiveBits(n int) (total int) {
	for ; n > 1; n /= 2 {
		total += n % 2
	}
	total += n
	return
}
