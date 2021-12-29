
## 大数加法

**题目来源**：[牛客网 - **NC1 大数加法**](https://www.nowcoder.com/practice/11ae12e8c6fe48f883cad618c2e81475?tpId=117&&tqId=37842&rp=1&ru=/activity/oj&qru=/ta/job-code-high/question-ranking)

### 题目描述

以字符串的形式读入两个数字，编写一个函数计算它们的和，以字符串形式返回

**数据范围**：len(s),len(t) ≤ 100000，字符串仅由'0'~‘9’构成

### 示例

**示例 1**

```go
输入："1","99"
返回值："100"
说明：1+99=100
```

**示例 2**

```go
输入："114514",""
返回值："114514"
```

## 解题

### 思路

这个比较简单，只要注意一下如何将字符类型转成整形（针对Go语言）（字符和数字正好差48）。其余的地方，跟通常计算两数相加是一样的

### 代码

```go
func solve( s string ,  t string ) string {
	sl, tl := len(s), len(t)
	if sl == 0 {
		return t
	}
	if tl == 0 {
		return s
	}

	if sl < tl { //让s为那个最长的串
		tl, sl = sl, tl
		t, s = s, t
	}

	resStr := make([]byte, sl+1) //因为可能存在进位，所以+1

	carry := 0 //进位
	for i:=0; i < sl; i++ {
		schInt := int(s[sl-i-1]-'0') // 减0是因为字符和数字正好差48，而“0” byte值是 48
		tchInt := 0
		if tl - i - 1 >= 0 {
			tchInt = int(t[tl-i-1]-'0')
		}
		sum := schInt + tchInt +carry
		carry = sum / 10
		sum = sum % 10
		resStr[sl-i] = byte(sum + '0')
	}

	if carry == 1 {
		resStr[0] = '1'
		return string(resStr)
	}

	return string(resStr[1:])
}
```