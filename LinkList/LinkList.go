package LinkList
//定义结点中数据的类型为接口类型，可收任意类型数据
type Object interface {}

//定义结点的结构体
type Node struct {
	Data int
	Next *Node
}

//定义链表的结构体
type List struct {
	HeadNode *Node
}
