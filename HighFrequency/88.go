package HighFrequency

//solution-one: two pointers
func Merge2(nums1 []int, m int, nums2 []int, n int)  {
	mi, ni := 0, 0
	mergeArr := make([]int, 0, m+n)

	for {
		if mi == m {
			mergeArr = append(mergeArr, nums2[ni:]...)
			break
		}
		if ni == n {
			mergeArr = append(mergeArr, nums1[mi:]...)
			break
		}
		if nums1[mi] < nums2[ni] {
			mergeArr = append(mergeArr, nums1[mi])
			mi++
		} else {
			mergeArr = append(mergeArr, nums2[ni])
			ni++
		}
	}

	copy(nums1, mergeArr)
}