package main

import (
	"fmt"
	"reflect"
	"testing"
)

func makeTestMaze() Maze {
	var contents = make(map[Point]int)
	contents[Point{-2, -2}] = 1
	contents[Point{1, 0}] = 1
	contents[Point{-1, 1}] = 1
	return Maze{
		contents: contents,
		droid:    Point{0, 0}}
}
func TestMazeRender(t *testing.T) {
	maze := makeTestMaze()
	output := maze.generateBoard()
	expected := [][]int64{{WALL, 0, 0, 0},
		{0, 0, 0, 0},
		{0, 0, DROID, WALL},
		{0, WALL, 0, 0}}

	if reflect.DeepEqual(output, expected) == false {
		t.Errorf("generated board \n%v \nexpected \n%v", output, expected)
	}
	fmt.Println(maze.render())
}

func TestMazeMoveWall(t *testing.T) {
	maze := makeTestMaze()
	maze.applyInformation(SOUTH, WALL)
	result, ok := maze.contents[Point{0, 1}]
	if !ok || result != WALL {
		t.Errorf("Moved south and got %v instead of %v", result, WALL)
	}
	expected := Point{0, 0}
	if maze.droid != expected {
		t.Errorf("Droid should not move - got %v instead of %v", maze.droid, expected)
	}
	fmt.Println("TestMazeMoveWall - Maze\n", maze.render())
}

func TestMazeMoveWater(t *testing.T) {
	maze := makeTestMaze()
	maze.applyInformation(SOUTH, EMPTY)
	result := maze.contents[Point{0, 1}]
	if result != EMPTY {
		t.Errorf("Moved south and go %v instead of %v", result, WALL)
	}
	expected := Point{0, 1}
	if maze.droid != expected {
		t.Errorf("Droid should not move - got %v instead of %v", maze.droid, expected)
	}
	fmt.Println("Maze\n", maze.render())
}
