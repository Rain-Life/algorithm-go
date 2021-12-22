package BinaryTree

import "math"

//solution-one: recursion
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt64 //初始化全局最大和
	var maxValue func(*TreeNode) int
	maxValue = func(node *TreeNode) int { // 用于计算某个节点通往下方的最大路径和
		if node == nil {
			return 0
		}

		leftValue := max124(maxValue(node.Left), 0) //去除负数的情况
		rightValue := max124(maxValue(node.Right), 0)

		priceNewPath := node.Val + leftValue + rightValue //这两步的目的是为了在全是负数的情况下，不舍弃当前结点本身
		maxSum = max124(priceNewPath, maxSum)

		return node.Val + max124(leftValue, rightValue) //选左右节点的最大值加上其本身的值
	}
	maxValue(root)

	return maxSum
}

func max124(x, y int) int {
	if x > y {
		return x
	}

	return y
}