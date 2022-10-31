package piscine

func FindNextPrime(nb int) int {
	if IsPrime(nb) == true {
		return nb
	} else {
		nb = FindNextPrime(nb + 1)
	}
	return nb
}
