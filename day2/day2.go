package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Program []int
type binary func(int, int) int

func readProgram() Program {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic("unable to read file")
	}
	strs := strings.Split(string(data), ",")
	ints := make([]int, len(strs))
	for i := 0; i < len(strs); i++ {
		ints[i], _ = strconv.Atoi(strs[i])
	}
	return ints
}

func (p *Program) processInstuction(pos int, fn binary) {
	x := (*p)[pos+1]
	y := (*p)[pos+2]
	z := (*p)[pos+3]
	(*p)[z] = fn((*p)[x], (*p)[y])
}

func (p *Program) processAddition(pos int) {
	p.processInstuction(pos, func(x int, y int) int { return x + y })
}

func (p *Program) processMultiplication(pos int) {
	p.processInstuction(pos, func(x int, y int) int { return x * y })
}

func (p *Program) processProgram() int {
	pos := 0
	for {
		switch (*p)[pos] {
		case 1:
			p.processAddition(pos)
			pos = pos + 4
		case 2:
			p.processMultiplication(pos)
			pos = pos + 4
		case 99:
			return (*p)[0]
		}
	}
}

func part1() int {
	var p = readProgram()
	p[1] = 12
	p[2] = 2
	return p.processProgram()
}

func part2() (int, error) {
	input := readProgram()
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			var target Program = make([]int, len(input))
			copy(target, input)
			target[1] = noun
			target[2] = verb
			if target.processProgram() == 19690720 {
				return 100*noun + verb, nil
			}
		}
	}
	return -1, errors.New("Failed to find answer")
}

func main() {
	fmt.Println("Part 1:", part1())
	part2, _ := part2()
	fmt.Println("Part 2:", part2)
}
