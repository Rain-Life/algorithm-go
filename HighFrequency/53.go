package HighFrequency

//solution-one: force
func MaxSubArray2(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n==1 {
		return nums[0]
	}
	max := nums[0]
	for i := 0; i < n-1; i++ {
		sum := 0
		for j := i; j < n; j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}

//solution-two: dynamic programming
// 连续子数组的最大和
func MaxSubArray(nums []int) int {
	n := len(nums)
	if n == 0 {
		return -1
	}
	if n==1 {
		return nums[0]
	}

	maxSum := nums[0]
	for i:=1; i < n; i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1] //注意，这就将num[i]位置的值，保存成了0~i这个子序列的最大和(相当于记录了0~i这段连续子数组的状态（最大和）)
		}
		if nums[i] > maxSum {
			maxSum = nums[i]
		}
	}

	return maxSum
}