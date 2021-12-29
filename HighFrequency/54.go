package HighFrequency

//螺旋矩阵
func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	left, right, top, bottom := 0, len(matrix[0])-1, 0, len(matrix)-1
	res := []int{}
	for  {
		//从左往右，将遍历的数字依次放入结果集res中(可以发现过程中行是不变的，也就是top)
		for i := left; i <= right; i++ {
			res = append(res, matrix[top][i])
		}
		//下一次循环（这个循环指的是最外层的大循环）的时候，就从下一行开始，所以top要+1
		top++
		if top > bottom {
			break
		}

		//从上往下，将遍历的数字依次放入结果集res中(可以发现过程中列是不变的，也就是right)
		for i := top; i <= bottom; i++ {
			res = append(res, matrix[i][right])
		}
		//下一次循环的时候，就开始从倒数第二列开始遍历了，所以right要-1
		right--
		if left > right {
			break
		}

		//从右往左，将遍历的数字依次放入结果集res中(可以发现过程中行是不变的，也就是bottom)
		for i := right; i >= left; i-- {
			res = append(res, matrix[bottom][i])
		}
		//下一次循环的时候，就开始从倒数第二行开始遍历了，所以bottom要-1
		bottom--
		if top > bottom {
			break
		}

		//从下往上
		for i := bottom; i >= top; i-- {
			res = append(res, matrix[i][left])
		}
		left++
		if left > right {
			break
		}
	}

	return res
}