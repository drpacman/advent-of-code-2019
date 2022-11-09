package intcomp

import (
	"advent-of-code-2019/files"
	"advent-of-code-2019/numbers"
	"fmt"
	"strconv"
)

type Program struct {
	instructions []int64
	Input        chan int64
	Output       chan int64
	Halted       bool
	relativeBase int64
}

type binary func(int64, int64) int64

func ReadProgram(filepath string) Program {
	strs := files.ReadSingleLineCSV(filepath)
	ints := make([]int64, len(strs))
	for i := 0; i < len(strs); i++ {
		n, _ := strconv.Atoi(strs[i])
		ints[i] = int64(n)
	}
	return Program{
		instructions: ints,
		Input:        make(chan int64, 1000),
		Output:       make(chan int64, 1000),
		Halted:       false,
		relativeBase: 0}
}

func CreateProgram(instructions []int64) Program {
	var ins = make([]int64, len(instructions))
	copy(ins, instructions)
	return Program{
		instructions: ins,
		Input:        make(chan int64, 1000),
		Output:       make(chan int64, 1000)}
}

func (p *Program) Poke(value int64, index int) {
	p.instructions[index] = value
}

func (p *Program) value(value, mode int64) int64 {
	if mode == 0 {
		return p.readMemory(value)
	} else if mode == 2 {
		return p.readMemory(p.relativeBase + value)
	}
	return value
}

func (p *Program) readMemory(pos int64) int64 {
	if pos >= int64(len(p.instructions)) {
		return 0
	}
	return p.instructions[pos]
}

func (p *Program) writeMemory(pos, mode, value int64) {
	if mode == 2 {
		pos = pos + p.relativeBase
	}
	currLimit := int64(len(p.instructions))
	if pos >= currLimit {
		memoryExtension := make([]int64, pos-currLimit+1)
		// extend the memory
		p.instructions = append(p.instructions, memoryExtension...)
	}
	p.instructions[pos] = value
}

func (p *Program) processBinaryInstuction(pos int64, mode []int64, fn binary) {
	x := p.value(p.readMemory(pos+1), mode[0])
	y := p.value(p.readMemory(pos+2), mode[1])
	p.writeMemory(p.readMemory(pos+3), mode[2], fn(x, y))
}

func (p *Program) processAddition(pos int64, modes []int64) {
	p.processBinaryInstuction(pos, modes, func(x int64, y int64) int64 { return x + y })
}

func (p *Program) processMultiplication(pos int64, modes []int64) {
	p.processBinaryInstuction(pos, modes, func(x int64, y int64) int64 { return x * y })
}

func (p *Program) processLessThen(pos int64, modes []int64) {
	p.processBinaryInstuction(pos, modes, func(x int64, y int64) int64 {
		if x < y {
			return 1
		} else {
			return 0
		}
	})
}

func (p *Program) processEquals(pos int64, modes []int64) {
	p.processBinaryInstuction(pos, modes, func(x int64, y int64) int64 {
		if x == y {
			return 1
		} else {
			return 0
		}
	})
}

func (p *Program) processSetValue(pos int64, modes []int64) {
	p.writeMemory(p.readMemory(pos+1), modes[0], <-p.Input)
}

func (p *Program) processGetValue(pos int64, modes []int64) {
	var x = p.value(p.readMemory(pos+1), modes[0])
	p.Output <- x
}

func (p *Program) processJump(pos int64, test func(int64) bool, modes []int64) int64 {
	value := p.value(p.readMemory(pos+1), modes[0])
	if test(value) {
		return p.value(p.readMemory(pos+2), modes[1])
	} else {
		return pos + 3
	}
}

func (p *Program) processJumpIfZero(pos int64, modes []int64) int64 {
	return p.processJump(pos, func(value int64) bool { return value == 0 }, modes)
}

func (p *Program) processJumpIfNotZero(pos int64, modes []int64) int64 {
	return p.processJump(pos, func(value int64) bool { return value != 0 }, modes)
}

func (p *Program) processUpdateRelativeBase(pos int64, modes []int64) {
	p.relativeBase += p.value(p.readMemory(pos+1), modes[0])
}

func (p *Program) unpackInstruction(pos int64) (int64, []int64) {
	var digits = numbers.Digits(p.readMemory(pos))
	var l = len(digits)
	switch l {
	case 1:
		return int64(digits[0]), []int64{0, 0, 0}
	case 2:
		return int64(digits[0]*10 + digits[1]), []int64{0, 0, 0}
	case 3:
		return int64(digits[1]*10 + digits[2]), []int64{int64(digits[0]), 0, 0}
	case 4:
		return int64(digits[2]*10 + digits[3]), []int64{int64(digits[1]), int64(digits[0]), 0}
	case 5:
		return int64(digits[3]*10 + digits[4]), []int64{int64(digits[2]), int64(digits[1]), int64(digits[0])}
	default:
		panic(fmt.Sprintf("Unsupported instruction %v", p.readMemory(pos)))
	}
}

func (p *Program) ProcessProgram() {
	var pos int64 = 0
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
			p.processSetValue(pos, modes)
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
		case 9:
			p.processUpdateRelativeBase(pos, modes)
			pos = pos + 2
		case 99:
			p.Output <- (*p).instructions[0]
			close(p.Output)
			p.Halted = true
			return
		default:
			panic(fmt.Errorf("Unknown instruction %v at position %v", instruction, pos))
		}
	}
}
