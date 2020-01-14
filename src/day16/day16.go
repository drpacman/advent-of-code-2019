package main

import (
	"files"
	"fmt"
	"strconv"
)

func ReadInput(path string) []int {
	line := files.ReadLines(path)[0]
	return stringToInput(line)
}

func stringToInput(s string) []int {
	input := make([]int, 0)
	for i := 0; i < len(s); i++ {
		n, _ := strconv.Atoi(s[i : i+1])
		input = append(input, n)
	}
	return input
}

func lastDigit(a int64) int {
	s := fmt.Sprintf("%v", a)
	digit, _ := strconv.Atoi(s[len(s)-1:])
	return digit
}

func multiPhaseFtt(input, base []int, phase int) []int {
	in := input
	for i := 0; i < phase; i++ {
		in = FTT(in, base)
	}
	return in
}
func FTT(input, base []int) []int {
	output := make([]int, len(input))
	for i := 1; i <= len(input); i++ {
		var value int64 = 0
		for j := 0; j < len(input); j++ {
			basePos := (j + 1) / i
			m := base[basePos%len(base)]
			value += int64((input[j] * m))
		}
		output[i-1] = lastDigit(value)
	}
	return output
}

/*
9898

(1) 10-10
(2) 0110
(3) 0011
(4) 0001

for the digits over half way through thecomplete list their value
is sum of digits from the end mod 10
so
* last digit is in the above is 8
* last but one digit is 9+8 % 10 = 7

for next phase
* last digit remain unchanged - is 8
* next digit is 7 + 8 % 10= 5

for next phase
* last digit remain unchanged - is 8
* next digit is 5 + 8 % 10= 3

*/

// for the digits over half way through the complete list their value
// at the end of the phase is sum of digits from the end mod 10
func calcTail(input []int) []int {
	// get first 7 digits from input to get our target tail position
	targetPos := 0
	for i := 0; i < 7; i++ {
		targetPos = targetPos*10 + input[i]
	}
	srcLen := len(input)
	tailPos := (srcLen * 10000) - targetPos
	repeats := 1 + tailPos/srcLen
	remainder := tailPos % srcLen
	// initialise the target with suitable copies of the input
	target := make([]int, 0)
	for i := 0; i < repeats; i++ {
		target = append(target, input...)
	}
	// drop the excessive bit at the front
	target = target[srcLen-remainder:]
	// run the phases
	for p := 0; p < 100; p++ {
		for i := len(target) - 2; i >= 0; i-- {
			target[i] = (target[i+1] + target[i]) % 10
		}
	}
	return target[0:8]
}

func part2() {
	s := ReadInput("input")
	fmt.Println("Part 2", calcTail(s))
}

func part1() {
	s := ReadInput("input")
	base := []int{0, 1, 0, -1}
	fmt.Println("Part 1", multiPhaseFtt(s, base, 100)[0:8])
}

func main() {
	part1()
	part2()
}
