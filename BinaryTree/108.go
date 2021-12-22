package BinaryTree

//solution-one: recursion
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return convert(nums, 0, len(nums)-1)
}

func convert(nums []int, left, right int) *TreeNode {
	if(left > right) {
		return nil
	}
	mid := (left+right) / 2
	root := &TreeNode{Val: nums[mid]}
	root.Left = convert(nums, left, mid-1)
	root.Right = convert(nums, mid + 1, right)

	return root
}
