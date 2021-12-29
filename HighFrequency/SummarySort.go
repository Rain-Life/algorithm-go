package HighFrequency

import "fmt"

// BubbleSort
func BubbleSort(arr []int)  {
	flag := false
	n := len(arr)
	for i:=0; i < n; i++ {
		flag = false
		for j:=0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			break
		}
	}

	fmt.Println(arr)
}

//InsertSort
func InsertSort(arr []int)  {
	n := len(arr)
	for i:=1; i < n; i++ { //i是待排区的元素
		value := arr[i]
		j := i-1
		for ; j>=0; j-- { //j遍历的是已排区的每一个元素
			if arr[j] > value {
				arr[j+1] = arr[j]
			}else {
				break
			}
		}
		arr[j+1] = value
	}

	fmt.Println(arr)
}

//SelectSort
func SelectSort(arr []int) {
	n := len(arr)

	for i:=0; i < n-1; i++ {
		for j:=i+1; j < n; j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	fmt.Println(arr)
}

//MergeSort
func MergeSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	medium := len(arr) / 2
	leftArr := MergeSort(arr[:medium])
	rightArr := MergeSort(arr[medium:])

	res := merge(leftArr, rightArr)

	return res
}

func merge(leftArr, rightArr []int) []int {
	var mergeRes []int
	leftN := len(leftArr)
	rightN := len(rightArr)

	leftIndex :=0
	rightIndex := 0

	for leftIndex < leftN && rightIndex < rightN {
		if leftArr[leftIndex] < rightArr[rightIndex] {
			mergeRes = append(mergeRes, leftArr[leftIndex])
			leftIndex++
		} else {
			mergeRes = append(mergeRes, rightArr[rightIndex])
			rightIndex++
		}
	}

	if leftIndex == leftN {
		mergeRes = append(mergeRes, rightArr[rightIndex:]...)
	}

	if rightIndex == rightN {
		mergeRes = append(mergeRes, leftArr[leftIndex:]...)
	}

	return mergeRes
}

//QuickSort
func QuickSort(arr []int, start, end int)  {
	if start >= end {
		return
	}

	pivot := partition(arr, start, end)
	QuickSort(arr, start, pivot-1)
	QuickSort(arr, pivot+1, end)
}

func partition(arr []int, start, end int) int {
	pivotValue := arr[end]

	i:=start
	for j:=start; j<len(arr); j++ {
		if arr[j] > pivotValue {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}

	arr[end], arr[i] = arr[i], arr[end]

	return i
}

//HeapSort

// 堆排序
func HeapSort(a []int, n int) {
	BuildHeap(a, n)
	k := n
	for k > 1 {
		swap(a, 1, k)
		k--
		heapify(a, k, 1)
	}
}

// 建堆
func BuildHeap(a []int, n int)  {
	//因为对于一个完全二叉树来说，n/2+1到n这个下边内的元素都是叶子结点，所以只需要对1到n/2的非叶子结点进行堆化即可
	for  i:=n/2; i>0; i-- {
		heapify(a, n, i)
	}
	fmt.Println(a)
}

func heapify(a []int, n, i int) {
	for true {
		maxPos := i
		if 2*i < n && a[i] < a[2*1] {
			maxPos = 2*i
		}
		if 2*i+1 < n && a[maxPos] < a[2*i+1] {
			maxPos = 2*i+1
		}
		if maxPos == i {
			break
		}
		swap(a, i, maxPos)
		i = maxPos
	}
}
func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

