package main

import (
	"fmt"
	"math"
	"io/ioutil"
	"strings"
	"strconv"
)

func calcFuel(m int) int {
	return int(math.Floor(float64(m/3.0))) - 2
}

func totalFuel(m int) int {
	total := calcFuel(m)
	if total > 0 {
		if extra := totalFuel(total); extra > 0 {
			total += extra
		}
	}
	return total
}

func readItems() []int {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	result := make([]int, len(lines))
	for i,v := range lines {
		result[i], _ = strconv.Atoi(v)
	}
	return result
}

func part1(items []int) int {
	sum := 0
	for _,v := range items {
		sum += calcFuel(v)
	}
	return sum	
}

func part2(items []int) int {
	sum := 0
	for _,v := range items {
		sum += totalFuel(v)
	}
	return sum	
}

func main() {
	items := readItems()
	fmt.Println("Part 1", part1(items))
	fmt.Println("Part 2", part2(items))
	
}
