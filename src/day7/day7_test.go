package main

import (
	"testing"
)

func TestCombinations(t *testing.T) {
	result := Combinations([]int{0,1,2,3,4})
	if len(result) != 120 {
		t.Errorf("Incorrect number of combinations - got %v expected %v", len(result), 120)
	}
}

func TestProgramPart1Example1(t *testing.T) {
	test1 := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	values1 := []int{4, 3, 2, 1, 0}
	expected := 43210
	result := runProgram(test1, values1)
	if result != expected {
		t.Errorf("Program output of %v, expected output %v", result, expected)
	}
}

func TestProgramPart1Example2(t *testing.T) {
	test1 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	values1 := []int{0, 1, 2, 3, 4}
	expected := 54321
	result := runProgram(test1, values1)
	if result != expected {
		t.Errorf("Program output of %v, expected output %v", result, expected)
	}
}

func TestProgramPart1Example3(t *testing.T) {
	test1 := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}
	values1 := []int{1, 0, 4, 3, 2}
	expected := 65210
	result := runProgram(test1, values1)
	if result != expected {
		t.Errorf("Program output of %v, expected output %v", result, expected)
	}
}

func TestProgramPart2Example1(t *testing.T) {
	test1 := []int{3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5}
	values1 := []int{9,8,7,6,5}
	expected := 139629729
	result := runProgram(test1, values1)
	if result != expected {
		t.Errorf("Program output of %v, expected output %v", result, expected)
	}
}