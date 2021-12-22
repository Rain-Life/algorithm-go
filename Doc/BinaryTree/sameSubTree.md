## 判断t1树中是否有与t2树完全相同的子树

**题目来源**：[牛客网 - 判断t1树中是否有与t2树完全相同的子树](https://www.nowcoder.com/practice/4eaccec5ee8f4fe8a4309463b807a542?tpId=117&&tqId=37821&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking)

### 题目描述

给定彼此独立的两棵二叉树，树上的节点值两两不同，判断 t1 树是否有与 t2 树完全相同的子树

子树指一棵树的某个节点的全部后继节点

数据范围：树的节点数满足 0<*n*≤500000，树上每个节点的值一定在32位整型范围内

进阶：空间复杂度: O(1)，时间复杂度 O(n)

### 示例

**示例 1**

```go
输入：{1,2,3,4,5,6,7,#,8,9},{2,4,5,#,8,9}
返回值：true
```

**提示：**

- `1≤*n*≤500000`

## 解题

### 解法一：递归（前序遍历）

**思路**

判断一个树t2是不是另一颗树t1的子树，只要t2是以t1的左结点或右结点为根节点的树的子树就可以了，又是将一个大的问题，化解成了多个相同的子问题，就应该想到用递归来解。所以写成递归公式就应该是这样

```go
isContains(root1.Left, root2) || isContains(root1.Right, root2)
```

如果说t1中有一个结点和t2中的一个结点相等的时候，我们就可以分别继续比较t1和t2的左右子节点，直到t1或t2有一个遍历结束。翻译成代码就是

```go
func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}

	return root1.Val == root2.Val && isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}
```

所以，可以发现上边整个过程就是一个前序遍历，先比t1的根结点是否和t2的根结点相等，如果相等，继续比较这两个结点的子节点是否相等，当他们的左右节点都相等的时候，继续往下比较，直到t1或t2遍历完，比较结束的时候，看他们最后的结点是否相等，如果相等，就说明以t2是t1的子树

如果t1的根节点和t2的根节点不相等，则分别判断t1的左右子结点是否和t2相等，如果相等，重复上边的步骤

代码一看就懂，可以结合代码，理解上边的意思

**代码**

```go
// 判断t1树中是否有与t2树完全相同的子树
func isContains( root1 *TreeNode ,  root2 *TreeNode ) bool {
	if root1 == nil {
		return false
	}
	
	return isSame(root1, root2) || isContains(root1.Left, root2) || isContains(root1.Right, root2)
}

func isSame(root1 *TreeNode, root2 *TreeNode) bool {
	if root1 == nil || root2 == nil {
		return root1 == root2
	}

	return root1.Val == root2.Val && isSame(root1.Left, root2.Left) && isSame(root1.Right, root2.Right)
}
```