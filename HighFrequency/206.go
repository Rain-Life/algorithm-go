package HighFrequency

import (
	algorithm_go "algorithm-go"
	"algorithm-go/LinkList"
	"fmt"
)

//solution-one
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *algorithm_go.ListNode) *algorithm_go.ListNode {
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

//solution-two
func ReverseListHead(list *LinkList.List) {
	if list.HeadNode == nil {
		fmt.Println("链表为空")
		return
	}

	newList := &LinkList.List{}
	currentNode := list.HeadNode
	nextNode := currentNode.Next
	for currentNode!=nil {
		if newList.HeadNode == nil {
			newList.HeadNode = currentNode
			newList.HeadNode.Next = nil
			currentNode = nextNode
			continue
		}
		nextNode = currentNode.Next
		currentNode.Next = newList.HeadNode
		newList.HeadNode = currentNode
		currentNode = nextNode
	}
}

