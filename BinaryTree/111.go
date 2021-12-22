package BinaryTree

//solution-one: BFS
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minDepth := 1
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		tmpQueue := nodeQueue
		nodeQueue = []*TreeNode{}
		for len(tmpQueue) != 0 {
			node := tmpQueue[0]
			tmpQueue = tmpQueue[1:]
			if node.Left == nil && node.Right == nil {
				return minDepth
			}
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
		minDepth++
	}

	return minDepth
}
