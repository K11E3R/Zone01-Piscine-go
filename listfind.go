package piscine

// type NodeL struct {
// 	Data interface{}
// 	Next *NodeL
// }

// type List struct {
// 	Head *NodeL
// 	Tail *NodeL
// }

func CompStr(a, b interface{}) bool {
	return a == b
}

func ListFind(l *List, ref interface{}, comp func(a, b interface{}) bool) *interface{} {
	t := l.Head
	for t != nil {
		if CompStr(t.Data, ref) {
			return &t.Data
		}
		t = t.Next
	}
	return nil
}
