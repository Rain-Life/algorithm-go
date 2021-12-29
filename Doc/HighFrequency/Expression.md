
## 表达式求值

**题目来源**：[牛客网-表达式求值](https://www.nowcoder.com/practice/c215ba61c8b1443b996351df929dc4d4?tpId=117&&tqId=37849&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking)

### 题目描述

请写一个整数计算器，支持加减乘三种运算和括号

数据范围：0 ≤ |s| ≤ 100，保证计算结果始终在整型范围内

要求：空间复杂度： O(n)，时间复杂度 O(n)

### 示例

**示例 1**

```go
输入："1+2"
返回值：3
```

**示例 2**

```go
输入："(2*(3-4))*5"
返回值：-10
```

**示例 3**

```go
输入："3+2*3*4-1"
返回值：26
```

## 解题

### **思路**

**常规的表达式求值**

一个表达式包含两个部分，数字和运算符。可以用两个栈来实现表达式求值，一个栈用来存储数字，一个栈用来存储运算符

从左向右遍历表达式，当遇到数字时，将数字放入到存储数字的栈；如果遇到运算符，将存储运算符栈的栈顶元素取出，进行优先级比较

**如果比运算符栈顶元素优先级高，则将当前运算符压入到存储运算符的栈中；如果比运算符栈顶元素低或优先级一样，则从存储数字的栈中取出两个元素，然后进行计算，将计算的结果放入到存储数字的栈中**。重复上边的操作

**存在小括号的表达式求职**

- 如果遇到(，直接将其压入到存储运算符的栈中
- 如果遇到)，则取出符号栈中的运算符和数字栈中的数字进行计算，将计算结果放入数字栈，直到遇到(
- 如果遇到运算符，发现栈顶元素是(，则运算符入符号栈

以 ( 5 * ( 8 - 3 * 2 ) + 5 ) * 2 表达式为例 例，按照上边的思路，图解它的计算过程

![image](https://github.com/Rain-Life/algorithm-go/blob/master/photos/HighFrequency/Expression/1.png)

代码实现起来还是有些困难的，建议自己敲一遍，需要考虑的细节挺多的。代码敲了俩小时才实现（**做算法题还是不能求快，搞懂最重要，不要感觉自己脑子里边有思路就感觉自己会了，共勉！**）

### **代码**

```go
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
```