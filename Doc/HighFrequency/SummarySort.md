
# 九大排序算法总结

## 冒泡排序

### 核心思想

每次冒泡操作都会对**相邻的两个元素**进行比较，看是否满足大小关系要求。如果不满足就让它俩互换。**一次冒泡会让至少一个元素移动到它应该在的位置**，重复n次，就完成了n个数据的排序工作

### 排序过程

冒泡过程还可以优化。当某次冒泡操作已经没有数据交换时，说明已经达到完全有序，不用再继续执行后续的冒泡操作。如图

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/1.png)

### 代码实现

```go
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
```

### 冒泡排序算法分析

首先**冒泡排序是一个原地排序算法**，因为冒泡排序只涉及相邻数据的交换，需要常量级的临时空间，所以空间复杂度是O(1)

在冒泡排序中，只有交换才可以改变两个元素的前后顺序。为了保证冒泡排序算法的稳定性，当有相邻的两个元素大小相等的时候，我们不做交换，相同大小的数据在排序前后不会改变顺序，所以**冒泡排序是稳定的排序算法**

在最好的情况下，也就是待排数据是完全有序的，那只需要进行一次冒泡操作即可，所以**最好情况下的时间复杂度是O(n)**

最坏情况下是待排数据是完全无序的，这个时候就需要n次冒泡，所以**最坏情况下的时间复杂度是O(n^2)**

## 插入排序

### 核心思想

将待排序的数组分成两个区间，有序区和无序区。刚开始的时候，有序区只有第一个元素。插入排序的过程就是每次从无序区中取出一个元素，放入到有序区中对应的位置，保证插入到有序区中之后，有序区依然是有序的。不断的重复这个过程，直到无序区为空

### 排序过程

i指向待排序区的元素，j指向已排序区的元素。i负责遍历无序区中的每一个元素，j负责找到无序区中元素在有序区中正确的位置（这里只画了一轮，后边的每一轮都是同样的过程）

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/2.png)

### 代码实现

```go
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
```

### 插入排序算法分析

插入排序也不需要额外的存储空间，空间复杂度是O(1)，所以它是**原地排序算法**

在插入排序中，对于值相同的元素，我们可以选择将后面出现的元素，插入到前面出现元素的后面，这样就可以保持原有的前后顺序不变，所以插入排序是**稳定的排序算法**

如果待排序的数据是完全有序的，并不需要搬移任何数据。如果从尾到头在有序数据组里面查找插入位置，每次只需要比较一个数据就能确定插入的位置。所以这种情况下，最好是时间复杂度为O(n)。注意，**这里是从尾到头遍历已经有序的数据**

如果数组是倒序的，每次插入都相当于在数组的第一个位置插入新的数据，所以需要移动大量的数据，所以**最坏情况时间复杂度为O(n^2)**。**平均时间复杂度也是O(n^2)**

## 选择排序

### 核心思想

选择排序的思想和插入排序的思想有些类似，选择排序是每次从无序区中选择一个最小的元素放入到有序区中

### 排序过程

在无序区中找最小元素的过程就是，取无序区中的第一个元素，和无序区中的每一个进行比较。具体如图：

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/3.png)

### 代码实现

```go
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
```

### 选择排序算法分析

选择排序的空间复杂度也是O(1)，是**原地排序算法**。选择排序的最好情况和最坏情况的时间复杂度都是O(n^2)，这个很简单，看一下它的执行过程就知道了

**选择排序不是一个稳定排序**，选择排序每次都要找剩余未排序元素中的最小值，并和前面的元素交换位置，这样破坏了稳定性

比如7，3，5，7，1，9 这样一组数据，使用选择排序算法来排序的话，第一次找到最小元素1，与第一个7交换位置，那第一个7和中间的7顺序就变了，所以就不稳定了

## 归并排序

### 核心思想

如果要排序一个数组，我们先把数组从中间分成前后两部分，然后对前后两部分分别排序，再将排好序的两部分合并在一起，这样整个数组就都有序了

归并排序使用的就是**分治思想**。分治，顾名思义，就是分而治之，将一个大问题分解成小的子问题来解决。小的子问题解决了，大问题也就解决了

### 排序过程

每次取数组的中间元素为分区点，将数组分成前后两个部分，当每部分数组元素个数只剩1个之后，开始将拆分的数组进行合并，合并的过程中，保证其有序

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/4.png)

关于有序数组的合并非常简单，这里不画图展示

### 代码实现

```go
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
```

### 归并排序算法分析

首先**「归并排序是稳定排序算法」**，这个主要取决于merge操作，也就是将两个有序数组合并成一个有序数组。在合并的过程中，当两个数组中有相同的元素时，先把前边那部分数组中元素放入到tmp中，这样就可以保证相同的元素，在合并前后顺序不变

从归并排序的原理图中可以看出来，**「归并排序的执行效率与要排序的原始数组的有序程度无关，所以其时间复杂度是非常稳定的，不管是最好情况、最坏情况，还是平均情况，时间复杂度都是O(nlogn)」**

归并排序和下边的快速排序相比，虽然时间复杂度都是O(nlogn)，归并排序应用却不是那么广泛，因为它**「不是原地排序」**。归并排序中合并过程需要借助额外的存储空间（空间复杂度是O(n)）

## 快速排序

### 核心思想

如果要排序数组中下标从 p 到 r 之间的一组数据，我们选择 p 到 r 之间的任意一个数据作为 pivot（分区点）

遍历 p 到 r 之间的数据，将小于 pivot 的放到左边，将大于 pivot 的放到右边，将 pivot 放到中间。经过这一步骤之后，数组 p 到 r 之间的数据就被分成了三个部分，前面 p 到 q-1 之间都是小于 pivot 的，中间是 pivot，后面的 q+1 到 r 之间是大于 pivot 的

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/5.png)

根据分治、递归的处理思想，我们可以用递归排序下标从 p 到 q-1 之间的数据和下标从 q+1 到 r 之间的数据，直到区间缩小为 1，就说明所有的数据都有序了

### 排序过程

在归并排序里边有一个merge()方法，是将两个有序的数组合并成一个有序的数组。而这里用到了一个partition()函数，就是上边说到的，随机选择一个元素作为分区点（pivot）（一般情况下，可以选择start到end区间的最后一个元素），然后对A[start…end]分区，函数返回分区点（pivot）的下标

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/6.png)

partition()函数实现的过程

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/7.png)

### 代码实现

```go
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
```

### 快速排序算法分析

上边使用了一种巧妙的方法，在空间复杂度为O(1)的情况下，实现了快速排序，所以**「快排是一个稳定的排序算法」**

因为分区的过程涉及交换操作，如果数组中有两个相同的元素，比如序列 6，8，7，6，3，5，9，4，在经过第一次分区操作之后，两个6的相对先后顺序就会改变（跟着上边的分区算法图，很容易可以推出来）。所以，**「快速排序并不是一个稳定的排序算法」**

## 堆排序

要了解堆排序，需要先了解堆这个数据结构，包括建堆、从堆中删除一个元素、往堆中插入一个元素

### **什么是堆**

- 堆是一个完全二叉树；
- 堆中每一个节点的值都必须大于等于（或小于等于）其子树中每个节点的值

对于每个节点的值都大于等于子树中每个节点值的堆，我们叫做“大顶堆”。对于每个节点的值都小于等于子树中每个节点值的堆，我们叫做“小顶堆”

### 如何实现一个堆

**完全二叉树比较适合用数组来存储**。用数组来存储完全二叉树是非常节省存储空间的。因为我们不需要存储左右子节点的指针，单纯地通过数组的下标，就可以找到一个节点的左右子节点和父节点

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/8.png)

从图中可以看到，数组中下标为 i 的节点的左子节点，就是下标为 i∗2 的节点，右子节点就是下标为 i∗2+1 的节点，父节点就是下标为 i/2 的节点

### 向堆中插入元素

向堆中插入一个元素之后，需要保证继续满足堆的两个特性

如果把新插入的元素**放到堆的最后**，可以看下边这个图，是不是不符合堆的特性了？于是，就需要进行调整，让其重新满足堆的特性，这个过程就叫做**堆化**（heapify）

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/9.png)

堆化非常简单，就是顺着节点所在的路径，向上或者向下，对比，然后交换

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/10.png)

```go
package Heap

import "fmt"

var a []int //数组
var count int //堆中已存储的数据数量
var n int //堆的容量

func Heap(capacity int) {
	a = make([]int, capacity)
	n = capacity
	count = 0
}

// 建堆
func BuildHeap(a []int, n int)  {
	//因为对于一个完全二叉树来说，n/2+1到n这个下边内的元素都是叶子结点，所以只需要对1到n/2的非叶子结点进行堆化即可
	for  i:=n/2; i>0; i-- {
		heapify(a, n, i)
	}
	fmt.Println(a)
}

func swap(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

//堆化（插入语数据）
func InsertHeap(data int) {
	if count == n {
		return //堆满了
	}
	count++
	a[count] = data
	i := count
	for i/2 > 0 && a[i]>a[i/2] {
		swap(a, i, i/2)
		i = i/2
	}
	fmt.Println("插入后", a)
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
```

### 删除堆顶元素

把最后一个节点放到堆顶（覆盖掉堆顶元素，也就是删除），然后利用同样的父子节点对比方法。对于不满足父子节点大小关系的，互换两个节点，并且重复进行这个过程，直到父子节点之间满足大小关系为止

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/11.png)

```go
func RemoveHeap()  {
	if count == 0 {
		return //堆是空的
	}
	a[1] = a[count]
	count--
	heapify(a, count, 1)//最后一个参数表示要堆化哪个元素
	fmt.Println(a)
}
```

一个包含 n 个节点的完全二叉树，树的高度不会超过 log2n。堆化的过程是顺着节点所在路径比较交换的，所以堆化的时间复杂度跟树的高度成正比，也就是 O(logn)。插入数据和删除堆顶元素的主要逻辑就是堆化，所以，往堆中插入一个元素和删除堆顶元素的时间复杂度都是 O(logn)

### 堆排序

**建堆**

因为叶子节点往下堆化只能自己跟自己比较，所以我们直接从最后一个非叶子节点开始，依次堆化就行了

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/12.png)

```go
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
```

建堆的时间复杂度是 O(n)

**排序**

建堆结束之后，数组中的数据已经是按照大顶堆的特性来组织的。数组中的第一个元素就是堆顶，也就是最大的元素。把它跟最后一个元素交换，那最大元素就放到了下标为 n 的位置

这个过程有点类似上面讲的“删除堆顶元素”的操作，**当堆顶元素移除之后，把下标为 n 的元素放到堆顶，然后再通过堆化的方法，将剩下的 n−1 个元素重新构建成堆**。堆化完成之后，我们再取堆顶的元素，放到下标是 n−1 的位置，一直重复这个过程，直到最后堆中只剩下标为 1 的一个元素，排序工作就完成了

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/13.png)

```go
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
```

### **堆排序算法分析**

整个堆排序的过程，都只需要极个别临时存储空间，所以堆排序是**原地排序算法**。堆排序包括建堆和排序两个操作，建堆过程的时间复杂度是 O(n)，排序过程的时间复杂度是 O(nlogn)，所以，堆排序整体的时间复杂度是 **O(nlogn)**

堆排序**不是稳定的排序算法**，因为在排序的过程，存在将堆的最后一个节点跟堆顶节点互换的操作，所以就有可能改变值相同数据的原始相对顺序

## 桶排序

### 核心思想

将要排序的数据分到几个有序的桶里，每个桶里的数据再单独进行排序。桶内排完序之后，再把每个桶里的数据按照顺序依次取出，组成的序列就是有序的了

### 排序过程

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/14.png)

### 桶排序算法分析

假设需要排序的数据个数是n，将这n个数据均匀的划分到m个桶中，每个桶里就有 k=n/m 个元素。每个桶内部使用快速排序，时间复杂度为 O(k * logk)。m 个桶排序的时间复杂度就是 O(m * k * logk)，因为 k=n/m，所以整个桶排序的时间复杂度就是 O(n*log(n/m))。当桶的个数 m 接近数据个数 n 时，log(n/m) 就是一个非常小的常量，这个时候桶排序的时间复杂度接近 **O(n)**

虽然桶排序算法的时间复杂度比前边的六大排序算法低，但是并不能取代前边的那六大排序算法，因为**桶排序对要排序的数据要求很苛刻**

1. 要排序的数据需要很容易就能划分成 m 个桶，并且，**桶与桶之间有着天然的大小顺序**。这样每个桶内的数据都排序完之后，桶与桶之间的数据不需要再进行排序
2. 数据在各个桶之间的分布是比较均匀的（如果数据在桶之间分布不均匀，时间复杂度就可能退化到nlogn，比如所有数据划分之后，都在一个桶中）

**适用的场景**：数据的范围不大

## 计数排序

### 核心思想

**计数排序其实是桶排序的一种特殊情况**。当要排序的 n 个数据，所处的范围并不大的时候，比如最大值是 k，我们就可以把数据划分成 k 个桶。每个桶内的数据值都是相同的，省掉了桶内排序的时间

比如说高考，假设有 50 万考生，考生的满分是 900 分，最小是 0 分，这个数据的范围很小，所以我们可以分成 901 个桶，对应分数从 0 分到 900 分

根据考生的成绩，将这 50 万考生划分到这 901 个桶里。**桶内的数据都是分数相同的考生**，所以并不需要再进行排序。只需要依次扫描每个桶，将桶内的考生依次输出到一个数组中，就实现了 50 万考生的排序。因为只涉及扫描遍历操作，所以时间复杂度是 **O(n)**

### 排序过程

该排序算法叫计数排序，计数从哪里体现的？可以看下边这个例子

假设只有 8 个考生，分数在 0 到 5 分之间。这 8 个考生的成绩我们放在一个数组 A[8]中，它们分别是：2，5，3，0，2，3，0，3

考生的成绩从 0 到 5 分，使用大小为 6 的数组 C[6]表示桶，其中**下标对应分数**。不过，C[6]内存储的并不是考生，而是对应的考生个数。只需要遍历一遍考生分数，就可以得到 C[6]的值

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/15.png)

从图中可以看出，分数为 3 分的考生有 3 个，小于 3 分的考生有 4 个，所以，成绩为 3 分的考生在排序之后的有序数组 R[8]中，会保存下标 4，5，6 的位置

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/16.png)

现在就是要找出相同分数的考生，应该在什么位置。思路就是，对C[6]数组顺序求和，C[6]存储的数据就变成了下面这样子。**C[k]里存储小于等于分数 k 的考生个数**

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/17.png)

现在就从后到前依次扫描数组 A（A中保存的是乱序的学生成绩）。比如，当扫描到 3 时，可以从数组 C 中取出下标为 3 的值 7（C中下标表示的是分数），也就是说，到目前为止，包括自己在内，分数小于等于 3 的考生有 7 个，也就是说 3 是数组 R 中的第 7 个元素（也就是数组 R 中下标为 6 的位置）。当 3 放入到数组 R 中后，小于等于 3 的元素就只剩下了 6 个了，所以相应的 C[3]要减 1，变成 6

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/18.png)

以此类推，当扫描到第 2 个分数为 3 的考生的时候，就会把它放入数组 R 中的第 6 个元素的位置（也就是下标为 5 的位置）。当扫描完整个数组 A 后，数组 R 内的数据就是按照分数从小到大有序排列的了

### 计数排序算法分析

计数排序**只能用在数据范围不大的场景中**，如果数据范围 k 比要排序的数据 n 大很多，就不适合用计数排序了。而且，计数排序**只能给非负整数排序**，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数

## 基数排序

### 核心思想

将待排序数据的每一个数据的每个小的单元逐一进行比较排序，当每个单元排好序之后，这组数据就是有序的了

比如有10w个手机号，希望将这 10 万个手机号码从小到大排序。手机号码有 11 位，范围太大，不适合用前边提到的排序算法。假设要比较两个手机号码 a，b 的大小，如果在前面几位中，a 手机号码已经比 b 手机号码大了，那后面的几位就不用看了

先按照最后一位来排序手机号码，然后，再按照倒数第二位重新排序，以此类推，最后按照第一位重新排序。经过 11 次排序之后，手机号码就都有序了

### 排序过程

下边以字符串排序为例来画图

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/SummarySort/19.png)

**注意，这里按照每位来排序的排序算法要是稳定的，否则这个实现思路就是不正确的**。因为如果是非稳定排序算法，那最后一次排序只会考虑最高位的大小顺序，完全不管其他位的大小关系，那么低位的排序就完全没有意义了

### 基数排序算法分析

根据每一位来排序，我们可以用上边提到的桶排序或者计数排序，它们的时间复杂度可以做到 

O(n)。如果要排序的数据有 k 位，那我们就需要 k 次桶排序或者计数排序，总的时间复杂度是 

O(k*n)。当 k 不大的时候，比如手机号码排序的例子，k 最大就是 11，所以基数排序的时间复杂度就

近似于 O(n)

基数排序对要排序的数据是有要求的，**需要可以分割出独立的“位”来比较**，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。除此之外，**每一位的数据范围不能太大**，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了

## 排序算法对比

| Name | 退化条件 | 适用场景 | 最好时间复杂度 | 最坏时间复杂度 | 平均时间复杂度 | 空间复杂度 | 是否是稳定排序 | 是否是原地排序 |
| --- | --- | --- | --- | --- | --- | --- | --- | --- |
| 冒泡排序 | 排序的数据刚好是倒序的 | - | O(n)（排序的数据已经是有序的情况） | O(n^2) | O(n^2) | O(1) | 是稳定排序 | 原地排序(空间复杂度是O(1)的就是原地排序) |
| 插入排序 | 排序的数据刚好是倒序的 | - | O(n)（排序的数据已经是有序的情况） | O(n^2) | O(n^2) | O(1) | 是稳定排序 | 原地排序 |
| 选择排序 | - | - | O(n^2) | O(n^2) | O(n^2) | O(1) | 不是稳定排序 | 原地排序 |
| 归并排序 | - | - | O(nlogn) | O(nlogn) | O(nlogn) | O(n) | 是稳定排序 | 不是原地排序 |
| 快速排序 | 排序的数据刚好是有序的 | - | O(nlogn) | O(n^2) | O(nlogn) | O(1) | 不是稳定排序 | 原地排序 |
| 堆排序 | - | - | - | - | O(nlogn) | O(1) | 不是稳定排序 | 原地排序 |
| 桶排序 | 在极端情况下，如果数据都被划分到一个桶里 | 1. 数据量比较大，内存有限，无法将数据全部加载到内存中| O(n) | O(nlogn) | O(n) | - | - | - |
| 计数排序 | - | 1. 计数排序只能用在数据范围不大的场景中，如果数据范围 k 比要排序的数据 n 大很多，就不适合用计数排序了
2. 计数排序只能给非负整数排序，如果要排序的数据是其他类型的，要将其在不改变相对大小的情况下，转化为非负整数 | O(n) | O(nlogn) | O(n) | - | - | - |
| 基数排序 | 基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了 | 基数排序对要排序的数据是有要求的，需要可以分割出独立的“位”来比较，而且位之间有递进的关系，如果 a 数据的高位比 b 数据大，那剩下的低位就不用比较了。除此之外，每一位的数据范围不能太大，要可以用线性排序算法来排序，否则，基数排序的时间复杂度就无法做到 O(n) 了 | O(n) | - | O(n) | - | - | - |