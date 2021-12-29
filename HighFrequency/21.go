package HighFrequency

import "algorithm-go/LinkList"

func MergeLinkList(l1 *LinkList.Node, l2 *LinkList.Node) *LinkList.Node {
	if l1 == nil && l2 == nil {
		return nil
	}
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	curr1 := l1  //链表1当前指向的结点
	curr2 := l2  //链表2当前指向的结点
	var newHead *LinkList.Node
	if curr1.Data <= curr2.Data {
		newHead = curr1
		curr1 = curr1.Next
	} else {
		newHead = curr2
		curr2 = curr2.Next
	}

	moveNode := newHead
	for ;curr1 != nil && curr2 != nil;moveNode = moveNode.Next {
		if curr1.Data < curr2.Data {
			moveNode.Next = curr1
			curr1 = curr1.Next
		} else {
			moveNode.Next = curr2
			curr2 = curr2.Next
		}
	}

	if curr1 != nil {
		moveNode.Next = curr1
	}
	if curr2 != nil{
		moveNode.Next = curr2
	}

	return newHead
}