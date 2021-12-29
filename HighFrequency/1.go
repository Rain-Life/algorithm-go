package HighFrequency
func TwoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	hashMap := make(map[int]int, len(nums))

	for j := 0; j < len(nums); j++ {
		if v, ok := hashMap[target-nums[j]]; ok {
			return []int{j, v}
		}
		hashMap[nums[j]] = j
	}

	return nil
}