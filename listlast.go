package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListLast(l *List) interface{} {
	if l.Head == nil {
		return nil
	}
	current := l.Head
	for current.Next != nil {
		current = current.Next
	}
	return current.Data
}
