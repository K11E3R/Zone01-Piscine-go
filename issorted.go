package piscine

func f(a, b int) int {
	return b - a
}

func IsSorted(f func(a, b int) int, a []int) bool {
	var result bool = true
	for i := 1; i < len(a)-1; i++ {
		if f(a[i], a[i-1]) < 0 && f(a[i+1], a[i]) > 0 {
			result = false
			break
		}
	}
	return result
}
