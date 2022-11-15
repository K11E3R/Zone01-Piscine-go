package piscine

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
