package BinaryTree

//solution-one: recursion
func kthSmallest(root *TreeNode, k int) int {
	var position, res int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		position++
		if position == k {
			res = node.Val
		}
		dfs(node.Right)

	}
	dfs(root)

	return res
}