package main

import (
	"testing"
)

func TestPanelCount(t *testing.T) {
	input := []int{0, 1, 0, 0, 1, 0, 1, 0, 0, 1, 1, 0, 1, 0}
	robot := Robot{}
	for i := 0; i < len(input); i = i + 2 {
		robot.processMove(Move{input[i], input[i+1]})
	}
	expected := 6
	result := robot.tileCount()
	if result != expected {
		t.Errorf("Number of painted tiles was %v, but expected %v", result, expected)
	}
}
