package BinaryTree

//solution-one: recursion
func upsideDownBinaryTree(root *TreeNode) *TreeNode {
	if root == nil || root.Left == nil {
		return root
	}
	//翻转前，拿到待反转的左子节点的兄弟节点
	leftNode, rightNode := root.Left, root.Right
	newRoot := upsideDownBinaryTree(leftNode) //对左子节点进行翻转
	leftNode.Left = rightNode
	leftNode.Right = root
	root.Left = nil //因为翻转后，原来的父节点的左右子节点要被替换了，所以这里在当前的左右子节点被处理后，将其置空
	root.Right = nil

	return newRoot
}