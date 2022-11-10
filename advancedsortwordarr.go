package piscine

func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if f(a[j], a[i]) == -1 {
				s := a[j]
				a[j] = a[i]
				a[i] = s
			}
		}
	}
}
