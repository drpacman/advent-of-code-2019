package main

import (
	"testing"
)

func TestReadInput(t *testing.T) {
	input := ReadInput("input")
	if len(input) != 650 {
		t.Errorf("Failed to read full input, expected len 650, received %v", len(input))
	}
}

func equals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, x := range a {
		if x != b[i] {
			return false
		}
	}
	return true
}
func testFTT(input, expected string, phases int, t *testing.T) {
	base := []int{0, 1, 0, -1}
	output := multiPhaseFtt(stringToInput(input), base, phases)[0:8]
	expectedOutput := stringToInput(expected)
	if !equals(output, expectedOutput) {
		t.Errorf("Expected output of FTT is %v, received %v", expectedOutput, output)
	}
}
func TestFTTSimple(t *testing.T) {
	testFTT("12345678", "48226158", 1, t)
	testFTT("48226158", "34040438", 1, t)
	testFTT("34040438", "03415518", 1, t)
	testFTT("03415518", "01029498", 1, t)
	testFTT("80871224585914546619083218645595", "24176176", 100, t)
	testFTT("19617804207202209144916044189917", "73745418", 100, t)
	testFTT("69317163492948606335995924319873", "52432133", 100, t)
}

func TestFTTRepeat(t *testing.T) {
	in := stringToInput("9898")
	base := []int{0, 1, 0, -1}
	output := FTT(in, base)
	expectedLastDigits := []int{7, 8}
	if !equals(output[2:4], expectedLastDigits) {
		t.Errorf("Expected output of FTT is %v, received %v", expectedLastDigits, output)
	}
	output = FTT(output, base)
	expectedLastDigits = []int{5, 8}
	if !equals(output[2:4], expectedLastDigits) {
		t.Errorf("Expected output of FTT is %v, received %v", expectedLastDigits, output)
	}
	output = FTT(output, base)
	expectedLastDigits = []int{3, 8}
	if !equals(output[2:4], expectedLastDigits) {
		t.Errorf("Expected output of FTT is %v, received %v", expectedLastDigits, output)
	}
}

func TestCalcTail(t *testing.T) {
	input := stringToInput("03036732577212944063491565474664")
	output := calcTail(input)
	expected := stringToInput("84462026")
	if !equals(output, expected) {
		t.Errorf("Expected output of FTT is %v, received %v", expected, output)
	}
}
