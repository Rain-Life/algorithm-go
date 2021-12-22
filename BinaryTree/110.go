package BinaryTree

//solution-one: Top-Down
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(height(root.Left) - height(root.Right)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right) //先看以当前节点为根节点的树高度差，再分别判断左右子节点为根节点的左右子树的高度差
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max1(height(root.Left), height(root.Right)) + 1 //每往下遍历一层就+1
}

func max1(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}

//solution-two: Down-Up
func isBalanced2(root *TreeNode) bool {
	return height2(root) >= 0
}

func height2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftHeight := height(root.Left)
	rightHeight := height(root.Right)
	if leftHeight == -1 || rightHeight == -1 || abs1(leftHeight - rightHeight) > 1 {
		return -1
	}
	return max2(leftHeight, rightHeight) + 1
}

func max2(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func abs1(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}


