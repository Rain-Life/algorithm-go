package BinaryTree

//solution-one: hashTable
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode{
	parentNodeMap := map[int]*TreeNode{} //记录每个结点的父节点
	visisted := map[int]bool{} // 记录p或q已经访问过的祖先结点

	//后续遍历实现记录每个结点的父节点
	var backTraverse func(*TreeNode)
	backTraverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parentNodeMap[node.Left.Val] = node
			backTraverse(node.Left)
		}
		if node.Right != nil {
			parentNodeMap[node.Right.Val] = node
			backTraverse(node.Right)
		}
	}
	backTraverse(root)

	for p != nil { //p向上遍历并记录访问过的祖先结点
		visisted[p.Val] = true
		p = parentNodeMap[p.Val]
	}
	for q != nil {
		if visisted[q.Val] {
			return q
		}
		q = parentNodeMap[q.Val]
	}

	return nil
}

//solution-two: flash back
func lowestCommonAncestor236(root, p, q *TreeNode) *TreeNode {
	// check
	if root == nil {
		return root
	}
	// 相等 直接返回root节点即可
	if root == p || root == q {
		return root
	}
	// Divide
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// Conquer
	// 左右两边都不为空，则根节点为祖先
	if left != nil && right != nil {
		return root
	}
	if left != nil {
		return left
	}
	if right != nil {
		return right
	}
	return nil
}
