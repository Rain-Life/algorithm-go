package BinaryTree

//solution-one: twice traversal
func lowestCommonAncestor1(root, p, q *TreeNode) *TreeNode {
	pathP := getPath(root, p)
	pathQ := getPath(root, q)
	ancestor := root
	for i := 0; i < len(pathP) && i < len(pathQ) && pathP[i] == pathQ[i]; i++ {
		ancestor = pathP[i]
	}
	return ancestor
}

func getPath(root, target *TreeNode) []*TreeNode {
	path := []*TreeNode{}
	node := root
	for node != target {
		path = append(path, node)
		if target.Val < node.Val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	path = append(path, node)
	return path
}

// solution-two: once traversal
func lowestCommonAncestor3(root, p, q *TreeNode) *TreeNode {
	ancestor := root
	for  {
		if ancestor.Val > p.Val && ancestor.Val > q.Val {
			ancestor = ancestor.Left
		} else if ancestor.Val < p.Val && ancestor.Val < q.Val {
			ancestor = ancestor.Right
		} else {
			return ancestor
		}
	}
}
