package main

import (
	"fmt"
	"intcomp"
)

func runAmplifiers(createProgram func() intcomp.Program, values []int) int {
	amplifiers := make([]intcomp.Program, 5)
	for i := 0; i < 5; i++ {
		amplifiers[i] = createProgram()
		go amplifiers[i].ProcessProgram()
		amplifiers[i].Input <- values[i]
	}

	input := 0
	for !amplifiers[0].Halted {
		amplifiers[0].Input <- input
		amplifiers[1].Input <- <-amplifiers[0].Output
		amplifiers[2].Input <- <-amplifiers[1].Output
		amplifiers[3].Input <- <-amplifiers[2].Output
		amplifiers[4].Input <- <-amplifiers[3].Output

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

func Combinations(combo []int) [][]int {
	if len(combo) == 1 {
		return [][]int{[]int{combo[0]}}
	} else {
		result := make([][]int, 0)
		for i := 0; i < len(combo); i++ {
			remainder := make([]int, 0)
			remainder = append(remainder, combo[:i]...)
			remainder = append(remainder, combo[i+1:]...)
			var tail [][]int = Combinations(remainder)
			for j := 0; j < len(tail); j++ {
				result = append(result, append([]int{combo[i]}, tail[j]...))
			}
		}
		return result
	}
}

func run(phaseValues []int) (int, []int) {
	var programFactory = func() intcomp.Program {
		return intcomp.ReadProgram("input")
	}
	var max = 0
	var maxValues = []int{0, 0, 0, 0, 0}
	for _, values := range Combinations(phaseValues) {
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
