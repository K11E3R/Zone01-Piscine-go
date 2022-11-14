package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
	prev *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func (l *List) Reverse() {
	curr := l.Head
	var prev *NodeL
	l.Tail = l.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l.Head = prev
}

func ListPushBack(l *List, data interface{}) {
	l.Reverse()
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
	l.Reverse()
}
