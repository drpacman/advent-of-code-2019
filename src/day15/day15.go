package main

import (
	"display"
	"fmt"
	"intcomp"
	"numbers"
	"render"
	"time"
)

const (
	NORTH   = 1
	SOUTH   = 2
	EAST    = 4
	WEST    = 3
	UNKNOWN = 0
	WALL    = 1
	EMPTY   = 2
	OXYGEN  = 3
	DROID   = 4
	VISITED = 5
)

type Point struct {
	x, y int
}

func (p1 Point) dist(p2 Point) int {
	return numbers.AbsInt(p1.x-p2.x) + numbers.AbsInt(p1.y-p2.y)
}

type Maze struct {
	contents map[Point]int
	droid    Point
}

func (m *Maze) applyInformation(direction int, outcome int) bool {
	visited := false
	target := m.droid
	switch direction {
	case NORTH:
		target = Point{m.droid.x, m.droid.y - 1}
	case SOUTH:
		target = Point{m.droid.x, m.droid.y + 1}
	case EAST:
		target = Point{m.droid.x + 1, m.droid.y}
	case WEST:
		target = Point{m.droid.x - 1, m.droid.y}
	}
	if outcome == EMPTY {
		switch m.contents[target] {
		case UNKNOWN:
			m.contents[target] = EMPTY
		case EMPTY:
			m.contents[target] = VISITED
		case VISITED:
			visited = true
		}
	} else {
		m.contents[target] = outcome
	}

	if outcome != WALL {
		m.droid = target
	}
	return visited
}

func (m *Maze) getMoves() []int {
	var newMoves = make([]int, 0)
	var oldMoves = make([]int, 0)
	var options = map[int]Point{
		NORTH: Point{m.droid.x, m.droid.y - 1},
		SOUTH: Point{m.droid.x, m.droid.y + 1},
		EAST:  Point{m.droid.x + 1, m.droid.y},
		WEST:  Point{m.droid.x - 1, m.droid.y}}
	for move := range options {
		entry, ok := m.contents[options[move]]
		if !ok {
			newMoves = append(newMoves, move)
		} else {
			if entry == UNKNOWN {
				newMoves = append(newMoves, move)
			} else if entry == EMPTY {
				oldMoves = append(oldMoves, move)
			}
		}
	}
	return append(newMoves, oldMoves...)
}

func (m *Maze) getEmptySlotsAt(p Point) []Point {
	var emptySlots = make([]Point, 0)
	var slots = []Point{
		Point{p.x, p.y - 1},
		Point{p.x, p.y + 1},
		Point{p.x + 1, p.y},
		Point{p.x - 1, p.y}}
	for _, slot := range slots {
		entry, ok := m.contents[slot]
		if ok && (entry == EMPTY || entry == VISITED) {
			emptySlots = append(emptySlots, slot)
		}
	}
	return emptySlots
}

func (m *Maze) generateBoard() [][]int64 {
	// convert map to 2D int array
	minX, maxX, minY, maxY := 0, 0, 0, 0
	for p := range m.contents {
		if p.x > maxX {
			maxX = p.x
		}
		if p.x < minX {
			minX = p.x
		}
		if p.y > maxY {
			maxY = p.y
		}
		if p.y < minY {
			minY = p.y
		}
	}
	cols := numbers.AbsInt(maxX-minX) + 1
	rows := numbers.AbsInt(maxY-minY) + 1
	entries := make([][]int64, rows)
	for y := 0; y < rows; y++ {
		entries[y] = make([]int64, cols)
		for x := 0; x < cols; x++ {
			pos := Point{x + minX, y + minY}
			entry := m.contents[pos]
			entries[y][x] = int64(entry)
		}
	}
	entries[m.droid.y-minY][m.droid.x-minX] = DROID
	return entries
}

func (m Maze) render() string {
	chars := map[int64]string{
		UNKNOWN: "?",
		EMPTY:   ".",
		WALL:    "#",
		OXYGEN:  "X",
		DROID:   "O",
		VISITED: "-",
	}
	return render.GenerateBoardText(m.generateBoard(), chars)
}

func (m Maze) renderDisplay(display *display.Display) {
	chars := map[int64]string{
		UNKNOWN: "?",
		EMPTY:   " ",
		WALL:    "#",
		OXYGEN:  "X",
		DROID:   "O",
		VISITED: " "}
	display.SetContent(render.GenerateBoardHtml(m.generateBoard(), chars))
}

func (m *Maze) exploreMaze(p intcomp.Program, d *display.Display) Point {
	var oxygen Point
	direction := SOUTH
	reverse := NORTH
	reversing := false
	breadcrumb := make([]int, 0)
	for {
		p.Input <- int64(direction)
		moved := true
		switch <-p.Output {
		case 0:
			m.applyInformation(direction, WALL)
			moved = false
		case 1:
			m.applyInformation(direction, EMPTY)
		case 2:
			m.applyInformation(direction, OXYGEN)
			oxygen = m.droid
		}

		if !reversing && moved {
			switch direction {
			case NORTH:
				reverse = SOUTH
			case SOUTH:
				reverse = NORTH
			case EAST:
				reverse = WEST
			case WEST:
				reverse = EAST
			}
			breadcrumb = append(breadcrumb, reverse)
		}

		moves := m.getMoves()
		if len(moves) == 0 {
			if len(breadcrumb) > 0 {
				direction = breadcrumb[len(breadcrumb)-1]
				breadcrumb = breadcrumb[:len(breadcrumb)-1]
				reversing = true
			} else {
				return oxygen
			}
		} else {
			direction = moves[0]
			reversing = false
		}

		m.renderDisplay(d)
	}
}

func (m Maze) fillWithOxygen() int {
	count := 0
	start := Point{-16, -20}
	targets := [][]Point{m.getEmptySlotsAt(start)}
	for {
		newTargets := make([][]Point, 0)
		for _, target := range targets {
			for _, slot := range target {
				m.contents[slot] = OXYGEN
				newTargets = append(newTargets, m.getEmptySlotsAt(slot))
			}
		}
		if len(newTargets) > 0 {
			count++
			targets = newTargets
		} else {
			break
		}
	}
	return count
}

func (m Maze) findShortestRoute(start, target Point) int {
	paths := [][]Point{[]Point{start}}
	current := start

	for len(paths) > 0 {
		newPaths := make([][]Point, 0)
		for _, path := range paths {
			current = path[len(path)-1]
			// find all the enties around the position which aren't walls
			options := m.getEmptySlotsAt(current)

		Options:
			for _, o := range options {
				if o == target {
					fmt.Println("Got home in ", len(path))
					return len(path)
				}
				// ignore if we have already been there on this path
				for _, p := range path {
					if p == o {
						continue Options
					}
				}
				clone := make([]Point, len(path))
				copy(clone, path)
				newPaths = append(newPaths, append(clone, o))
			}
		}
		paths = newPaths
	}
	return 0
}

func run(d *display.Display) {
	defer d.Close()
	m := Maze{
		contents: make(map[Point]int),
		droid:    Point{0, 0}}
	p := intcomp.ReadProgram("input")
	go p.ProcessProgram()
	pos := m.exploreMaze(p, d)

	m.renderDisplay(d)
	time.Sleep(time.Duration(0) * time.Second)
	fmt.Println("Part 1", m.findShortestRoute(pos, Point{0, 0}))
	fmt.Println("Part 2", m.fillWithOxygen())
	m.renderDisplay(d)
	time.Sleep(time.Duration(0) * time.Second)
}

func main() {
	d := display.CreateDisplay(800, 800)
	go run(d)
	d.Show()
}

/*
Point src, dst;// Source and destination coordinates
// cur also indicates the coordinates of the current location
int MD_best = MD(src, dst);// It stores the closest MD we ever had to dst
// A productive path is the one that makes our MD to dst smaller
while (cur != dst) {
    if (there exists a productive path) {
        Take the productive path;
    } else {
        MD_best = MD(cur, dst);
        Imagine a line between cur and dst;
        Take the first path in the left/right of the line; // The left/right selection affects the following hand rule
        while (MD(cur, dst) != MD_best || there does not exist a productive path) {
            Follow the right-hand/left-hand rule; // The opposite of the selected side of the line
    }
}
*/
