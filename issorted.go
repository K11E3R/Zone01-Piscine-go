package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	var result bool = true
	for i := 1; i < len(a)-1; i++ {
		if (a[i-1]-a[i]) < 0 && (a[i]-a[i+1]) > 0 {
			result = false
			break
		}
	}
	return result
}
