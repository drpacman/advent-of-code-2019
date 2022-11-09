package main

import (
	"advent-of-code-2019/display"
	"flag"
	"fmt"
	"advent-of-code-2019/intcomp"
	"advent-of-code-2019/render"
	"time"
)

type Coordinate struct {
	x int64
	y int64
}

func part1() {
	p := intcomp.ReadProgram("input")
	go p.ProcessProgram()
	var tile int64
	count := 0
	for !p.Halted {
		_ = <-p.Output
		_ = <-p.Output
		tile = <-p.Output
		if tile == 2 {
			count++
		}
	}
	fmt.Println("Part 1: ", count)
}

func generateBoard(board [][]int64) string {
	var chars = map[int64]string{
		1: "X",
		2: "+",
		3: "=",
		4: "O"}

	return render.GenerateBoardHtml(board, chars)
}

func part2(display *display.Display, delay int64) {
	if display != nil {
		defer display.Close()
	}
	p := intcomp.ReadProgram("input")
	p.Poke(2, 0)
	go p.ProcessProgram()
	var x, y, tile int64 = 0, 0, 0
	board := make([][]int64, 26)
	for i, _ := range board {
		board[i] = make([]int64, 46)
	}
	var ball = Coordinate{0, 0}
	var bat = Coordinate{0, 0}
	var score int64 = 0

	for !p.Halted {
		x = <-p.Output
		y = <-p.Output
		if x == -1 && y == 0 {
			score = <-p.Output
		} else {
			tile = <-p.Output
			board[y][x] = tile
			switch tile {
			case 3:
				bat = Coordinate{x, y}
			case 4:
				ball = Coordinate{x, y}
				switch dist := bat.x - ball.x; {
				case dist < 0:
					p.Input <- 1
				case dist > 0:
					p.Input <- -1
				default:
					p.Input <- 0
				}
				if display != nil {
					display.SetContent(generateBoard(board))
					time.Sleep(time.Duration(int64(delay)) * time.Millisecond)
				}
			}
		}
	}
	fmt.Println("Part 2: ", score)

}

func main() {
	delay := flag.Int("delay", 1, "How long to pause between frames")
	flag.Parse()

	fmt.Println(*delay)

	part1()
	if *delay != 0 {
		d := display.CreateDisplay(800, 800)
		go part2(d, int64(*delay))
		d.Show()
	} else {
		part2(nil, 0)
	}
}
