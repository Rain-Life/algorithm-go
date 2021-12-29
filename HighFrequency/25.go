package HighFrequency

import algorithm_go "algorithm-go"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *algorithm_go.ListNode, k int) *algorithm_go.ListNode {
	dummy := &algorithm_go.ListNode{}
	pre := dummy
	end := dummy

	dummy.Next = head
	for end.Next != nil {
		for i :=0; i < k && end != nil; i++ {
			end = end.Next
		}
		if end == nil {
			break
		}

		start := pre.Next
		next := end.Next
		end.Next = nil //这个是为了方便反转这个k个节点
		pre.Next = reverseList1(start) //前边部分与刚反转完的这部分连接
		start.Next = next // 反转之后start变成了这个k组中的最后一个结点的位置，所以让他的Next指针指向next结点，将后边未翻转的连接起来
		pre = start
		end = pre
	}

	return dummy.Next
}

func reverseList1(head *algorithm_go.ListNode) *algorithm_go.ListNode {
	if head == nil  || head.Next == nil{
		return head
	}

	newHead := &algorithm_go.ListNode{}
	newHead.Next = head
	prevNode := newHead.Next
	currentNode := prevNode.Next
	for currentNode != nil {
		prevNode.Next = currentNode.Next
		currentNode.Next = newHead.Next
		newHead.Next = currentNode
		currentNode = prevNode.Next
	}
	head = newHead.Next

	return head
}