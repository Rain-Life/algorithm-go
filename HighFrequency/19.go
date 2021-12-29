package HighFrequency

import "algorithm-go/LinkList"

//solution-one: stack
func removeNthFromEnd(head *LinkList.Node, n int) *LinkList.Node {
	stack := []*LinkList.Node{}
	dummy := &LinkList.Node{0, head}
	for curr:=dummy; curr!=nil;curr=curr.Next {
		stack = append(stack, curr)
	}
	pre := stack[len(stack) - n - 1] //获取前驱结点
	pre.Next = pre.Next.Next
	return dummy.Next
}

//solution-two: two pointers

func removeNthFromEnd1(head *LinkList.Node, n int) *LinkList.Node {
	dummy := &LinkList.Node{0, head} //这里加了一个哨兵结点的原因是为了方便找到待删除结点的前驱结点（链表的问题应该经常想到哨兵结点，或者叫虚拟头结点）
	first, second := head,  dummy
	//first先走n个节点
	for i:=0; i<n; i++ {
		first = first.Next
	}
	for ;first != nil; first = first.Next {
		second = second.Next //因为second在first的前一个位置开始遍历的，所以当first到尾部的时候，second就在倒数第n个节点的前一个位置
	}

	second.Next = second.Next.Next
	return head
}