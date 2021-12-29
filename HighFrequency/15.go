package HighFrequency

import "sort"

//solution-one: force
func ThreeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	//排序
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}
	for i:=0; i < n; i++  {
		//如果当前元素和上一个元素相同，则跳过（避免重复）
		if i > 0 && nums[i] == nums[i-1]  {
			continue
		}
		for j := i + 1; j < n; j++ {
			//如果当前元素和上一个元素相同，则跳过（避免重复）
			if j > i + 1 && nums[j] == nums[j-1] {
				continue
			}
			for k := j + 1; k < n; k++ {
				//如果当前元素和上一个元素相同，则跳过（避免重复）
				if k > j + 1 && nums[k] == nums[k-1] {
					continue
				}
				if (nums[i] + nums[j] + nums[k]) == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
			}
		}
	}

	return res
}

//solution-two: sort + two pointers
func ThreeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}

	//排序
	sort.Ints(nums)
	n := len(nums)
	res := [][]int{}

	for first := 0; first < n; first ++ {
		//如果和上一个数字相等，跳过（避免重复）
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n-1 // 指向最后一个元素（利用双指针）
		target := -1 * nums[first] //我们要找的第二个和第三个数的和，一定是和第一个数互为相反数的，因此这里取反，target = nums[first] + nums[second]

		//second就是指向数组开头的那个指针
		for second := first + 1; second < n; second++ {
			if second > first +1 && nums[second] == nums[second-1] {
				continue
			}
			//如果第二个和第三个数字的和大于目标值，后边那个指针向前移动（因为现在数字是有序的）
			for second < third && nums[second] + nums[third] > target {
				third--
			}
			if second == third { //说明元素遍历完了
				break
			}
			if nums[second] + nums[third] == target {
				res = append(res, []int{nums[first], nums[second], nums[third]})
			}
		}
	}

	return res
}