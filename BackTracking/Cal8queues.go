package BackTracking

import (
	"fmt"
	"math"
)

var result [8]int
var count = 1
func Cal8queues(row int) {
	if row == 8 {
		printQueues(result)
		return
	}

	for column:=0; column < 8; column++ {
		if isOk(row, column) {
			result[row] = column
			Cal8queues(row+1)
		}
	}
}

func isOk(row, column int) bool {
	for i := row-1; i>=0; i-- {
		if result[i] == column {
			return false
		}
		if math.Abs(float64(row - i)) == math.Abs(float64(column - result[i])) {
			return false
		}
		if (row + column) == (i + result[i]) {
			return false
		}
	}

	return true
}

func printQueues(result [8]int) {
	fmt.Printf("第%v种解法\n", count)
	for row :=0; row < 8; row++ {
		for column:=0; column < 8; column++ {
			if result[row] == column {
				fmt.Printf("Q\t")
			} else {
				fmt.Printf("*\t")
			}
		}
		fmt.Println()
	}
	fmt.Println()
	count++
}
