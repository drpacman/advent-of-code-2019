package main

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSequenceFromFile(t *testing.T) {
	var instruction = ReadInputAsLinearCongruentialGenerator("test_input", 10)
	var expected = int64(0)
	var result = instruction.posForCard(9)
	if result != expected {
		t.Errorf("Got result %v instead of expected %v", result, expected)
	}
}

func TestModPow(t *testing.T) {
	result := modPow(3, 10, 1000000)
	expected := int64(59049)
	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}

	result = modPow(3, 10, 100)
	expected = int64(49)
	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func TestNTimes(t *testing.T) {
	instruction := ReadInputAsLinearCongruentialGenerator("input1", 119315717514047)
	fmt.Println(instruction)
	fmt.Println("Single composed LCG: ", instruction)
	instruction = instruction.applyNTimes(119315717514046)
	result := instruction.posForCard(1)
	expected := int64(1)
	if result != expected {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func TestLocallyImplementedModPowAgainstBigInt(t *testing.T) {
	k := int64(10315717514047)
	n := int64(119315717514047)
	result := modPow(10, k, n)
	expected := big.NewInt(10)
	expected.Exp(expected, big.NewInt(k), big.NewInt(n))
	if result != expected.Int64() {
		t.Errorf("Expected %v got %v", expected, result)
	}
}

func TestLocallyImplementedModInverseAgainstBigInt(t *testing.T) {
	size := int64(119315717514047)
	result := modInverse(2020, size)
	expected := big.NewInt(0)
	expected.ModInverse(big.NewInt(2020), big.NewInt(size))
	if result != expected.Int64() {
		t.Errorf("Expected %v got %v", expected, result)
	}
}
