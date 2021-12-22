package BinaryTree

//solution-one: Preorder
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	nodeQueue := []*TreeNode{}
	nodeStack := []*TreeNode{}
	for len(nodeStack) != 0 || root != nil {
		for root != nil {
			nodeQueue = append(nodeQueue, root)
			nodeStack = append(nodeStack, root)
			root = root.Left
		}
		node := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		root = node.Right
	}
	for i:=1; i< len(nodeQueue); i++ {
		prev, curr := nodeQueue[i-1], nodeQueue[i]
		prev.Left, prev.Right = nil, curr
	}
}

//solution-two: preorder + open
func flatten1(root *TreeNode) {
	if root == nil {
		return
	}
	nodeStack := []*TreeNode{root}
	var prev *TreeNode
	for len(nodeStack) != 0 {
		curr := nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if prev != nil {
			prev.Left, prev.Right = nil, curr
		}

		leftNode, rightNode := curr.Left, curr.Right
		if rightNode != nil {
			nodeStack = append(nodeStack, rightNode)
		}
		if leftNode != nil {
			nodeStack = append(nodeStack, leftNode)
		}

		prev = curr
	}
}

//solution-three: find preNode
func flatten2(root *TreeNode) {
	if root == nil {
		return
	}
	curr := root
	for curr != nil {
		if curr.Left != nil {
			next := curr.Left
			prev := next
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = curr.Right
			curr.Left, curr.Right = nil, next
		}
		curr = curr.Right
	}
}
