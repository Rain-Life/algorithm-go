package BinaryTree

import "strconv"

//solution-one: DFS
var treePath []string
func binaryTreePaths(root *TreeNode) []string {
	treePath = []string{}
	dfsEachPath(root, "")
	return treePath
}

func dfsEachPath(root *TreeNode, path string) {
	if root != nil {
		eachPath := path
		eachPath += strconv.Itoa(root.Val)
		if root.Left == nil && root.Right == nil {
			treePath = append(treePath, eachPath)
		} else {
			eachPath += "->"
			dfsEachPath(root.Left, eachPath)
			dfsEachPath(root.Right, eachPath)
		}
	}
}

//solution-two: BFS
func binaryTreePaths2(root *TreeNode) []string {
	paths := []string{}
	if root == nil {
		return paths
	}
	nodeQueue := []*TreeNode{root}
	pathQueue := []string{strconv.Itoa(root.Val)}
	for i :=0; i < len(nodeQueue); i++ {
		node, path := nodeQueue[i], pathQueue[i]
		if node.Left == nil && node.Right == nil {
			paths = append(paths, path)
			continue
		}
		if node.Left != nil {
			nodeQueue = append(nodeQueue, node.Left)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Left.Val))
		}
		if node.Right != nil {
			nodeQueue = append(nodeQueue, node.Right)
			pathQueue = append(pathQueue, path+"->"+strconv.Itoa(node.Right.Val))
		}
	}

	return paths
}
