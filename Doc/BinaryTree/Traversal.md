# 二叉树遍历

## 二叉树递归遍历

二叉树的递归实现很简单，只需要按照每种顺序去打印相关结点即可

-   前序遍历：当前节点、左节点、右节点
-   中序遍历：左节点、当前节点、右节点
-   后序遍历：左节点、右节点、当前节点

以下图的二叉树为例

![2.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/22b82dcc1aed47c080bf09dc2cf51a76~tplv-k3u1fbpfcp-watermark.image?)


树基础结构代码
```
type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}
```


### 前序遍历

前序遍历顺序：当前节点、左节点、右节点
```
func preorderTraversal(root *TreeNode) []int {
	var res []int
	var dfs func(*TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root)

	return res
}
```
遍历结果：A B D F E C G H I

### 中序遍历

中序遍历：左节点、当前节点、右节点
```
//递归实现
func inorderTraversal(root *TreeNode) []int {
	var inorderArr []int
	var inorder func(*TreeNode)
	inorder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inorder(node.Left)
		inorderArr = append(inorderArr, node.Val)
		inorder(node.Right)
	}
	inorder(root)

	return inorderArr
}
```

遍历结果：**D B E F A G H C I**

### 后后序遍历

```
func postorderTraversal(root *TreeNode) []int {
    var res[] int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		dfs(node.Right)
		res = append(res, node.Val)
	}
	dfs(root)

	return res
}
```
遍历结果：D E F B H G I C A


## 二叉树非递归遍历

二叉树的非递归遍历，需要借助栈这种数据结构。为什么会想到使用栈这种数据结构来实现？我个人觉得是因为递、归的过程，因为递归的过程就是最先走到的，最后才执行，比如斐波那契数列，假设我要求f(5)的值，那我就需要依次求f(5)、f(4)、f(3)、f(2)、f(1)，这是递的过程，归的过程就是知道f(1)，f(2)，然后回推到f(5)。虽然起初走到f(5)、f(4)，但是并不知道值，可以想象是把它压入了栈中，后边知道f(1)之后，再回头计算。很像栈的先进后出

知道了为什么使用栈来实现二叉树的非递归遍历，下边就看如何实现（还是以上边那棵二叉树为例）

### 中序遍历

**核心步骤：**

1.  当遍历到一个节点时，先将其压入栈，然后去遍历它的左子树（因为中序遍历的顺序是：左、中、右）
1.  如果左子树为空了，取出栈顶元素，并打印它的值
1.  遍历该元素的右子树（重复1、2、3）

**过程：**

-   根节点A，A入栈，左子树不为空，遍历左子树
-   节点B，B入栈，左子树不为空，遍历左子树
-   节点D，D入栈，左子树为空，取出栈顶元素D，并打印

![3.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/0786783eac364fe8956d1e17b284f7d8~tplv-k3u1fbpfcp-watermark.image?)

-   遍历D的右子树，发现右子树为空，则弹出栈顶元素B，并打印
-   遍历B的右子树F，F入栈，左子树不为空，遍历左子树
-   节点E入栈，发现E没有左右子树，弹出栈顶元素F，并访问

![4.png](https://p6-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/26a8247c51254e62b7d866d46d2e9556~tplv-k3u1fbpfcp-watermark.image?)

-   F没有右子树，弹出栈顶元素A，并访问
-   遍历A的右子树，所以C入栈
-   C有左子树，遍历C的左子树
-   ...(同上)

![5.png](https://p1-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/c785a9e02f034ad5a3685c67332945f7~tplv-k3u1fbpfcp-watermark.image?)

**代码实现：**

```
func inorderTraversal1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	var inorderArr []int
	inorderStack := []*TreeNode{}
	for root != nil || len(inorderStack) != 0 {
		for root != nil {
			inorderStack = append(inorderStack, root)
			root = root.Left
		}
		root = inorderStack[len(inorderStack)-1]
		inorderStack = inorderStack[:len(inorderStack)-1]
		inorderArr = append(inorderArr, root.Val)
		root = root.Right
	}

	return inorderArr
}
```
### 前序遍历

**核心步骤：**

前序跟中序的区别就是，打印节点的时机。前序遍历时在访问到当前结点时，先打印，再入栈。而中序遍历是，先入栈，第二次访问当前节点时，才打印

```
func preorderTraversal1(root *TreeNode) []int {
	var res []int
	preOrderStack := []*TreeNode{}
	for root != nil || len(preOrderStack) != 0 {
		for root != nil {
			res = append(res, root.Val)
			preOrderStack = append(preOrderStack, root)
			root = root.Left
		}
		node := preOrderStack[len(preOrderStack)-1]
		preOrderStack = preOrderStack[:len(preOrderStack)-1]
		root = node.Right
	}

	return res
}
```

### 后后序遍历

**核心思想：**

核心思想其实跟前边两个一样，后续遍历的访问顺序是左、右、中。后序遍历的时候，需要做个标记，从栈中弹出元素的时候，需要先判断它是否有右子树或者有没有被访问过。如果没有右子树，或者右子树被访问过，则直接打印当前结点。否则，将该节点重新放回队列，遍历该节点的又子树

代码实现：
```
func postorderTraversal1(root *TreeNode) []int {
	var res []int
	postOrderStack := []*TreeNode{}
	visisted := &TreeNode{}
	for root != nil || len(postOrderStack) != 0 {
		for root != nil {
			postOrderStack = append(postOrderStack, root)
			root = root.Left
		}
		node := postOrderStack[len(postOrderStack)-1]
		postOrderStack = postOrderStack[:len(postOrderStack)-1]
		if node.Right == nil || node.Right == visisted {
			res = append(res, node.Val)
			visisted = node
		} else {
			postOrderStack = append(postOrderStack, node)
			root = node.Right
		}
	}

	return res
}
```


## 二叉树层序遍历遍历

### 核心思想

层序遍历比较简单，需要借助队列这种数据结构（先进先出）

1.  根节点先入队列
1.  从队列中取出一个元素并打印
1.  将该元素的左右子节点依次放入队列
1.  重复2、3，直到队列为空

![6.png](https://p3-juejin.byteimg.com/tos-cn-i-k3u1fbpfcp/bf06125c1c0c4e0eb30422b4fdebf328~tplv-k3u1fbpfcp-watermark.image?)

### 代码实现

```
//二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	levelQueue := []*TreeNode{root}
	res := [][]int{}
	for len(levelQueue) != 0 {
		tmpValArr := []int{}
		i := 0
		n := len(levelQueue)
		for i < n {
			node := levelQueue[i]
			tmpValArr = append(tmpValArr, node.Val)
			if node.Left != nil {
				levelQueue = append(levelQueue, node.Left)
			}
			if node.Right != nil {
				levelQueue = append(levelQueue, node.Right)
			}
			i++
		}

		levelQueue = levelQueue[i:]
		res = append(res, tmpValArr)
	}

	return res
}
```


