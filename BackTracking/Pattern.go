package BackTracking

var pattern = "ab*f?j"
var plen = len(pattern)

var matched bool = false

func Match(text string, tlen int)  {
	rmatch(0, 0, tlen, text)
}

func rmatch(ti, pi, tlen int, text string) {
	if matched {
		return
	}
	if plen == pi {
		if tlen == ti {
			matched = true
		}
		return
	}

	if pattern[pi] == '*' {
		for i:=0; i < tlen-ti; i++ {
			rmatch(ti+i, pi+1, tlen, text)
		}
	} else if pattern[pi] == '?' {
		rmatch(ti, pi+1, tlen, text)
		rmatch(ti+1, pi+1, tlen, text)
	} else if ti < tlen && pattern[pi] == text[ti] {
		rmatch(ti+1, pi+1, tlen, text)
	}
}
