package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListRemoveIf(l *List, data_ref interface{}) {
	var v []interface{}
	t := l.Head
	for t != nil {
		if t.Data != data_ref {
			v = append(v, t.Data)
		}
		t = t.Next
	}
	i := 0
	if l.Head != nil {
		l.Head = nil
		l.Tail.Next = nil
	}

	for i < len(v) {
		list := &NodeL{
			Next: l.Head,
			Data: v[i],
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
		i++
	}
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
