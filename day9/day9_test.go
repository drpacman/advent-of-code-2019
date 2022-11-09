package main

import (
	"advent-of-code-2019/intcomp"
	"math"
	"testing"
)

func TestProgramExample1(t *testing.T) {
	var contents = []int64{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}
	var p = intcomp.CreateProgram(contents)
	go p.ProcessProgram()
	for _, c := range contents {
		v := <-p.Output
		if v != c {
			t.Errorf("Expected %v got %v", c, v)
		}
	}

}

func TestProgramExample2(t *testing.T) {
	var contents = []int64{1102, 34915192, 34915192, 7, 4, 7, 99, 0}
	var p = intcomp.CreateProgram(contents)
	go p.ProcessProgram()
	result := <-p.Output
	expected := int64(math.Pow10(16))
	if result > expected {
		t.Errorf("Expected %v to be greater then %v", result, expected)
	}

}

func TestProgramExample3(t *testing.T) {
	var contents = []int64{104, 1125899906842624, 99}
	var p = intcomp.CreateProgram(contents)
	go p.ProcessProgram()
	result := <-p.Output
	expected := int64(1125899906842624)
	if result != expected {
		t.Errorf("Expected %v to be %v", result, expected)
	}
}

func TestRelativeSet(t *testing.T) {
	var contents = []int64{109, 1, 203, 6, 204, 6, 99, -1}
	var p = intcomp.CreateProgram(contents)
	go p.ProcessProgram()
	p.Input <- 1
	result := <-p.Output
	expected := int64(1)
	if result != expected {
		t.Errorf("Expected %v to be %v", result, expected)
	}
}
