
## 用两个栈实现队列

**题目来源**：[LeetCode-**剑指 Offer 09. 用两个栈实现队列**](https://leetcode-cn.com/problems/yong-liang-ge-zhan-shi-xian-dui-lie-lcof/)

### 题目描述

用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

### 示例

**示例 1**

```go
**输入**：["CQueue","appendTail","deleteHead","deleteHead"]
[[],[3],[],[]]
**输出**：[null,null,3,-1]
```

**示例 2**

```go
**输入**：
["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
[[],[],[5],[2],[],[]]
**输出**：[null,-1,null,null,5,2]
```

**提示：**

- `1 <= values <= 10000`
- `最多会对 appendTail、deleteHead 进行 10000 次调用`

## 解题

### **思路**

看到题目首先肯定是想到栈和队列的特性
**队列**：先进先出（FIFO）
**栈**：后进先出（LIFO）

要用栈实现队列，首先知道的是，栈只能从一个口进出元素，栈底是最先进去的元素。要想用栈实现队列，那就必须让栈底的元素先出来，这样就需要借助第二个栈。第一个栈复杂”队列“的插入操作，第二个栈负责删除操作，具体过程如下

- 插入操作，数据直接压入栈1
- 删除操作，如果栈2不为空，则将栈1中的元素，全部压入栈2中，然后取出栈2的栈顶元素

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/StackImplementQueue/1.png)

**说明**：在出队的时候，只有当stack2**为空**的时候，才会将stack1中的元素全部压入stack2中。具体看代码

### **代码**

```go
// 两个栈实现队列
type CQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() CQueue {
	return CQueue{
		stack1: []int{},
		stack2: []int{},
	}
}

// 入队都从stack1入
func (this *CQueue) AppendTail(value int)  {
	this.stack1 = append(this.stack1, value)
}

// 出队
func (this *CQueue) DeleteHead() int {
	if len(this.stack2) == 0 {
		for len(this.stack1) > 0 {
			this.stack2 = append(this.stack2, this.stack1[len(this.stack1)-1])
			this.stack1 = this.stack1[:len(this.stack1)-1]
		}
	}
	if len(this.stack2) > 0 {
		value := this.stack2[len(this.stack2)-1]
		this.stack2 = this.stack2[:len(this.stack2)-1]
		return value
	}

	return -1
}
```