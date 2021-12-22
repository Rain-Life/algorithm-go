package BinaryTree

//solution-one: DFS
func rightSideView(root *TreeNode) []int {
	rightViewMap := map[int]int{} //记录每一层第一个访问到的节点值（用map是为了知道每一层是否已经有访问过的结点）
	rightView := []int{} //最终存右视图结果的
	nodeStack := []*TreeNode{}
	depthStack := []int{}
	maxDepth := -1 //记录最大深度是为了最后变量右视图

	nodeStack = append(nodeStack, root) //根节点入栈
	depthStack = append(depthStack, 0) //跟结点属于第0层
	for len(nodeStack) != 0 {
		node := nodeStack[len(nodeStack)-1] //取出栈顶元素
		nodeStack = nodeStack[:len(nodeStack)-1]

		currDepth := depthStack[len(depthStack)-1] //当前节点属于哪一层
		depthStack = depthStack[:len(depthStack)-1]

		if node != nil {
			//维护二叉树的最大深度
			if currDepth > maxDepth {
				maxDepth = currDepth
			}

			//判断当前层，是否已经有节点被记录了.如果没有，才记录下来
			if _, ok := rightViewMap[currDepth];!ok {
				rightViewMap[currDepth] = node.Val
			}

			nodeStack = append(nodeStack, node.Left)
			nodeStack = append(nodeStack, node.Right)//右节点后入栈（后进先出）
			depthStack = append(depthStack, currDepth+1) //新进入的这两个结点的层数是相同的
			depthStack = append(depthStack, currDepth+1)
		}
	}
	for i:=0; i <= maxDepth; i++ {
		rightView = append(rightView, rightViewMap[i])
	}

	return rightView
}

//solution-two: BFS
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	rightViewMap := map[int]int{}
	rightView := []int{}
	maxDepth := 0

	nodeQueue := []*TreeNode{root}
	depthQueue := []int{}
	depthQueue = append(depthQueue, 0)

	for len(nodeQueue) != 0 {
		node := nodeQueue[0]
		if len(nodeQueue) <= 1 {
			nodeQueue = []*TreeNode{}
		} else {
			nodeQueue = nodeQueue[1:]
		}

		currDepth := depthQueue[0]
		if len(depthQueue) <= 1 {
			depthQueue = []int{}
		} else {
			depthQueue = depthQueue[1:]
		}

		if node != nil {
			if currDepth > maxDepth {
				maxDepth = currDepth
			}
			rightViewMap[currDepth] = node.Val

			nodeQueue = append(nodeQueue, node.Left)
			nodeQueue = append(nodeQueue, node.Right)
			depthQueue = append(depthQueue, currDepth+1)
			depthQueue = append(depthQueue, currDepth+1)
		}
	}

	for i := 0; i <= maxDepth; i++ {
		rightView = append(rightView, rightViewMap[i])
	}

	return rightView
}
