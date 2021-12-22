package BinaryTree

//solution-one: falsh back
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return helper(1, n)
}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTrees := []*TreeNode{}
	//枚举可行的根节点
	for i := start; i<=end; i++ {
		//获得所有可行的左子树集合
		leftTrees := helper(start, i-1)
		//获得所有可行的右子树集合
		rightTrees := helper(i+1, end)
		//从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, nil, nil}
				currTree.Left = left
				currTree.Right = right
				allTrees = append(allTrees, currTree)
			}
		}
	}

	return allTrees
}