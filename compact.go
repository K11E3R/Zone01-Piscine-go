package piscine

func Compact(ptr *[]string) int {
	count := 0
	for i := 0; i < ptr.lenght ; i++ {
		if (*ptr)[i] != " " {
			count++
		}
	}
	return count
}
