package main

import (
	"fmt"
	"numbers"
)

func meetsCriteriaPartB(value int) bool {
	var digits = numbers.Digits(value)
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
	var digits = numbers.Digits(value)
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
