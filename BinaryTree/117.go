package BinaryTree

//solution-one: level Traversal
func connect2(root *Node) *Node {
	if root == nil {
		return  root
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
func connect3(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}