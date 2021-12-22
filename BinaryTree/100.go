package BinaryTree

//solution-one: DFS
func isSameTree(p *TreeNode, q *TreeNode) bool {
	isSame := true
	var dfs func(*TreeNode, *TreeNode)
	dfs = func(pNode *TreeNode, qNode *TreeNode) {
		if pNode == nil && qNode == nil {
			return
		}
		if (pNode == nil && qNode != nil) || (qNode == nil && pNode != nil) {
			isSame = false
			return
		}
		if pNode.Val != qNode.Val {
			isSame = false
			return
		}
		dfs(pNode.Left, qNode.Left)
		dfs(pNode.Right, qNode.Right)

	}
	dfs(p, q)

	return isSame
}
