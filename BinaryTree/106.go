package BinaryTree

//solution-one: Recursion
func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	inorderMap := map[int]int{}
	for k, v := range inorder {
		inorderMap[v] = k
	}

	var build func(int, int) *TreeNode
	build = func(inorderLeft int, inorderRight int) *TreeNode {
		if inorderLeft > inorderRight {
			return nil
		}

		rootVal := postorder[len(postorder)-1]
		postorder = postorder[:len(postorder)-1]
		root := TreeNode{Val: rootVal}

		rootPos := inorderMap[rootVal]
		root.Right = build(rootPos+1, inorderRight)
		root.Left = build(inorderLeft, rootPos-1)

		return &root
	}

	return build(0, len(inorder)-1)
}
