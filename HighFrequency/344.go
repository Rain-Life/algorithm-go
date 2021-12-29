package HighFrequency

//solution-two: two pointers
//反转字符串
func reverseString(s []byte)  {
	left, right := 0, len(s)-1
	for ; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}