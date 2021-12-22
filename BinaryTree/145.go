package BinaryTree

//solution-one: recursion
func postorderTraversal(root *TreeNode) []int {
	var res[] int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)

	return res
}

//solution-two: iteration
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	postOrderStack := []*TreeNode{}
	visisted := &TreeNode{}
	for root != nil || len(postOrderStack) != 0 {
		for root != nil {
			postOrderStack = append(postOrderStack, root)
			root = root.Left
		}
		node := postOrderStack[len(postOrderStack)-1]
		postOrderStack = postOrderStack[:len(postOrderStack)-1]
		if node.Right == nil || node.Right == visisted {
			res = append(res, node.Val)
			visisted = node
		} else {
			postOrderStack = append(postOrderStack, node)
			root = node.Right
		}
	}

	return res
}
