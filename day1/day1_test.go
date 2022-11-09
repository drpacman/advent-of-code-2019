package main

import "testing"

func TestCalcValue(t *testing.T) {
	if b := calcFuel(100756); b != 33583 {
		t.Errorf("Value %v not equal to %v", b, 33583)
	}
} 