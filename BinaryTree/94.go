package BinaryTree


//solution-one: recursion
func inorderTraversal(root *TreeNode) []int {
	var inorderArr []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		inorderArr = append(inorderArr, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	return inorderArr
}

//solution-two: Iterator
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var inorderArr []int
	inorderStack := []*TreeNode{}
	for root != nil || len(inorderStack) != 0 {
		for root != nil {
			inorderStack = append(inorderStack, root)
			root = root.Left
		}
		root = inorderStack[len(inorderStack)-1]
		inorderStack = inorderStack[:len(inorderStack)-1]
		inorderArr = append(inorderArr, root.Val)
		root = root.Right
	}

	return inorderArr
}
