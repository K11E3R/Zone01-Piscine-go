package piscine

func BTreeTransplant(root, node, repla *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	replacement := node
	if node.Parent == nil {
		root = repla
	} else if node == node.Parent.Left {
		replacement.Parent.Left = repla
	} else {
		replacement.Parent.Right = repla
	}
	if repla != nil {
		repla.Parent = node.Parent
	}

	return root
}
