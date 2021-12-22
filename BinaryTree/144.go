package BinaryTree

//solution-one: recursion
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return res
}

//solution-two: stack
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	preOrderStack := []*TreeNode{}
	for root != nil || len(preOrderStack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			preOrderStack = append(preOrderStack, root)
			root = root.Left
		}
		node := preOrderStack[len(preOrderStack)-1]
		preOrderStack = preOrderStack[:len(preOrderStack)-1]
		root = node.Right
	}

	return res
}
