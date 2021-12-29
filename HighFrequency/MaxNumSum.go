package HighFrequency

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