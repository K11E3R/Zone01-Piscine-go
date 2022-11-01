package piscine

func ConcatParams(args []string) string {
	var r string
	for j := 0; j < len(args); j++ {
		if j == 0 {
			r = args[0] + "\n"
		} else {
			if j < len(args)-1 {
				r = r + args[j] + "\n"
			} else {
				r = r + args[j]
			}
		}
	}
	return r
}
