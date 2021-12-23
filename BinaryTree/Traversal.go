package BinaryTree

//preorder: recursion
func preorderTraversalNew(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return res
}

//inorder: recursion
func inorderTraversalNew(root *TreeNode) []int {
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

//postorder: recursion
func postorderTraversalNew(root *TreeNode) []int {
	var res[] int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)

	return res
}

//preorder: iteration
func preorderTraversalNew1(root *TreeNode) []int {
	var res []int
	preOrderStack := []*TreeNode{}
	for root != nil || len(preOrderStack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			preOrderStack = append(preOrderStack, root)
			root = root.Left
		}
		node := preOrderStack[len(preOrderStack)-1]
		preOrderStack = preOrderStack[:len(preOrderStack)-1]
		root = node.Right
	}

	return res
}

//inorder: iteration
func inorderTraversalNew1(root *TreeNode) []int {
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

//postorder: iteration
func postorderTraversalNew1(root *TreeNode) []int {
	var res []int
	postOrderStack := []*TreeNode{}
	visisted := &TreeNode{}
	for root != nil || len(postOrderStack) != 0 {
		for root != nil {
			postOrderStack = append(postOrderStack, root)
			root = root.Left
		}
		node := postOrderStack[len(postOrderStack)-1]
		postOrderStack = postOrderStack[:len(postOrderStack)-1]
		if node.Right == nil || node.Right == visisted {
			res = append(res, node.Val)
			visisted = node
		} else {
			postOrderStack = append(postOrderStack, node)
			root = node.Right
		}
	}

	return res
}
