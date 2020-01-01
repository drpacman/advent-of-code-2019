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

func Combinations(combo []int) [][]int {
	if len(combo) == 1 {
		return [][]int{[]int{combo[0]}}
	} else {
		result := make([][]int, 0)
		for i := 0; i < len(combo); i++ {
			remainder := make([]int, 0)
			remainder = append(remainder, combo[:i]...)
			remainder = append(remainder, combo[i+1:]...)
			var tail [][]int = Combinations(remainder)
			for j := 0; j < len(tail); j++ {
				result = append(result, append([]int{combo[i]}, tail[j]...))
			}
		}
		return result
	}
}
