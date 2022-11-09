package main

import (
	"fmt"
	"advent-of-code-2019/intcomp"
)

type Program []int
type binary func(int, int) int

func main() {
	p := intcomp.ReadProgram("input")
	p.Input <- 1
	go p.ProcessProgram()
	done := false
	for o := range p.Output {
		if done {
			break
		}
		if o != 0 {
			fmt.Printf("Part 1: %v\n", o)
			done = true
		}
	}
	p2 := intcomp.ReadProgram("input")
	go p2.ProcessProgram()
	p2.Input <- 5
	done = false
	for o := range p2.Output {
		if done {
			break
		}
		if o != 0 {
			fmt.Printf("Part 1: %v\n", o)
			done = true
		}
	}
}
