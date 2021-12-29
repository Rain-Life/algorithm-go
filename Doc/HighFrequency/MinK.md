
## 最小的K个数

**题目来源**：[LeetCode-**剑指 Offer 40. 最小的k个数**](https://leetcode-cn.com/problems/zui-xiao-de-kge-shu-lcof/)

### 题目描述

输入整数数组 `arr` ，找出其中最小的 `k` 个数。例如，输入4、5、1、6、2、7、3、8这8个数字，则最小的4个数字是1、2、3、4

### 示例

**示例 1**

```go
输入：arr = [3,2,1], k = 2
输出：[1,2] 或者 [2,1]
```

**示例 2**

```go
输入：arr = [0,1,2,1], k = 1
输出：[0]
```

**限制**

- `0 <= k <= arr.length <= 10000`
- `0 <= arr[i] <= 10000`

## 解题

### 解法一：大顶堆

**思路**

求最小的K个数，跟求最大的K个树是一样的，很容易想到TopK，因此就不难想到用堆来解决这个问题。所以本题可以构建一个大小为K的大顶堆，具体过程是

1. 遍历数组中的每一个元素
2. 将数组中的前K个数初始化成一个大顶堆
3. 遍历数组中的剩余元素，如果数组中的值小于堆顶元素，则堆化到小顶堆中
4. 遍历大顶堆

用最容易的顺序思想来写的代码，主要是为了方便理解（关于堆，及堆的插入、删除，可以看[这里](https://juejin.cn/post/7025499503285403685) 的堆部分）

**代码**

```go
func MinK(arr []int, k int) []int {
	if k == 0 {
		return []int{}
	}

	length := len(arr)
	minK := make([]int, k+1) // 大顶堆(用数组来保存堆，下标为0的位置不存元素，这样是为了方便求左右结点及父节点)
	for j := 0;j < k; j++ { //arr中的前k个复制到minK中
		minK[j+1] = arr[j]
	}

	for i:=k/2; i>0; i-- { // 堆化大顶堆
		heapify(minK, k+1, i)
	}

	for x := k; x < length; x ++ {
		if arr[x] < minK[1] { //如果当前元素大于堆顶元素
			//移除堆顶元素（其实就是从堆中删除一个元素） 删除堆顶元素，然后将最后一个元素放到堆顶，然后再进行堆化
			count := len(minK)
			minK[1] = minK[count-1]
			count--
			heapify(minK, count, 1)

			//将当前元素插入到堆中
			count++
			minK[count-1] = arr[x]
			m := count-1
			for m/2 > 0 && minK[m] > minK[m/2] {
				minK[m], minK[m/2] = minK[m/2], minK[m]
				m = m/2
			}
		}
	}

	return minK[1:]
}

func heapify(a []int, n, i int) {
	for true {
		maxPos := i
		if 2*i < n && a[i] < a[2*i] {
			maxPos = 2*i
		}
		if 2*i+1 < n && a[maxPos] < a[2*i+1] {
			maxPos = 2*i+1
		}
		if maxPos == i {
			break
		}
		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}

```

### 解法二：快排的思想

**思路**

本题其实就是类似TopK的问题，自然想到用快排来解决。要解决本题，只要找到第K大的数字就可以了，左边的都是小与这个数字的，右边都是大于这个数字的。因为题目并不要求按顺序打印出最小的K个数，所以就不用排好序再输出了。具体过程是

1. 按照快排思想，找到一个分区点，将数组中大于该分区点元素的，放到分区点右边，小于的放左边
2. 判断分区点的位置与k的大小关系，如果分区点位置大于k，则取分区点的左边部分数组，重复步骤1
3. 如果分区点位置小于k，则取分区点的右边部分数组，重复步骤1

**代码**

```go
func QuickSortFindMinK(arr []int, k int) []int {
	if k <= 0 {
		return []int{}
	}

	quickSortFind(arr, 0, len(arr)-1, k)
	return arr[:k]
}

func quickSortFind(arr []int, start, end, k int) bool {
	for {
		if start > end{
			return false
		}
		pivot := partion(arr, start, end)
		if pivot == k {
			return true
		} else if pivot > k {
			end = pivot-1
		} else {
			start = pivot+1
		}
	}
}

func partion(arr []int, start, end int) int {
	pivotValue := arr[end]

	i := start
	for j:=start; j < len(arr); j++ {
		if arr[j] < pivotValue {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[end], arr[i] = arr[i], arr[end]

	return i
}
```