package BinaryTree


//solution-one: DFS
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

//solution-two: BFS
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	widthQueue := []*TreeNode{root}
	depth := 0

	for len(widthQueue) != 0 {
		count := len(widthQueue)
		for count > 0 {
			node := widthQueue[0]
			if len(widthQueue) > 1 {
				widthQueue = widthQueue[1:]
			} else {
				widthQueue = []*TreeNode{}
			}
			if node.Left != nil {
				widthQueue = append(widthQueue, node.Left)
			}
			if node.Right != nil {
				widthQueue = append(widthQueue, node.Right)
			}

			count--
		}
		depth++
	}

	return depth
}

