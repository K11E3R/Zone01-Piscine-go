package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	height := BTreeLevelCount(root)
	for level := 1; level <= height; level++ {
		BTreeApplyToGivenLevel(root, level, f)
	}
}

func BTreeApplyToGivenLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyToGivenLevel(root.Left, level-1, f)
		BTreeApplyToGivenLevel(root.Right, level-1, f)
	}
}

func BTreeLevelCount(root *TreeNode) int {
	countl := 0
	countr := 0

	if root != nil {
		countl = BTreeLevelCount(root.Left)
		countl++
		countr = BTreeLevelCount(root.Right)
		countr++
	}
	if countl > countr {
		return countl
	}
	return countr
}
