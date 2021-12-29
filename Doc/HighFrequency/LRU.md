
## 设计LRU缓存结构

**题目来源**：[牛客网-**NC93 设计LRU缓存结构**](https://www.nowcoder.com/practice/e3769a5f49894d49b871c09cadd13a61?tpId=188&&tqId=38550&rp=1&ru=/activity/oj&qru=/ta/job-code-high-week/question-ranking)

### 题目描述

设计LRU(最近最少使用)缓存结构，该结构在构造时确定大小，假设大小为 k ，并有如下两个功能

1. set(key, value)：将记录(key, value)插入该结构

2. get(key)：返回key对应的value值

提示:

1.某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的，然后都会刷新缓存。

2.当缓存的大小超过k时，移除最不经常使用的记录。

3.输入一个二维数组与k，二维数组每一维有2个或者3个数字，第1个数字为opt，第2，3个数字为key，value

若opt=1，接下来两个整数key, value，表示set(key, value)若opt=2，接下来一个整数key，表示get(key)，若key未出现过或已被移除，则返回-1对于每个opt=2，输出一个答案

4.为了方便区分缓存里key与value，下面说明的缓存里key用""号包裹

**要求：set和get操作复杂度均为 O(1)**

### 示例

**示例 1**

```go
输入：[[1,1,1],[1,2,2],[1,3,2],[2,1],[1,4,4],[2,2]],3

返回值：[1,-1]

说明：
[1,1,1]，第一个1表示opt=1，要set(1,1)，即将(1,1)插入缓存，缓存是{"1"=1}
[1,2,2]，第一个1表示opt=1，要set(2,2)，即将(2,2)插入缓存，缓存是{"1"=1,"2"=2}
[1,3,2]，第一个1表示opt=1，要set(3,2)，即将(3,2)插入缓存，缓存是{"1"=1,"2"=2,"3"=2}
[2,1]，第一个2表示opt=2，要get(1)，返回是[1]，因为get(1)操作，缓存更新，缓存是{"2"=2,"3"=2,"1"=1}
[1,4,4]，第一个1表示opt=1，要set(4,4)，即将(4,4)插入缓存，但是缓存已经达到最大容量3，移除最不经常使用的{"2"=2}，插入{"4"=4}，缓存是{"3"=2,"1"=1,"4"=4}
[2,2]，第一个2表示opt=2，要get(2)，查找不到，返回是[1,-1]
```

**示例 2**

```go
输入：[[1,1,1],[1,2,2],[2,1],[1,3,3],[2,2],[1,4,4],[2,1],[2,3],[2,4]],2

返回值：[1,-1,-1,3,4]
```

## 解题

### **思路**

首先看LRU有哪些特性（假设用单链表实现）：

越靠后的结点，存的是越少访问的数据。因此，当有一个数据被访问之后，需要有以下两种情况考虑

1. 如果该数据在单链表中
此时，找到这个结点，并将它从原来的位置删除，然后插入到链表的头部
2. 如果该数据不在单链表中（有两种情况）
如果单链表满了：则删除链表的尾节点，并将该数据插入到链表的头部
如果单链表没满：则直接将该数据插入到单链表头部

首先实现LRU，最想想到的就是通过简单的单链表实现，越靠后的结点，表示越最不常访问的。题目中要求set和get的时间复杂度都是O(1)（也就是插入和查询），单纯的单链表，显然是做不到的，因为插入和获取的过程中，都涉及到查询，单链表的查询时间复杂度是O(n)，所以排除单链表实现

那就再想想其他的数据结构。要能够实现LRU，那么这个数据结构就必须有以下特性：

- 获取数据的时间复杂度是O(1)
- 插入数据的时间复杂度是O(1)
- 删除操作的时间复杂度是O(1)

看来数组好像可以，因为它支持下标的随机访问。但是在实现LRU插入的过程中可能会存在数据的移动，那插入操作的时间复杂度就做不到O(1)了

如果能利用数组的随机访问特性就好了，这样就能保住查询操作能满足要求，下边就是插入操作，只要能找到数据，那链表的单纯插入操作的时间复杂度就是O(1)。OK，那等于说利用数组的随机访问加上链表的插入删除操作，保证可以满足题目要求了

所以这里就可以使用**散列表+双向链表**。散列表利用了数组的随机访问特性，能保证查询的时间复杂度是O(1)。用双向链表的原因是，可以在O(1)的时间复杂度下，找到前驱结点（因为删除操作需要用到前驱结点）

这里通过链表法来解决散列冲突问题

最终这个结构就是长这样

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/LRU/1.png)

前驱（prev）和后继（next）指针是为了将结点串在双向链表中，hnext 指针是为了将结点串在散列表的拉链中（解决散列冲突问题）

然后看一下查询、插入、删除操作

- **查找一个数据**：当找到数据之后，我们还需要将它移动到双向链表的尾部
- **删除一个数据**：找到数据所在的结点，因为链表使用的是双向链表，可以通过前驱指针 O(1) 时间复杂度获取前驱结点
- **插入一条数据**：先看这个数据是否已经在缓存中。如果已经在其中，需要将其移动到双向链表的尾部；如果不在其中，还要看缓存有没有满。如果满了，则将双向链表头部的结点删除，然后再将数据放到链表的尾部；如果没有满，就直接将数据放到链表的尾部

### **代码**

```go
//这里主要是提供思路路，不考虑hash冲突的问题，实际场景中需要考虑

package LinkList

//散列表 + 双向链表 实现LRU

var head *LruNode // 双向链表头结点
var tail *LruNode // 双向链表尾结点

//定义双向链表的结点结构
type LruNode struct {
	Key string
	Value string
	pre *LruNode
	next *LruNode
	//hnext *Node // 用来解决哈希冲突的指针
}

//LRUCache结构
type LruCache struct {
	length int
	HashMap map[string]*LruNode
}

// 初始化LRU缓存
func Initialize(length int) *LruCache {
	LruCache := LruCache{}
	LruCache.length = length
	LruCache.HashMap = make(map[string]*LruNode, length)

	return &LruCache
}

//向双向链表中插入一个结点（在LRU这种特殊场景中，只会在头部插入）
func (Lru *LruCache) InsertNode(node *LruNode) {
	if head != nil {
		head.pre = node
		node.next = head
	} else if head == nil {
		head = node
		tail = node
	}
	head = node
}

//移除一个结点
func (Lru *LruCache) DelNode(node *LruNode) string {
	if node == tail {
		tail = tail.pre
	} else if node == head {
		head = head.next
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}

	return node.Key
}

//刷新缓存（当访问一个数据的时候，需要刷新一下缓存）
func (Lru *LruCache) RefreshCache(node *LruNode) {
	if node == tail {
		return
	}

	Lru.DelNode(node)
	Lru.InsertNode(node)
}

//获取LRU缓存数据
func (Lru *LruCache) Get(key string) string {
	if node, ok := Lru.HashMap[key]; ok {
		Lru.RefreshCache(node)
		return node.Value
	}

	return ""
}

func (Lru *LruCache) Set(key, value string) {
	if node, ok := Lru.HashMap[key]; !ok { //缓存里边没有
		if len(Lru.HashMap) >= Lru.length {
			removeNode := Lru.DelNode(head)
			delete(Lru.HashMap, removeNode)
		}
		newNode := LruNode{Key: key, Value: value}
		Lru.InsertNode(&newNode)
		Lru.HashMap[key] = &newNode
	} else {
		node.Value = value
		Lru.RefreshCache(node)
	}
}
```