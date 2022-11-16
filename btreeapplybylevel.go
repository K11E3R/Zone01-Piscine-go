package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := cc(root)
	for level := 1; level <= height; level++ {
		ttt(root, level, f)
	}
}

func ttt(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		ttt(root.Left, level-1, f)
		ttt(root.Right, level-1, f)
	}
}

func cc(root *TreeNode) int {
	countl := 0
	countr := 0

	if root != nil {
		countl = cc(root.Left)
		countl++
		countr = cc(root.Right)
		countr++
	}
	if countl > countr {
		return countl
	}
	return countr
}
