package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListSize(l *List) int {
	count := 0
	it := l.Head
	for it != nil {
		count++
		it = it.Next
	}
	return count
}
