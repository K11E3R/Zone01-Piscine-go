package piscine

func ListAt(l *NodeL, pos int) *NodeL {
	if l.Data == nil {
		return nil
	}
	count := 0
	current := l
	for current != nil {
		if count == pos {
			return current
		}
		count++
		current = current.Next
	}
	return nil
}
