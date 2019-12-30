package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

//Wire Holder for a wire
type Wire struct {
	dir byte
	len int
}

type Wires []Wire
type coords struct {
	x int
	y int
}

func (w Wire) String() string {
	return fmt.Sprintf("Wire %v length %v", w.dir, w.len)
}

func (w Wires) getCoords() []coords {
	result := make([]coords, 0)
	start := coords{0, 0}
	for i := 0; i < len(w); i++ {
		var wireCoords = w[i].getCoords(start)
		result = append(result, wireCoords...)
		start = wireCoords[len(wireCoords)-1]
	}
	return result
}

func (w Wire) getCoords(start coords) []coords {
	result := make([]coords, w.len)
	for i := 0; i < w.len; i++ {
		switch w.dir {
		case 'R':
			result[i] = coords{start.x + 1 + i, start.y}
		case 'L':
			result[i] = coords{start.x - 1 - i, start.y}
		case 'U':
			result[i] = coords{start.x, start.y + 1 + i}
		case 'D':
			result[i] = coords{start.x, start.y - 1 - i}
		}
	}
	return result
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func closestCrossPoint(w1, w2 Wires) int {
	w1Coords := w1.getCoords()
	w2Coords := w2.getCoords()

	m := make(map[coords]bool)
	for _, item := range w1Coords {
		m[item] = true
	}
	var closest = math.MaxInt32
	for _, coord := range w2Coords {
		if _, ok := m[coord]; ok {
			var dist = abs(coord.x) + abs(coord.y)
			if dist < closest {
				closest = dist
			}
		}
	}
	return closest
}

func shortestStepsCrossPoint(w1, w2 Wires) int {
	w1Coords := w1.getCoords()
	w2Coords := w2.getCoords()

	m := make(map[coords]int)
	for i, item := range w1Coords {
		m[item] = i + 1
	}
	var shortest = math.MaxInt32
	for j, coord := range w2Coords {
		if i, ok := m[coord]; ok {
			var dist = i + j + 1
			if dist < shortest {
				shortest = dist
			}
		}
	}
	return shortest
}

func readInput() [2]Wires {
	data, _ := ioutil.ReadFile("input")
	lines := strings.Split(string(data), "\n")
	var wires [2]Wires
	for j := 0; j < 2; j++ {
		entries := strings.Split(string(lines[j]), ",")
		wires[j] = make([]Wire, len(entries))
		for i := 0; i < len(entries); i++ {
			len, _ := strconv.Atoi(entries[i][1:])
			wires[j][i] = Wire{entries[i][0], len}
		}
	}
	return wires
}

func main() {
	wires := readInput()
	fmt.Printf("Part 1 : %v\n", closestCrossPoint(wires[0], wires[1]))
	fmt.Printf("Part 2 : %v\n", shortestStepsCrossPoint(wires[0], wires[1]))
}
