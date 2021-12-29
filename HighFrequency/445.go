package HighFrequency
//栈实现
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	stack1 := []int{}
	stack2 := []int{}
	//head1, head2 := l1, l2
	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	//头插法创建新的链表
	var newHead *ListNode
	carry := 0
	for len(stack1) > 0 || len(stack2) > 0 || carry != 0{
		v1, v2 := 0, 0
		if len(stack1) > 0 {
			v1 = stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}
		if len(stack2) > 0 {
			v2 = stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}
		sum := v1 + v2 +carry
		carry = sum/10
		sum = sum % 10
		node := ListNode{sum, nil}
		node.Next = newHead
		newHead = &node
	}
	return newHead
}