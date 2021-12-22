package BinaryTree

//solution-one: inorder
func recoverTree(root *TreeNode)  {
	if root == nil {
		return
	}

	var prevNode, xNode, yNode *TreeNode
	nodeStack := []*TreeNode{}
	for len(nodeStack)!=0 || root != nil {
		for root != nil {
			nodeStack = append(nodeStack, root)
			root = root.Left
		}
		root = nodeStack[len(nodeStack)-1]
		nodeStack = nodeStack[:len(nodeStack)-1]
		if prevNode != nil && root.Val < prevNode.Val {
			yNode = root
			if xNode == nil {
				xNode = prevNode
			} else {
				break
			}
		}
		prevNode = root
		root = root.Right
	}
	xNode.Val, yNode.Val = yNode.Val, xNode.Val
}

//solution-two: mirrors inorder
func recoverTree1(root *TreeNode)  {
	var x, y, pred, predecessor *TreeNode

	for root != nil {
		if root.Left != nil {
			// predecessor 节点就是当前 root 节点向左走一步，然后一直向右走至无法走为止
			predecessor = root.Left
			for predecessor.Right != nil && predecessor.Right != root {
				predecessor = predecessor.Right
			}

			// 让 predecessor 的右指针指向 root，继续遍历左子树
			if predecessor.Right == nil {
				predecessor.Right = root
				root = root.Left
			} else { // 说明左子树已经访问完了，我们需要断开链接
				if pred != nil && root.Val < pred.Val {
					y = root
					if x == nil {
						x = pred
					}
				}
				pred = root
				predecessor.Right = nil
				root = root.Right
			}
		} else { // 如果没有左孩子，则直接访问右孩子
			if pred != nil && root.Val < pred.Val {
				y = root
				if x == nil {
					x = pred
				}
			}
			pred = root
			root = root.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
