package main

import (
	"testing"
)

func TestAddition(t *testing.T) {
	var p = Program{1, 0, 0, 0}
	p.processAddition(0)
	if p[0] != 2 {
		t.Errorf("Failed to calculate addition correctly got %v instead of %v", p[0], 2)
	}
}

func TestMultiplication(t *testing.T) {
	var p = Program{2, 3, 0, 3}
	p.processMultiplication(0)
	if p[3] != 6 {
		t.Errorf("Failed to calculate addition correctly got %v instead of %v", p[3], 6)
	}
}

func TestProgram(t *testing.T) {
	var p = Program{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	p.processProgram()
	if p[0] != 3500 {
		t.Errorf("Failed to execute program correctly got %v instead of %v", p[0], 3500)
	}
}
