package HighFrequency

// solution-two: hashtable
type ListNode struct {
	Val int
	Next *ListNode
}
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	mapNode := make(map[*ListNode]bool)
	for ;headA != nil; headA = headA.Next {
		mapNode[headA] = true
	}
	for ;headB != nil;headB = headB.Next {
		if _, ok := mapNode[headB]; ok {
			return headB
		}
	}

	return nil
}

// solution-two: two pointers
//双指针实现
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}

	pA, pB := headA, headB
	for pA != pB {
		if pA == nil {
			pA = headB
		} else {
			pA = pA.Next
		}
		if pB == nil {
			pB = headA
		} else {
			pB = pB.Next
		}
	}

	return pA
}