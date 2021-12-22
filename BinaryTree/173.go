package BinaryTree

//solution-one: recursion
type BSTIterator struct {
	iteratorArr []int
}

func constructor(root *TreeNode) BSTIterator {
	iterator := BSTIterator{}
	iterator.inorderTraverse(root)

	return iterator
}

func (iterator *BSTIterator) inorderTraverse(node *TreeNode) {
	if node == nil {
		return
	}
	iterator.inorderTraverse(node.Left)
	iterator.iteratorArr = append(iterator.iteratorArr, node.Val)
	iterator.inorderTraverse(node.Right)
}

func (this *BSTIterator) Next() int {
	value := this.iteratorArr[0]
	this.iteratorArr = this.iteratorArr[1:]

	return value
}

func (this *BSTIterator) HasNext() bool {
	return len(this.iteratorArr) > 0
}