package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func ListMerge(l1 *List, l2 *List) {
	var v []interface{}
	t := l1.Head
	for t != nil {
		v = append(v, t.Data)
		t = t.Next
	}
	t = l2.Head
	for t != nil {
		v = append(v, t.Data)
		t = t.Next
	}
	i := 0
	if l1.Head != nil {
		l1.Head = nil
		l1.Tail.Next = nil
	}

	for i < len(v) {
		list := &NodeL{
			Next: l1.Head,
			Data: v[i],
		}
		if l1.Head != nil {
			l1.Head.prev = list
		}
		l1.Head = list

		L := l1.Head
		for L.Next != nil {
			L = L.Next
		}
		l1.Tail = L
		i++
	}
	curr := l1.Head
	var prev *NodeL
	l1.Tail = l1.Head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	l1.Head = prev
}
