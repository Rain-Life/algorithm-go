
## 判断链表是否有环

**题目来源**：[牛客网-**NC4 判断链表中是否有环**](https://www.nowcoder.com/practice/650474f313294468a4ded3ce0f7898b9?tpId=117&&tqId=37714&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking)

### 题目描述

判断给定的链表中是否有环。如果有环则返回true，否则返回false

数据范围：链表长度 0 ≤ n ≤ 100000，链表中任意节点的值满足 |val| <= 100000

要求：空间复杂度 O(1)，时间复杂度 O(n)

输入分为2部分，第一部分为链表，第二部分代表是否有环，然后回组成head头结点传入到函数里面。-1代表无环，其他的数字代表有环，这些参数解释仅仅是为了方便读者自测调试。实际在编码时读入的是链表的头节点

例如输入{3,2,0,-4},1时，对应的链表结构如下图所示：

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/HaveRing/1.png)

可以看出环的入口结点为从头结点开始的第1个结点，所以输出true

### 示例

**示例 1**

```go
输入：{3,2,0,-4},1
返回值：true
说明：第一部分{3,2,0,-4}代表一个链表，第二部分的1表示，-4到位置1，即-4->2存在一个链接，组成传入的head为一个带环的链表 ,返回true
```

**示例 2**

```go
输入：{1},-1
返回值：false
说明：第一部分{1}代表一个链表，-1代表无环，组成传入head为一个无环的单链表，返回false
```

**示例 3**

```go
输入：{-1,-7,7,-4,19,6,-9,-5,-2,-5},6
返回值：true
```

## 解题

### **思路**

假设链表中有环，会是什么情况？就是无论怎样遍历都不会遍历到next等于空的情况。如何才能知道它是一直在循环遍历？有个暴力的方法就是，我遍历到每一个结点的时候，都去判断一下，它的next，有没有指向我前边已经遍历过的结点。但是这中方式，复杂度比较高

你假设有两个人在操场上跑步，一个人速度快，一个人速度慢（他俩的速度假设都不变），会出现什么情况？总有某一个时刻，他们会相遇。那把这种情况用在环的判断上，假设有两个指针，他们都遍历链表，但是步长不一样（每次遍历结点的个数），如果链表有环的话，他们一定会在某一个点相遇

比较好理解，这里就不画图了

### **代码**

```go
func hasCycle(head *LinkList.Node) bool {
	if head == nil || head.Next == nil {
		return false
	}
	if head.Next == head {
		return true
	}

	slow, fast := head, head
	for fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == nil {
			return false
		}
		if fast == slow {
			return true
		}
	}

	return false
}
```