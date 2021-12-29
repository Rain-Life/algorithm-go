package HighFrequency

//solution-one:big Heap
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
		heapify1(minK, k+1, i)
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

func heapify1(a []int, n, i int) {
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

//solution-two:quick sort

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