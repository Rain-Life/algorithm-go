package HighFrequency

import "math"

//solution-one: order merge
//合并两个有序链表
func MergeTwoList23(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	dummyHead := &ListNode{}
	prev := dummyHead
	currNode1, currNode2 := head1, head2
	for currNode1 != nil && currNode2 != nil {
		if currNode1.Val <= currNode2.Val  {
			prev.Next = currNode1
			currNode1 = currNode1.Next
		} else {
			prev.Next = currNode2
			currNode2 = currNode2.Next
		}
		prev = prev.Next
	}

	if currNode1 == nil {
		prev.Next = currNode2
	}
	if currNode2 == nil {
		prev.Next = currNode1
	}

	return dummyHead.Next
}

//顺序合并 - 合并K个有序链表
func mergeKLists(lists []*ListNode) *ListNode {
	baseList := &ListNode{math.MinInt32, nil}
	for i := 0; i < len(lists); i++ {
		baseList = MergeTwoList(baseList, lists[i])
	}

	return baseList.Next
}


//solution-two: divide and conquer
//分治思想，实现k个有序链表的合并
func MergeKLists2(lists []*ListNode) *ListNode {
	return merge23(lists, 0, len(lists)-1)
}

func merge23(lists []*ListNode, l, r int) *ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}

	mid := (l+r) >> 2
	return MergeTwoList(merge23(lists, l, mid), merge23(lists, mid+1, r))
}

//合并两个有序链表
func MergeTwoList(head1 *ListNode, head2 *ListNode) *ListNode {
	if head1 == nil {
		return head2
	} else if head2 == nil {
		return head1
	}

	dummyHead := &ListNode{}
	prev := dummyHead
	currNode1, currNode2 := head1, head2
	for currNode1 != nil && currNode2 != nil {
		if currNode1.Val < currNode2.Val  {
			prev.Next = currNode1
			currNode1 = currNode1.Next
		} else {
			prev.Next = currNode2
			currNode2 = currNode2.Next
		}
		prev = prev.Next
	}

	if currNode1 == nil {
		prev.Next = currNode2
	}
	if currNode2 == nil {
		prev.Next = currNode1
	}

	return dummyHead.Next
}