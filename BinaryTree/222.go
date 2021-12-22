package BinaryTree

//solution-one: traversal
func countNodes(root *TreeNode) int {
	var count int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		count++
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)
	return count
}

//solution-two: binary search

