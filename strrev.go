package piscine

func StrRev(s string) string {
	R := []byte(s)
	i := 0
	j := len(s) - 1
	for j >= 0 {
		R[j] = s[i]
		i++
		j--
	}
	return string(R)
}
