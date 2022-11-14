package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListPushFront(l *List, data interface{}) {
	list := &NodeL{
		Next: l.Head,
		Data: data,
	}
	if l.Head != nil {
		l.Head.prev = list
	}
	l.Head = list

	L := l.Head
	for L.Next != nil {
		L = L.Next
	}
	l.Tail = L
}
