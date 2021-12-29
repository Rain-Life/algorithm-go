package HighFrequency
//跳台阶
var depth int
var hashMap = map[int]int{}
func Step(n int) int {
	if depth > 1000 {
		return -1
	}
	if n <= 1 {
		return 1
	}
	if n==2 {
		return 2
	}

	if hashMap[n] != 0 {
		return hashMap[n]
	}

	res := Step(n-1) + Step(n-2)
	hashMap[n] = res

	return res
}