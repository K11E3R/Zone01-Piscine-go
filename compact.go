package piscine

func Compact(ptr *[]string) int {
	count := 0
	for i := 0; i < len(string); i++ {
		if (*ptr)[i] != " " {
			count++
		}
	}
	return count
}
