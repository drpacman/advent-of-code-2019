package main

import (
	"fmt"
	"math"
)

func digits(value int) []int {
	var digits = make([]int, 6)
	for i := 6; i > 0; i-- {
		n := int(math.Pow(10, float64(i)))
		m := int(math.Pow(10, float64(i-1)))
		r := value % n
		digits[6-i] = int(math.Floor(float64(r / m)))
	}
	return digits
}

func meetsCriteriaPartB(value int) bool {
	var digits = digits(value)
	var foundPair = false
	for i := 0; i < 5; {
		n := digits[i]
		if digits[i+1] < n {
			return false
		}
		if digits[i+1] == n {
			run := 1
			for i < 5 && digits[i+1] == n {
				run++
				i++
			}
			if run == 2 {
				foundPair = true
			}
		} else {
			i++
		}
	}
	return foundPair
}

func meetsCriteriaPartA(value int) bool {
	var digits = digits(value)
	var foundPair = false
	for i := 0; i < 5; i++ {
		if digits[i+1] < digits[i] {
			return false
		}
		if digits[i+1] == digits[i] {
			foundPair = true
		}
	}
	return foundPair
}

func countMatchingCriteria(start, end int, check func(int) bool) int {
	count := 0
	for i := start; i < end; i++ {
		if check(i) {
			count++
		}
	}
	return count
}

func part1(start, end int) int {
	return countMatchingCriteria(start, end, meetsCriteriaPartA)
}

func part2(start, end int) int {
	return countMatchingCriteria(start, end, meetsCriteriaPartB)
}

func main() {
	fmt.Println("Part 1: ", part1(265275, 781584))
	fmt.Println("Part 2: ", part2(265275, 781584))
}
