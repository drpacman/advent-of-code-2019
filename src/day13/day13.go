package main

import (
	"display"
	"fmt"
	"intcomp"
	"strings"
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
	var output strings.Builder
	for y := 0; y < len(board); y++ {
		output.WriteString("<div><pre>")
		for x := 0; x < len(board[y]); x++ {
			entry := " "
			switch board[y][x] {
			case 1:
				entry = "X"
			case 2:
				entry = "+"
			case 3:
				entry = "="
			case 4:
				entry = "O"
			}
			output.WriteString(entry)
		}
		output.WriteString("</pre></div>")
	}
	return output.String()
}

func part2(display *display.Display) {
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
				display.SetContent(generateBoard(board))
				time.Sleep(time.Duration(1 * time.Millisecond))
			}
		}
	}
	fmt.Println("Part 2: ", score)
	display.Close()
}

func main() {
	part1()
	display := display.CreateDisplay(800, 800)
	go part2(display)
	display.Show()
}
