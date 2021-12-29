package HighFrequency

import "algorithm-go/LinkList"

func hasCycle(head *LinkList.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}

	return false
}