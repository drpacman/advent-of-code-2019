package numbers

import (
	"math"
)

func Digits(value int) []int {
	var digits = make([]int, 0)
	for value >= 10 {
		r := value % 10
		digits = append([]int{r}, digits...)
		value = int(math.Floor(float64(value / 10)))
	}
	digits = append([]int{value}, digits...)
	return digits
}
