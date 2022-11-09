package main

import (
	"fmt"
	"advent-of-code-2019/intcomp"
	"advent-of-code-2019/render"
)

func calculateTractorBeam(path string, startX,startY,dx,dy int) int {
	count := 0
	board := make([][]int64, dy)
	for y := startY; y < startY+dy; y++ {
		board[y-startY] = make([]int64, dx)
		for x := startX; x < startX+dx; x++ {
			p := intcomp.ReadProgram(path)
			go p.ProcessProgram()
			p.Input <- int64(x)
			p.Input <- int64(y)
			value := <-p.Output
			if value == 1 {
				count++
			}
			board[y-startY][x-startX] = value
		}
	}
	chars := make(map[int64]string)
	chars[0] = "."
	chars[1] = "#"
	fmt.Println(render.GenerateBoardText(board, chars))
	return count
}

func part1() {
	fmt.Println("Part1", calculateTractorBeam("input",0,0,50,50))
}

func scanUntil(start, y int, stopAt int) int {
	x := start
	for {
		p := intcomp.ReadProgram("input")
		go p.ProcessProgram()
		p.Input <- int64(x)
		p.Input <- int64(y)
		value := <-p.Output
		if value == int64(stopAt) {
			return x
		}
		x++		
	}
}

func part2() {
	lxA, rxA, lxB, y := 5, 5, 5, 5
	for {
		lxA = scanUntil( lxA, y, 1 )
		rxA = scanUntil( lxA+1, y, 0 )
		lxB= scanUntil( lxB, y+99, 1 )
		
		if lxB == rxA - 100 {
			posX := rxA-100
			calculateTractorBeam("input",posX,y,100,100)
			fmt.Println("Part2", posX * 10000 + y)
			break
		}
		y++
	}
}

func main() {
	part1()
	part2()
}
