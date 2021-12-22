package BinaryTree

import "math"

//solution-one: recursion
func isValidBST(root *TreeNode) bool {
	return isValid(root, math.MinInt64, math.MaxInt64)
}

func isValid(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if root.Val <= min || root.Val >= max {
		return false
	}

	// isValid(root.Left, min, root.Val) 左子树的值，都应该比当前这个节点小，因此它的下边界的值应该是当前结点的值
	//isValid(root.Right, root.Val, max) 右子树的值，都应该比当前这个节点大，因此它的上边界的值应该是当前结点的值
	return isValid(root.Left, min, root.Val) && isValid(root.Right, root.Val, max)
}

//solution-two: inorder
func isValidBST2(root *TreeNode) bool {
	nodeStack := []*TreeNode{}
	preNodeVal := math.MinInt64
	for len(nodeStack) != 0 || root != nil {
		for root != nil {
			nodeStack = append(nodeStack, root)
			root = root.Left
		}
		root = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if root.Val <= preNodeVal {
			return false
		}
		preNodeVal = root.Val
		root = root.Right
	}

	return true
}
