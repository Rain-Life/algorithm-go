package HighFrequency

//solution-one:quick sort
func FindK(arr []int, k int) int {
	if k == 0 {
		return 0
	}

	return QuickSortK(arr, 0, len(arr)-1, k)
}

func QuickSortK(arr []int, start, end, k int) int {
	for true {
		if start > end {
			return -1
		}

		pivot := partionK(arr, start, end)
		if pivot == k {
			return arr[k-1]
		} else if pivot < k {
			start = pivot + 1
		} else {
			end = pivot - 1
		}
	}

	return -1
}

func partionK(arr []int, start, end int) int {
	pivotValue := arr[end]

	i := start
	for j := start; j < len(arr); j++ {
		if arr[j] > pivotValue {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[end], arr[i] = arr[i], arr[end]

	return i
}

//solution-two: big heap
func HeapFinK(arr []int, k int) int {
	if k == 0 {
		return -1
	}

	length := len(arr)
	maxK := make([]int, k+1)
	for j := 0;j < k; j++ { //arr中的前k个复制到minK中
		maxK[j+1] = arr[j]
	}

	for i:=k/2; i>0; i-- { // 堆化小顶堆
		heapifyK(maxK, k+1, i)
	}

	for x := k; x < length; x ++ {
		if arr[x] > maxK[1] { //如果当前元素大于堆顶元素
			//移除堆顶元素（其实就是从堆中删除一个元素） 删除堆顶元素，然后将最后一个元素放到堆顶，然后再进行堆化
			count := len(maxK)
			maxK[1] = maxK[count-1]
			count--
			heapifyK(maxK, count, 1)

			//将当前元素插入到堆中
			count++
			maxK[count-1] = arr[x]
			m := count-1
			for m/2 > 0 && maxK[m] < maxK[m/2] {
				maxK[m], maxK[m/2] = maxK[m/2], maxK[m]
				m = m/2
			}
		}
	}

	return maxK[1]
}

func heapifyK(a []int, n, i int) {
	for true {
		maxPos := i
		if 2*i < n && a[i] > a[2*i] {
			maxPos = 2*i
		}
		if 2*i+1 < n && a[maxPos] > a[2*i+1] {
			maxPos = 2*i+1
		}
		if maxPos == i {
			break
		}
		a[i], a[maxPos] = a[maxPos], a[i]
		i = maxPos
	}
}
