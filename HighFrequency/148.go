package HighFrequency

//solution-one: Merge Sort
func sortList(head *ListNode) *ListNode {
	return Sort(head, nil)
}

func Sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}
	if head.Next == tail {
		head.Next = nil
		return head
	}
	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	return MergeList(Sort(head, mid), Sort(mid, tail))
}

func MergeList(head1, head2 *ListNode) *ListNode {
	dummy := &ListNode{} // 哨兵头结点
	tmpNode, tmpNode1, tmpNode2 := dummy, head1, head2
	for tmpNode1 != nil && tmpNode2 != nil {
		if tmpNode1.Val <= tmpNode2.Val {
			tmpNode.Next = tmpNode1
			tmpNode1 = tmpNode1.Next
		} else {
			tmpNode.Next = tmpNode2
			tmpNode2 = tmpNode2.Next
		}

		tmpNode = tmpNode.Next
	}

	if tmpNode1 != nil {
		tmpNode.Next = tmpNode1
	}
	if tmpNode2 != nil {
		tmpNode.Next = tmpNode2
	}

	return dummy.Next
}