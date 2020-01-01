package numbers

import (
	"testing"
)

func TestCombinations(t *testing.T) {
	result := Combinations([]int{0, 1, 2, 3, 4})
	if len(result) != 120 {
		t.Errorf("Incorrect number of combinations - got %v expected %v", len(result), 120)
	}
}
