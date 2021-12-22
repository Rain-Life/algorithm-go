package BinaryTree

//solution-one:recursion
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := TreeNode{preorder[0], nil, nil}
	//找到跟结点在中序遍历结果中的位置
	i := 0
	for ; i < len(inorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i]) //因为切片是左闭右开的，所以需要+1
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return &root
}