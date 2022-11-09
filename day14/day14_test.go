package main

import (
	"testing"
)

func TestReadFormulas(t *testing.T) {
	fs := ReadInput("test_input_1")
	expected := Formula{[]Ingredient{Ingredient{"ORE", 10}}, Ingredient{"A", 10}}
	if !IngredientsEqual(fs[0].input, expected.input) {
		t.Errorf("Unexpected first formula, found %v expected %v", fs[0], expected)
	}
}

func TestCountOreExample1(t *testing.T) {
	fs := ReadInput("test_input_1")
	nanoFactory := CreateNanoFactory(fs)
	nanoFactory.GenerateChemical("FUEL", 1)
	result := nanoFactory.spentOre
	if result != 31 {
		t.Errorf("Unexpected count, found %v expected %v", result, 31)
	}
}

func TestCountOreExample2(t *testing.T) {
	fs := ReadInput("test_input_2")
	nanoFactory := CreateNanoFactory(fs)
	nanoFactory.GenerateChemical("FUEL", 1)
	result := nanoFactory.spentOre
	if result != 13312 {
		t.Errorf("Unexpected count, found %v expected %v", result, 13312)
	}
}

func TestCycleInput2(t *testing.T) {
	fs := ReadInput("test_input_2")
	nanoFactory := CreateNanoFactory(fs)
	result := CalculateFuelGeneratedForOre(nanoFactory, 1000000000000)
	var expected int64 = 82892753
	if result != expected {
		t.Errorf("Unexpected fuel count for 1 trillion ore, got %v, expected %v", result, expected)
	}
}

func TestCycleInput3(t *testing.T) {
	fs := ReadInput("test_input_3")
	nanoFactory := CreateNanoFactory(fs)
	result := CalculateFuelGeneratedForOre(nanoFactory, 1000000000000)
	var expected int64 = 5586022
	if result != expected {
		t.Errorf("Unexpected fuel count for 1 trillion ore, got %v, expected %v", result, expected)
	}
}
