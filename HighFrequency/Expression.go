package HighFrequency

//表达式求职
func SolveExpression( s string ) int {
	// write code here

	//我这里定义一个map，用来存储运算符的优先级
	mapOperator := map[byte]int{
		'+': 1,
		'-': 1,
		'*': 2,
		'/': 2,
	}

	n := len(s)
	opsStack := []byte{} // 符号栈
	numStack := []int{} // 数字栈
	for i := 0; i < n; i++ {
		ch := s[i]
		if ch == '(' {
			//遇到 ( 直接入符号栈
			opsStack = append(opsStack, ch)
		} else if ch == ')' { // 遇到 ）则取出符合和数字进行计算，直到遇到)
			for len(opsStack) > 0 {
				//取出符号栈的栈顶元素
				operator := opsStack[len(opsStack)-1]

				if operator != '(' {
					//计算
					if len(numStack) < 2 {
						continue
					}
					if len(opsStack) < 1 {
						continue
					}
					firstVal := numStack[len(numStack)-1]
					secondVal := numStack[len(numStack)-2]
					operator := opsStack[len(opsStack)-1]
					res := calculate(firstVal, secondVal, operator)

					opsStack = opsStack[:len(opsStack)-1]
					numStack = numStack[:len(numStack)-2]

					numStack = append(numStack, res)
				} else {
					opsStack = opsStack[:len(opsStack)-1]
					break
				}
			}

		} else {
			if isNumber(ch) {
				u := 0
				j := i
				//因为可能是多位数（例如：12*(3+(12/4)+5)）
				for j < n && isNumber(s[j]) {
					u = u * 10 + int((s[j] - '0'))
					j++
				}
				numStack = append(numStack, u)
			} else {
				// 只要符号栈不是空的，并且栈顶运算符不是（，就可以进行不断的运算
				for len(opsStack) > 0 && popOperator(opsStack, false) != '(' {
					operator := popOperator(opsStack, false)
					if mapOperator[operator] >= mapOperator[ch] { //站顶运算符优先级 > 当前运算符
						firstVal := numStack[len(numStack)-1]
						secondVal := numStack[len(numStack)-2]
						operator := opsStack[len(opsStack)-1]
						res := calculate(firstVal, secondVal, operator)

						opsStack = opsStack[:len(opsStack)-1]
						numStack = numStack[:len(numStack)-2]

						numStack = append(numStack, res)
					} else {
						break
					}
				}
				opsStack = append(opsStack, ch)
			}
		}

	}
	//把剩余的计算完
	if len(opsStack) > 0 && popOperator(opsStack, false) != '(' {
		firstVal := numStack[len(numStack)-1]
		secondVal := numStack[len(numStack)-2]
		operator := opsStack[len(opsStack)-1]
		res := calculate(firstVal, secondVal, operator)

		opsStack = opsStack[:len(opsStack)-1]
		numStack = numStack[:len(numStack)-2]

		numStack = append(numStack, res)
	}

	return numStack[len(numStack)-1]
}

func calculate(firstVal, secondVal int, operator byte) int {

	res := 0
	switch operator {
	case '+':
		res = secondVal + firstVal //计算的时候注意顺序，栈中靠后那个数，在表达式中是在前边的
	case '-':
		res = secondVal - firstVal
	case '*':
		res = secondVal * firstVal
	case '/':
		res = secondVal / firstVal
	}

	return res
}

func isNumber(ch byte) bool {
	if ch > 47 && ch < 58 {
		return true
	}

	return false
}

func popOperator(stack []byte, flag bool) byte {
	if len(stack) == 0 {
		return 0
	}
	operator := stack[len(stack)-1]
	if flag {
		stack = stack[:len(stack)-1]
	}

	return operator
}