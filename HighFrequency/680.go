package HighFrequency

//solution-one: force
func ValidPalindrome2(s string) bool {
	if IsPalindrome(s) {
		return true
	}

	sNew := ""
	for i := 0; i < len(s); i++ {
		if i == 0 {
			sNew = s[1:]
		} else if i == len(s)-1 {
			sNew = s[:len(s)-1]
		} else {
			sNew = s[0:i] + s[i+1:]
		}
		if IsPalindrome(sNew) {
			return true
		}
	}

	return false
}

//判断某一段字符串是否是回文字符串
func IsPalindrome(s string) bool {

	left, right := 0, len(s)-1
	for left <= right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

//solution-two: two pointers
func ValidPalindrome(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}

	left, right := 0, len(s)-1
	for left <= right {
		if s[left] == s[right] {
			left++
			right--
		} else {
			//判断删除左边元素后的字符串是否是回文
			flagLeft, flagRight := true, true
			i, j := left+1, right
			for i < j {
				if s[i] != s[j] {
					flagLeft = false
					break
				}
				i++
				j--
			}
			//判断删除右边元素后的字符串是否是回文
			i, j = left, right-1
			for i < j {
				if s[i] != s[j] {
					flagRight = false
					break
				}
				i++
				j--
			}

			return flagLeft || flagRight
		}
	}

	return true
}

