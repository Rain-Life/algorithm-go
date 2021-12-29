package HighFrequency

//solution-one: slide window
func SlideWindow(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	left, right, window, maxLength := 0, 0, make(map[byte]int), 1
	for right < n {
		var rightChar = s[right]
		index, ok := window[s[right]] //是否出现过
		if ok && index > left {
			left = index
		}

		if right - left + 1 > maxLength {
			maxLength = right - left + 1
		}

		window[rightChar] = right + 1
		right++
	}

	return maxLength
}

// solution-two: force
func lengthOfLongestSubstring(s string) int {
	lastOccured := make(map[byte]int) //用来记录已经遍历过的字符（下标是s的值，它的值是s的下标）
	maxLength := 0
	start :=0 //记录最长不含重复子串的起始位置

	for i, ch := range []byte(s) {
		lastId, ok := lastOccured[ch]
		if ok && lastId >= start { //如果该字符出现过，并且在start的后边出现的，更新start的位置
			start = lastId + 1
		}
		if i - start + 1 > maxLength {
			maxLength = i - start + 1
		}

		lastOccured[ch] = i
	}

	return maxLength
}