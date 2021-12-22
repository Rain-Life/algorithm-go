package BinaryTree

import "math"

//solution-one: inorder
func ClosestValue(root *TreeNode, target float64) int {
	var closestVal int
	min := float64(math.MaxInt64)
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		diffVal := math.Abs(float64(node.Val) - target)
		if diffVal < min {
			min = diffVal
			closestVal = node.Val
		}
		dfs(node.Right)
	}
	dfs(root)

	return closestVal
}

//solution-two: like binary search
func ClosestValue1(root *TreeNode, target float64) int {
	closestVal := root.Val
	for root != nil {
		if math.Abs(float64(closestVal )- target) >= math.Abs(float64(root.Val) - target) {
			closestVal = root.Val
		}

		if float64(root.Val) > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}

	return closestVal
}