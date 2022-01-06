package BackTracking

var MaxW = 1
func Bag(i, cw int, items []int, n, w int) {
	if cw == w || i == n {
		if cw > MaxW {
			MaxW = cw
		}
		return
	}
	
	Bag(i+1, cw, items, n, w)
	if cw + items[i] <= w {
		Bag(i+1, cw+items[i], items, n, w)
	}
}
