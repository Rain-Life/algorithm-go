package HighFrequency
//这里主要是提供思路路，不考虑hash冲突的问题，实际场景中需要考虑

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