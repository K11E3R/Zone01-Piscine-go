package piscine

func Sqrt(nb int) int {
	k := 0
	for i := nb; i >= 0; i = i - (2*k - 1) {
		if i == 0 {
			if nb == 0 {
				return 0
			} else {
				nb = Sqrt(nb-(2*k-1)) + 1
			}
			return nb
		}
		k++
	}
	return 0
}
