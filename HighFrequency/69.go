package HighFrequency

func mySqrt(x int) int {
	if x == 1 || x == 0 {
		return x
	}
	if x <= 0 {
		return -1
	}

	left,  right := 0,  x
	res := -1
	for left <= right {
		mid := left + (right-left)/2
		if mid * mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return res
}