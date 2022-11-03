package piscine

func AppendRange(min, max int) []int {
	r := []int(nil)
	if min < max {
		for i := min; i < max; i++ {
			r = append(r, i)
		}
	}
	return r
}
