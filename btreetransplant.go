package piscine

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {
	if node == nil {
		return root
	}
	if node.Parent.Left.Data == node.Data {
		node.Parent.Left = rplc
	} else if node.Parent.Right.Data == node.Data {
		node.Parent.Right = rplc
	}
	return root
}
