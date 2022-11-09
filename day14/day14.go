package main

import (
	"advent-of-code-2019/files"
	"fmt"
	"strconv"
	"strings"
)

type Ingredient struct {
	chemical string
	quanity  int
}

type Formula struct {
	input     []Ingredient
	generated Ingredient
}

type NanoFactory struct {
	options  map[string]Formula
	reserves map[string]int64
	spentOre int64
}

func IngredientsEqual(a, b []Ingredient) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func ReservesEqual(a, b map[string]int64) bool {
	if len(a) != len(b) {
		return false
	}
	for v := range a {
		if a[v] != b[v] {
			return false
		}
	}
	for v := range b {
		if b[v] != a[v] {
			return false
		}
	}
	return true
}

func ReadInput(filepath string) []Formula {
	lines := files.ReadLines(filepath)
	formulas := make([]Formula, len(lines))
	for i, line := range lines {
		io := strings.Split(line, "=>")
		inputs := strings.Split(io[0], ",")
		ingredients := make([]Ingredient, len(inputs))
		for j, item := range inputs {
			in := strings.Split(strings.Trim(item, " "), " ")
			quantity, _ := strconv.Atoi(in[0])
			ingredients[j] = Ingredient{in[1], quantity}
		}
		out := strings.Split(strings.Trim(io[1], " "), " ")
		quantity, _ := strconv.Atoi(out[0])
		formulas[i] = Formula{input: ingredients, generated: Ingredient{out[1], quantity}}
	}
	return formulas
}

func CreateNanoFactory(formulas []Formula) NanoFactory {
	options := make(map[string]Formula)
	for _, f := range formulas {
		key := f.generated.chemical
		_, ok := options[key]
		if ok {
			panic(fmt.Sprintf("Found duplicate formula for chemical %v", key))
		}
		options[key] = f
	}

	return NanoFactory{
		options:  options,
		reserves: make(map[string]int64),
		spentOre: 0}
}

func (f *NanoFactory) GenerateChemical(chemical string, quantity int) {
	required := quantity - f.useQuantity(chemical, quantity)
	for required > 0 {
		if chemical == "ORE" {
			f.reserves["ORE"] += int64(quantity)
			f.spentOre += int64(quantity)
		} else {
			formula, ok := f.options[chemical]
			if !ok {
				panic(fmt.Sprintf("Unable to find chemical %v", chemical))
			}
			for _, ingredient := range formula.input {
				f.GenerateChemical(ingredient.chemical, ingredient.quanity)
			}
			f.reserves[chemical] += int64(formula.generated.quanity)
		}
		required -= f.useQuantity(chemical, required)
	}
}

func (f *NanoFactory) useQuantity(chemical string, quantity int) int {
	removed := 0
	available, ok := f.reserves[chemical]
	if ok {
		if available >= int64(quantity) {
			removed = quantity
			f.reserves[chemical] = available - int64(quantity)
		} else {
			removed = int(available)
			f.reserves[chemical] = 0
		}
	}
	return removed
}

func part1() {
	fs := ReadInput("input")
	factory := CreateNanoFactory(fs)
	factory.GenerateChemical("FUEL", 1)
	fmt.Println("Part 1: ", factory.spentOre)
}

func CalculateFuelGeneratedForOre(factory NanoFactory, availableOre int64) int64 {
	var fuelPerCycle int64 = 0
	var orePerCycle int64 = factory.spentOre
	var history = make([]map[string]int64, 0)
	done := false
Cycle:
	for !done {
		factory.GenerateChemical("FUEL", 1)
		for _, entry := range history {
			if ReservesEqual(entry, factory.reserves) {
				done = true
				break Cycle
			}
		}

		reserves := make(map[string]int64)
		for k := range factory.reserves {
			reserves[k] = factory.reserves[k]
		}
		history = append(history, reserves)

		fuelPerCycle++
		orePerCycle = factory.spentOre
	}
	fmt.Printf("Cycles at %v\n", fuelPerCycle)
	numFullCycles := availableOre / orePerCycle
	oreFromCycles := numFullCycles * orePerCycle
	remainder := factory.spentOre + (availableOre - oreFromCycles)
	additionalFuel := 0
	for factory.spentOre < remainder {
		factory.GenerateChemical("FUEL", 1)
		additionalFuel++
	}
	fmt.Printf("Num full cycles is %v, remaining ore after cycle is %v, additional fuel is %v\n", numFullCycles, remainder, additionalFuel)
	return numFullCycles*fuelPerCycle + int64(additionalFuel) - 1
}
func part2() {
	fs := ReadInput("input")
	factory := CreateNanoFactory(fs)
	result := CalculateFuelGeneratedForOre(factory, 1000000000000)
	fmt.Printf("Part 2: %v\n", result)
}

func main() {
	part1()
	part2()
}
