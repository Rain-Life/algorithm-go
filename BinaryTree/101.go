package BinaryTree

//solution-one: recursion
func IsSymmetric(root *TreeNode) bool {
	return checkSame(root.Left, root.Right)
}

func checkSame(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && checkSame(p.Left, q.Right) && checkSame(p.Right, q.Left)
}

//solution-two: iteration
func IsSymmetric1(root *TreeNode) bool {
	leftNode, rightNode := root, root
	nodeQueue := []*TreeNode{leftNode, rightNode}
	for len(nodeQueue) != 0 {
		leftNode = nodeQueue[0]
		rightNode = nodeQueue[1]
		nodeQueue = nodeQueue[2:]
		if leftNode == nil && rightNode == nil {
			continue
		}
		if leftNode == nil || rightNode == nil {
			return false
		}
		if leftNode.Val != rightNode.Val {
			return false
		}

		nodeQueue = append(nodeQueue, leftNode.Left)
		nodeQueue = append(nodeQueue, rightNode.Right)

		nodeQueue = append(nodeQueue, leftNode.Right)
		nodeQueue = append(nodeQueue, rightNode.Left)
	}

	return true
}
