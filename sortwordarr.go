package piscine

func SortWordArr(a []string) {
	for i := 0; i < len(a); i++ {
		for j := i; j < len(a); j++ {
			if a[j] < a[i] {
				s := a[j]
				a[j] = a[i]
				a[i] = s
			}
		}
	}
}
