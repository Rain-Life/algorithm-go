package HighFrequency

//solution-one: recursion
//斐波那契数列
var repeatNum = map[int]int{}
func Fib(n int) int {
	if n ==0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	num, ok := repeatNum[n]
	if ok {
		return num
	}

	res := Fib(n-1) + Fib(n-2)
	repeatNum[n] = res
	return res
}


// solution-two: dynamic programming
func fib(n int) int {
	if (n == 0) {
		return 0
	}
	if (n == 1) {
		return 1
	}
	status := make([]int, n+1)

	status[0] = 0
	status[1] = 1
	for j:=2; j <= n; j++ {
		status[j] = status[j-1] + status[j-2]
	}

	return status[n]
}