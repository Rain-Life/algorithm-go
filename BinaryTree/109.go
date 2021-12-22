package BinaryTree

//solution-one: recursion

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil{
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}
	return convertList(head, nil)
}

func convertList(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	//confirm mid Node
	slowPoint := left
	fastPoint := left
	for fastPoint != right && fastPoint.Next != right {
		slowPoint = slowPoint.Next
		fastPoint = fastPoint.Next.Next
	}

	root := &TreeNode{Val: slowPoint.Val}
	root.Left = convertList(left, slowPoint)
	root.Right = convertList(slowPoint.Next, right)

	return root
}
