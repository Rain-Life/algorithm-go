package HighFrequency

import "algorithm-go/LinkList"

//solution-one: hashtable
func detectCycle(head *LinkList.Node) *LinkList.Node {
	mapNode := map[*LinkList.Node]string{}
	for head != nil {
		if _, ok := mapNode[head]; ok {
			return head
		}
		mapNode[head] = "reached"
		head = head.Next
	}

	return nil
}

// solution-two: two pointers
func detectCycle1(head *LinkList.Node) *LinkList.Node {
	if head == nil || head.Next == nil {
		return nil
	}
	if head.Next == head {
		return head
	}

	fast, slow := head, head
	for fast != nil {
		if fast.Next == nil {
			return nil
		}
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			ptr := head
			for ptr != slow {
				ptr = ptr.Next
				slow = slow.Next
			}
			return ptr
		}
	}

	return nil
}