package BinaryTree

//solution-one: recursion
func isContains( root1 *TreeNode ,  root2 *TreeNode ) bool {
	if root1 == nil {
		return false
	}

	return isSame(root1, root2) || isContains(root1.Left, root2) || isContains(root1.Right, root2)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}

	return root1.Val == root2.Val && isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}
