package piscine

func BTreeApplyByLevel(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else if level > 1 {
		BTreeApplyByLevel(root.Left, level-1, f)
		BTreeApplyByLevel(root.Right, level-1, f)
	}
}
