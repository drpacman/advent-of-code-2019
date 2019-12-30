package main

import (
	"fmt"
	"io/ioutil"
	"numbers"
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

func (p *Program) value(value, mode int) int {
	if mode == 0 {
		return (*p)[value]
	}
	return value
}

func (p *Program) processBinaryInstuction(pos int, mode []int, fn binary) {
	x := p.value((*p)[pos+1], mode[0])
	y := p.value((*p)[pos+2], mode[1])
	z := (*p)[pos+3]
	(*p)[z] = fn(x, y)
}

func (p *Program) processAddition(pos int, modes []int) {
	p.processBinaryInstuction(pos, modes, func(x int, y int) int { return x + y })
}

func (p *Program) processMultiplication(pos int, modes []int) {
	p.processBinaryInstuction(pos, modes, func(x int, y int) int { return x * y })
}

func (p *Program) processLessThen(pos int, modes []int) {
	p.processBinaryInstuction(pos, modes, func(x int, y int) int {
		if x < y {
			return 1
		} else {
			return 0
		}
	})
}

func (p *Program) processEquals(pos int, modes []int) {
	p.processBinaryInstuction(pos, modes, func(x int, y int) int {
		if x == y {
			return 1
		} else {
			return 0
		}
	})
}

func (p *Program) processSetValue(pos int, input int, modes []int) {
	(*p)[(*p)[pos+1]] = input
}

func (p *Program) processGetValue(pos int, modes []int) {
	fmt.Printf("Output: %v\n", p.value((*p)[pos+1], modes[0]))
}

func (p *Program) processJump(pos int, test func(int) bool, modes []int) int {
	value := p.value((*p)[pos+1], modes[0])
	if test(value) {
		return p.value((*p)[pos+2], modes[1])
	} else {
		return pos + 3
	}
}

func (p *Program) processJumpIfZero(pos int, modes []int) int {
	return p.processJump(pos, func(value int) bool { return value == 0 }, modes)
}

func (p *Program) processJumpIfNotZero(pos int, modes []int) int {
	return p.processJump(pos, func(value int) bool { return value != 0 }, modes)
}

func (p *Program) unpackInstruction(pos int) (int, []int) {
	var digits = numbers.Digits((*p)[pos])
	var l = len(digits)
	switch l {
	case 1:
		return digits[0], []int{0, 0}
	case 2:
		return digits[0]*10 + digits[1], []int{0, 0}
	case 3:
		return digits[1]*10 + digits[2], []int{digits[0], 0}
	case 4:
		return digits[2]*10 + digits[3], []int{digits[1], digits[0]}
	default:
		panic(fmt.Sprintf("Unsupported instruction %v", (*p)[pos]))
	}
}

func (p *Program) processProgram(input int) int {
	pos := 0
	for {
		instruction, modes := p.unpackInstruction(pos)
		switch instruction {
		case 1:
			p.processAddition(pos, modes)
			pos = pos + 4
		case 2:
			p.processMultiplication(pos, modes)
			pos = pos + 4
		case 3:
			p.processSetValue(pos, input, modes)
			pos = pos + 2
		case 4:
			p.processGetValue(pos, modes)
			pos = pos + 2
		case 5:
			pos = p.processJumpIfNotZero(pos, modes)
		case 6:
			pos = p.processJumpIfZero(pos, modes)
		case 7:
			p.processLessThen(pos, modes)
			pos = pos + 4
		case 8:
			p.processEquals(pos, modes)
			pos = pos + 4
		case 99:
			return (*p)[0]
		}
	}
}

func main() {
	p := readProgram()
	fmt.Println("Part1")
	p.processProgram(1)
	fmt.Println("Part2")
	p = readProgram()
	p.processProgram(5)
}