package BinaryTree

//solution-one: level traversal
type Node struct {
	Val int
	Left,Right, Next *Node

}
func connect(root *Node) *Node {
	if root == nil {
		return root
	}

	nodeQueue := []*Node{root}
	for len(nodeQueue) != 0 {
		tmpQueue := nodeQueue
		nodeQueue = nil
		for i, node := range tmpQueue {
			if i+1 < len(tmpQueue) {
				node.Next = tmpQueue[i+1]
			}
			if node.Left != nil {
				nodeQueue = append(nodeQueue, node.Left)
			}
			if node.Right != nil {
				nodeQueue = append(nodeQueue, node.Right)
			}
		}
	}
	return root
}

//solution-two: iteration
func connect1(root *Node) *Node {
	if root == nil {
		return root
	}

	//每一层从该层的最左节点开始遍历
	for leftMost := root;leftMost.Left!=nil; leftMost = leftMost.Left {
		//通过Next指针遍历该层的节点，为下一层的节点建立Next连接
		for node:=leftMost; node != nil; node = node.Next {
			node.Left.Next = node.Right
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}

	return root
}

