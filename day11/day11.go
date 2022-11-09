package main

import (
	"fmt"
	"advent-of-code-2019/intcomp"
	"math"
	"strings"
)

const NORTH = 0
const EAST = 1
const SOUTH = 2
const WEST = 3

type Move struct {
	paint int
	turn  int
}

type Coordinate struct {
	x int
	y int
}

type Action struct {
	location Coordinate
	paint    int
}

type Robot struct {
	pos       Coordinate
	direction int
	trail     []Action
}

func createRobotPart1() Robot {
	return Robot{
		pos:       Coordinate{0, 0},
		direction: NORTH,
		trail:     make([]Action, 0)}
}

func createRobotPart2() Robot {
	return Robot{
		pos:       Coordinate{0, 0},
		direction: NORTH,
		trail:     []Action{Action{Coordinate{0, 0}, 1}}}
}

func (r *Robot) getInput() int {
	current := 0
	for _, e := range r.trail {
		if r.pos == e.location {
			current = e.paint
		}
	}
	return current
}

func (r *Robot) processMove(move Move) {
	r.trail = append(r.trail, Action{r.pos, move.paint})
	if move.turn == 0 {
		switch r.direction {
		case NORTH:
			r.direction = WEST
		case EAST:
			r.direction = NORTH
		case SOUTH:
			r.direction = EAST
		case WEST:
			r.direction = SOUTH
		}
	} else {
		switch r.direction {
		case SOUTH:
			r.direction = WEST
		case WEST:
			r.direction = NORTH
		case NORTH:
			r.direction = EAST
		case EAST:
			r.direction = SOUTH
		}
	}

	switch r.direction {
	case SOUTH:
		r.pos = Coordinate{r.pos.x - 1, r.pos.y}
	case WEST:
		r.pos = Coordinate{r.pos.x, r.pos.y + 1}
	case NORTH:
		r.pos = Coordinate{r.pos.x + 1, r.pos.y}
	case EAST:
		r.pos = Coordinate{r.pos.x, r.pos.y - 1}
	}
}

func (r *Robot) tileCount() int {
	set := make(map[Coordinate]bool)
	for _, entry := range r.trail {
		set[entry.location] = true
	}
	return len(set)
}

func (r *Robot) getTrailView() string {
	minX, minY := math.MaxInt16, math.MaxInt16
	maxX, maxY := math.MinInt16, math.MinInt16
	view := make(map[Coordinate]int)
	for _, entry := range r.trail {
		switch x := entry.location.x; {
		case x < minX:
			minX = x
		case x > maxX:
			maxX = x
		}

		switch y := entry.location.y; {
		case y < minX:
			minY = y
		case y > maxX:
			maxY = y
		}
		view[entry.location] = entry.paint
	}

	var output strings.Builder
	for y := maxY; y >= minY; y-- {
		for x := minX; x <= maxX; x++ {
			colour, ok := view[Coordinate{x, y}]
			if !ok || colour == 0 {
				output.WriteString(" ")
			} else {
				output.WriteString("#")
			}
		}
		output.WriteString("\n")
	}
	return output.String()
}

func part1() {
	p := intcomp.ReadProgram("input")
	robot := createRobotPart1()
	go p.ProcessProgram()
	for !p.Halted {
		p.Input <- int64(robot.getInput())
		robot.processMove(Move{int(<-p.Output), int(<-p.Output)})
	}
	fmt.Println("Part 1: ", robot.tileCount())
}

func part2() {
	p := intcomp.ReadProgram("input")
	robot := createRobotPart2()
	go p.ProcessProgram()
	for !p.Halted {
		p.Input <- int64(robot.getInput())
		robot.processMove(Move{int(<-p.Output), int(<-p.Output)})
	}
	fmt.Println("Part 2:\n", robot.getTrailView())
}

func main() {
	part1()
	part2()
}
