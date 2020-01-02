package numbers

import (
	"math"
)

func Digits(value int64) []int {
	var digits = make([]int, 0)
	for value >= 10 {
		r := int(value % 10)
		digits = append([]int{r}, digits...)
		value = int64(math.Floor(float64(value / 10)))
	}
	digits = append([]int{int(value)}, digits...)
	return digits
}

func Combinations(combo []int64) [][]int64 {
	if len(combo) == 1 {
		return [][]int64{[]int64{combo[0]}}
	} else {
		result := make([][]int64, 0)
		for i := 0; i < len(combo); i++ {
			remainder := make([]int64, 0)
			remainder = append(remainder, combo[:i]...)
			remainder = append(remainder, combo[i+1:]...)
			var tail [][]int64 = Combinations(remainder)
			for j := 0; j < len(tail); j++ {
				result = append(result, append([]int64{combo[i]}, tail[j]...))
			}
		}
		return result
	}
}
