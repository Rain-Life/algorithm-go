package BinaryTree

//solution-one: DFS
func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	eachPath := []int{}
	var dfsPathSum func(root *TreeNode, last int)
	dfsPathSum = func(root *TreeNode, last int) {
		if root == nil {
			return
		}
		last -= root.Val
		eachPath = append(eachPath, root.Val)
		defer func() {
			eachPath = eachPath[:len(eachPath)-1] //一条路径走完之后，回退到上一步的路径，继续往下走，所以需要去掉路径中的最后一个
		}()
		if root.Left == nil && root.Right == nil && last == 0 {
			res = append(res, append([]int(nil), eachPath...))
			return
		}

		dfsPathSum(root.Left, last)
		dfsPathSum(root.Right, last)
	}
	dfsPathSum(root, targetSum)

	return res
}

//solution-two: BFS
type pairs struct {
	node *TreeNode
	left int
}
func pathSum2(root *TreeNode, targetSum int) (res [][]int) {
	if root == nil {
		return [][]int{}
	}

	parent := map[*TreeNode]*TreeNode{} //用于记录每个节点的父节点

	//通过叶子结点，回推到根节点，记录路径
	getPath := func(node *TreeNode) (path []int) {
		for ;node != nil; node = parent[node] {
			path = append(path, node.Val)
		} //这个路径是叶子结点到根节点的，所以需要逆序一下

		for i, j := 0, len(path)-1;i < j; i++ {
			path[i], path[j] = path[j], path[i]
			j--
		}
		return
	}

	//常规的广度优先搜索思路，借助队列来记录每一层的结点

	bfsQueue := []pairs{{root, targetSum}}
	for len(bfsQueue) != 0 {
		pair := bfsQueue[0]
		bfsQueue = bfsQueue[1:]
		node := pair.node
		left := pair.left - node.Val
		if node.Left == nil && node.Right == nil {
			if left == 0 {
				res = append(res, getPath(node))
			}
		} else {
			if node.Left != nil {
				parent[node.Left] = node
				bfsQueue = append(bfsQueue, pairs{node.Left, left})
			}
			if node.Right != nil {
				parent[node.Right] = node
				bfsQueue = append(bfsQueue, pairs{node.Right, left})
			}
		}
	}

	return res
}
