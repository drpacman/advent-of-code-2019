package main

import (
	"fmt"
	"intcomp"
	"numbers"
)

func runAmplifiers(createProgram func() intcomp.Program, values []int) int {
	amplifiers := make([]intcomp.Program, len(values))
	for i := 0; i < len(values); i++ {
		amplifiers[i] = createProgram()
		go amplifiers[i].ProcessProgram()
		amplifiers[i].Input <- values[i]
	}

	input := 0
	for !amplifiers[0].Halted {
		amplifiers[0].Input <- input
		for i := 0; i < len(amplifiers)-1; i++ {
			amplifiers[i+1].Input <- <-amplifiers[i].Output
		}
		input = <-amplifiers[4].Output
	}
	return input
}

func runProgram(instructions []int, values []int) int {
	var programFactory = func() intcomp.Program {
		return intcomp.CreateProgram(instructions)
	}
	return runAmplifiers(programFactory, values)
}

func run(phaseValues []int) (int, []int) {
	var programFactory = func() intcomp.Program {
		return intcomp.ReadProgram("input")
	}
	var max = 0
	var maxValues = []int{0, 0, 0, 0, 0}
	for _, values := range numbers.Combinations(phaseValues) {
		var n = runAmplifiers(programFactory, values)
		if n > max {
			max = n
			maxValues = values
		}
	}
	return max, maxValues
}

func part1() (int, []int) {
	return run([]int{0, 1, 2, 3, 4})
}

func part2() (int, []int) {
	return run([]int{5, 6, 7, 8, 9})
}

func main() {
	max, maxValues := part1()
	fmt.Printf("Part1 %v with values %v\n", max, maxValues)
	max, maxValues = part2()
	fmt.Printf("Part2 %v with values %v\n", max, maxValues)

}
