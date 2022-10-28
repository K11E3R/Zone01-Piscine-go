package piscine

func SortIntegerTable(table []int) {
	for i := range table {
		for j := 0; j < len(table); j++ {
			if table[i] < table[j] {
				table[j], table[i] = table[i], table[j]
			}
		}
	}
}
