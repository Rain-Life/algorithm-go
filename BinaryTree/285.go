package BinaryTree

// solution-one: iteration
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	flag := false
	nodeStack := []*TreeNode{}
	for len(nodeStack) != 0 || root != nil {
		for root != nil {
			nodeStack = append(nodeStack, root)
			root = root.Left
		}
		root = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if flag {
			return root
		}
		if root == p {
			flag = true
		}
		root = root.Right
	}

	return nil
}
