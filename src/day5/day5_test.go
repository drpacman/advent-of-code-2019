package main

import (
	"intcomp"
	"testing"
)

func TestProgram(t *testing.T) {
	var testProgramB intcomp.Program = intcomp.CreateProgram([]int{3, 3, 1107, -1, 8, 3, 4, 3, 99})
	go testProgramB.ProcessProgram()
	testProgramB.Input <- 1
	result := <-testProgramB.Output
	expected := 1
	if result != expected {
		t.Errorf("Got result %v, expected %v", result, expected)
	}
}
