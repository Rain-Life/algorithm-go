package BinaryTree

import (
	"fmt"
	"strconv"
	"strings"
)

//solution-one: DFS
func Serialize(root *TreeNode) string {
	serializeStr := &strings.Builder{}
	dfsSerialize(root, serializeStr)

	return serializeStr.String()
}

func dfsSerialize(node *TreeNode, serializeStr *strings.Builder)  {
	if node == nil {
		serializeStr.WriteString("nil,")
		return
	}
	serializeStr.WriteString(strconv.Itoa(node.Val))
	serializeStr.WriteByte(',')
	dfsSerialize(node.Left, serializeStr)
	dfsSerialize(node.Right, serializeStr)
}

func Deserialize(data string) *TreeNode {
	deserializeArr := strings.Split(data, ",")
	var buildBinaryTree func() *TreeNode
	buildBinaryTree = func() *TreeNode {
		if len(deserializeArr) == 0 {
			return nil
		}
		if deserializeArr[0] == "nil" {
			deserializeArr = deserializeArr[1:]
			return nil
		}
		fmt.Printf("get：%v\n", deserializeArr)

		value, _ := strconv.Atoi(deserializeArr[0])
		deserializeArr = deserializeArr[1:]

		return &TreeNode{value, buildBinaryTree(), buildBinaryTree()}
	}

	return buildBinaryTree()
}

func buildBinaryTree(deserializeArr []string) *TreeNode {
	if deserializeArr[0] == "nil" {
		deserializeArr = deserializeArr[1:]
		return nil
	}
	fmt.Printf("get：%v\n", deserializeArr)

	value, _ := strconv.Atoi(deserializeArr[0])
	deserializeArr = deserializeArr[1:]

	return &TreeNode{value, buildBinaryTree(deserializeArr), buildBinaryTree(deserializeArr)}
}

//solution-two: Top-Down
//括号表示编码 + 递归下降解码
func Serialize1(root *TreeNode) string {
	if root == nil {
		return "X"
	}
	left := "(" + Serialize1(root.Left) + ")"
	right := "(" + Serialize1(root.Right) + ")"
	return left + strconv.Itoa(root.Val) + right
}

func Deserialize1(data string) *TreeNode {
	var parse func() *TreeNode
	parse = func() *TreeNode {
		if data[0] == 'X' {
			data = data[1:]
			return nil
		}
		node := &TreeNode{}
		data = data[1:] // 跳过左括号
		node.Left = parse()
		data = data[1:] // 跳过右括号
		i := 0
		//二叉树的结点可能是多位数
		for data[i] == '-' || '0' <= data[i] && data[i] <= '9' {
			i++
		}
		node.Val, _ = strconv.Atoi(data[:i]) //左子树解析完之后，获取节点值
		data = data[i:]
		//解析右子树
		data = data[1:] // 跳过左括号
		node.Right = parse()
		data = data[1:] // 跳过右括号
		return node
	}
	return parse()
}
