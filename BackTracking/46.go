package BackTracking


var ResPermute [][]int
func Permute(nums []int) [][]int {
	var path = []int{}
	BackTrackingPermute(nums, path)
	return ResPermute
}

func BackTrackingPermute(nums []int, path []int) {
	if len(path) == len(nums) {
		path1 := make([]int, len(nums))
		copy(path1, path)
		ResPermute = append(ResPermute, path1)
		return
	}
	for i := 0; i < len(nums); i++ {
		if isExist(path, nums[i]) {
			continue
		}
		path = append(path, nums[i])
		BackTrackingPermute(nums, path)
		path = path[:len(path)-1]
	}
}

func isExist(nums []int, elem int) bool {
	for _, v := range nums  {
		if v == elem {
			return true
		}
	}

	return false
}
