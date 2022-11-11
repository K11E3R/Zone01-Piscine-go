package piscine

func Max(a []int) int {
	maxx := a[0]
	for _, a := range a {
		if a > maxx {
			maxx = a
		}
	}
	return maxx
}
