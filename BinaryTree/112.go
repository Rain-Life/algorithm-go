package BinaryTree

//solution-one: DFS
func sumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}

func dfs(root *TreeNode, preNodeVal int) int {
	if root == nil {
		return 0
	}

	num := preNodeVal * 10 + root.Val
	if root.Left == nil && root.Right == nil {
		return num
	}

	return dfs(root.Left, num) + dfs(root.Right, num)
}

//solution-two: BFS
//广度优先搜索
func sumNumbers2(root *TreeNode) int {
	if root == nil {
		return 0
	}

	nodeQueue := []*TreeNode{root}
	numQueue := []int{root.Val}
	sum := 0
	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		if len(nodeQueue) > 1 {
			nodeQueue = nodeQueue[1:]
		} else {
			nodeQueue = []*TreeNode{}
		}
		num := numQueue[0]
		if len(numQueue) > 1 {
			numQueue = numQueue[1:]
		} else {
			numQueue = []int{}
		}

		leftNode, rightNode := node.Left, node.Right
		if leftNode == nil && rightNode == nil {
			sum += num
		} else {
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
				numQueue = append(numQueue, num * 10 + node.Left.Val)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
				numQueue = append(numQueue, num * 10 + node.Right.Val)
			}
		}
	}

	return sum
}
