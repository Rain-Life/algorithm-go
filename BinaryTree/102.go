package BinaryTree

//solution-one: BFS
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelQueue := []*TreeNode{root}
	res := [][]int{}
	for len(levelQueue) != 0 {
		tmpValArr := []int{}
		i := 0
		n := len(levelQueue)
		for i < n {
			node := levelQueue[i]
			tmpValArr = append(tmpValArr, node.Val)
			if node.Left != nil {
				levelQueue = append(levelQueue, node.Left)
			}
			if node.Right != nil {
				levelQueue = append(levelQueue, node.Right)
			}
			i++
		}

		levelQueue = levelQueue[i:]
		res = append(res, tmpValArr)
	}

	return res
}