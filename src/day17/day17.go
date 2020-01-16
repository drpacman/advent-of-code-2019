package main

import (
	"display"
	"fmt"
	"intcomp"
	"render"
	"strings"
	"time"
)

type Point struct {
	x, y int
}

type Dir struct {
	dx, dy int
}

type Leg struct {
	turn int
	len  int
}

const L = 0
const R = 1

var UP = Dir{0, -1}
var RIGHT = Dir{1, 0}
var LEFT = Dir{-1, 0}
var DOWN = Dir{0, 1}

var STATE_INIT = 0
var STATE_READY = 1
var STATE_RUNNING = 2
var STATE_COMPLETE = 3

type Game struct {
	state    int
	board    [][]int64
	in       chan int64
	out      chan int64
	graphics *display.Display
}

func (g *Game) readBoard(firstChar int64) {
	NEWLINE := int64(10)
	rows := make([][]int64, 0)
	row := make([]int64, 0)
	if firstChar > 0 {
		row = append(row, firstChar)
	}
Read:
	for c := range g.in {
		switch {
		case c == NEWLINE && len(row) == 0:
			break Read
		case c == NEWLINE:
			rows = append(rows, row)
			row = make([]int64, 0)
		default:
			row = append(row, c)
		}
	}
	g.board = rows
	if g.state == STATE_INIT {
		g.state = STATE_READY
	}
}

func (g *Game) processInput() {
Loop:
	for c := range g.in {
		if c > 127 {
			g.state = STATE_COMPLETE
			s := fmt.Sprintf("Result is %v", c)
			fmt.Println(s)
			g.graphics.SetContent(s)
			return
		}

		if c == '.' {
			g.readBoard(c)
			g.graphics.SetContent(render.GenerateBoardHtml(g.board, nil))
			time.Sleep(time.Duration(0) * time.Millisecond)
			continue Loop
		}

		s := strings.Builder{}
		s.WriteRune(rune(c))
		for c2 := range g.in {
			s.WriteRune(rune(c2))
			if c2 == '\n' {
				fmt.Println(s.String())
				g.graphics.SetContent(s.String())
				break
			}
		}
	}
}

func (g *Game) countIntersections() int {
	sum := 0
	PIPE := int64(35)
	for y := 1; y < len(g.board)-1; y++ {
		for x := 1; x < len(g.board[y])-1; x++ {
			if g.board[y][x] == PIPE && g.board[y+1][x] == PIPE && g.board[y-1][x] == PIPE && g.board[y][x-1] == PIPE && g.board[y][x+1] == PIPE {
				sum = sum + (x * y)
			}
		}
	}
	return sum
}

func (g *Game) identifyLegs() []Leg {
	var pos Point
	var dir Dir
	for y, row := range g.board {
		for x, value := range row {
			switch value {
			case '^', '>', '<', 'v':
				pos = Point{x, y}
				switch value {
				case '^':
					dir = UP
				case '>':
					dir = RIGHT
				case '<':
					dir = LEFT
				case 'v':
					dir = DOWN
				}
			}
		}
	}
	// walk scaffold and record journey
	legs := make([]Leg, 0)
	turn := -1
	done := false
	for !done {
		// work out next move
		// from current point where is next neighbouring #
		switch {
		case dir != RIGHT && pos.x > 0 && g.board[pos.y][pos.x-1] == '#':
			switch {
			case dir == UP:
				turn = L
			case dir == DOWN:
				turn = R
			default:
				done = true
			}
			dir = LEFT
		case dir != LEFT && pos.x < len(g.board[pos.y])-1 && g.board[pos.y][pos.x+1] == '#':
			switch {
			case dir == UP:
				turn = R
			case dir == DOWN:
				turn = L
			default:
				done = true
			}
			dir = RIGHT
		case dir != UP && pos.y < len(g.board)-1 && g.board[pos.y+1][pos.x] == '#':
			switch {
			case dir == LEFT:
				turn = L
			case dir == RIGHT:
				turn = R
			default:
				done = true
			}
			dir = DOWN
		case dir != DOWN && pos.y > 0 && g.board[pos.y-1][pos.x] == '#':
			switch {
			case dir == LEFT:
				turn = R
			case dir == RIGHT:
				turn = L
			default:
				done = true
			}
			dir = UP
		default:
			done = true
		}
		if !done {
			count := 0
			// go to end of current line
			for {
				nextPos := Point{pos.x + dir.dx, pos.y + dir.dy}
				if nextPos.y < 0 || nextPos.y >= len(g.board) || nextPos.x < 0 || nextPos.x >= len(g.board[nextPos.y]) {
					break
				}
				next := g.board[nextPos.y][nextPos.x]
				if next != '#' {
					break
				}
				pos = nextPos
				count++
			}
			legs = append(legs, Leg{turn, count})
		}
	}
	s := fmt.Sprintf("Walking map identified %v legs\n%v", len(legs), legs)
	fmt.Println(s)
	g.graphics.SetContent(s)
	return legs
}

func part1(d *display.Display) {
	p := intcomp.ReadProgram("input")
	go p.ProcessProgram()
	g := Game{graphics: d, in: p.Output, out: p.Input}
	g.readBoard(-1)
	fmt.Println("Part 1", g.countIntersections())
}

func (g *Game) processInstructions(input string) {
	for _, value := range input {
		g.out <- int64(value)
	}
	g.out <- 10
}

func part2(d *display.Display) {
	defer d.Close()
	p := intcomp.ReadProgram("input")
	p.Poke(2, 0)
	go p.ProcessProgram()
	g := Game{graphics: d, in: p.Output, out: p.Input}
	go g.processInput()
	for g.state != STATE_READY {
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
	g.identifyLegs()
	time.Sleep(time.Duration(1) * time.Second)

	mainProg := "A,A,B,C,B,C,B,C,B,A"
	progA := "R,6,L,12,R,6"
	progB := "L,12,R,6,L,8,L,12"
	progC := "R,12,L,10,L,10"
	g.processInstructions(mainProg)
	g.processInstructions(progA)
	g.processInstructions(progB)
	g.processInstructions(progC)
	g.processInstructions("y")
	g.state = STATE_RUNNING

	for g.state != STATE_COMPLETE {
		time.Sleep(time.Duration(10) * time.Millisecond)
	}
	time.Sleep(time.Duration(2) * time.Second)
}

func runParts(d *display.Display) {
	part1(d)
	part2(d)
}

func main() {
	d := display.CreateDisplay(600, 800)
	go runParts(d)
	d.Show()
}
