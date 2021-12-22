package BinaryTree

import "math"

//solution-one: stack + iteration
func VerifyPreorder(arr []int) bool {
	min := math.MinInt64
	stack := []int{}
	for _, v := range arr {
		topVal := math.MaxInt64
		if v < min {
			return false
		}
		if len(stack) != 0 {
			topVal = stack[len(stack)-1]
		}
		for len(stack) != 0 && v > topVal{
			min = topVal
			stack = stack[:len(stack)-1]
			topVal = stack[len(stack)-1]
		}
		stack = append(stack, v)
	}
	return true
}

//solution-two: iteration
func VerifyPreorder1(arr []int) bool {
	min := math.MinInt64
	i := -1
	for _, v := range arr {
		if v < min {
			return false
		}
		for i >= 0 && v > arr[i] {
			min = arr[i]
			i--
		}
		i++
		arr[i] = v
	}
	return true
}
