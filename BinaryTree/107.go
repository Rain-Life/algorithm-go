package BinaryTree


//solution-one: BFS
func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := [][]int{}
	nodeQueue := []*TreeNode{root}
	for len(nodeQueue) != 0 {
		levelVals := []int{}
		size := len(nodeQueue)
		for i := 0; i< size; i++ {
			node := nodeQueue[0]
			nodeQueue = nodeQueue[1:]
			levelVals = append(levelVals, node.Val)
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}

		res = append(res, levelVals)
	}

	//置换
	for j:=0; j < len(res)/2; j++ {
		res[j], res[len(res)-j-1] = res[len(res)-j-1], res[j]
	}

	return res
}