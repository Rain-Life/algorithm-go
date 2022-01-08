package BackTracking

var path []int
var Res [][]int

func BackTracking(n, k, startIndex int) {
	if len(path) == k { //循环终止条件
		path1 := make([]int, k)
		copy(path1, path)
		Res = append(Res, path1)
		return
	}

	for i:=startIndex; i <= n; i++ {//某一层的每一个节点
		path = append(path, i)
		BackTracking(n, k, i+1) //往下一层递归
		path = path[:len(path)-1] //回溯
	}
}

