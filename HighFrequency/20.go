package HighFrequency

func IsValid(s string) bool {
	n := len(s)
	if n % 2 != 0 {
		return false
	}
	mapBrackets := map[byte]byte{
		')':'(',
		']':'[',
		'}':'{',
	}
	stack := []byte{}
	for i:=0; i < n; i++ {
		if _, ok := mapBrackets[s[i]]; ok {
			if len(stack) == 0 || stack[len(stack)-1] != mapBrackets[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}

	if len(stack) == 0 {
		return true
	}

	return false
}