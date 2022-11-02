package piscine

func AtoiBase(s string, str string) int {
	indx := 0
	for _, res := range str {
		if res == '-' || res == '+' || count(str, string(res)) > 1 {
			indx = 1
			break
		}
	}
	if indx == 1 || len(str) < 2 {
		return 0
	} else {
		fin := 0
		for i, res := range s {
			ind := Index(str, string(res))
			fin += ind * RecursivePower(len(str), len(s)-1-i)
		}
		return fin
	}
}
