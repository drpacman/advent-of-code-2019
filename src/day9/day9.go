package main

import (
	"fmt"
	"intcomp"
)

func main() {
	p := intcomp.ReadProgram("input")
	go p.ProcessProgram()
	p.Input <- 1
	fmt.Printf("Part 1: %v\n", <-p.Output)
	p = intcomp.ReadProgram("input")
	go p.ProcessProgram()
	p.Input <- 2
	fmt.Printf("Part 2: %v\n", <-p.Output)
}
