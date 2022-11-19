package main

import (
	"testing"
)

func RunTest(t *testing.T, filename string, expected int) {
	m, start := ReadMaze(filename)
	var path = m.ProcessMaze(start)
	var path_length = len(path) - 1
	if expected != path_length {
		t.Errorf("Expected %v but got %v for filename %v", expected, path_length, filename)
	}
}

// func TestProcess1(t *testing.T) { RunTest(t, "input_test_1", 8) }

// func TestProcess2(t *testing.T) { RunTest(t, "input_test_2", 86) }

// func TestProcess3(t *testing.T) { RunTest(t, "input_test_3", 81) }

// func TestProcess4(t *testing.T) { RunTest(t, "input_test_4", 136) }

// func TestProcess5(t *testing.T) { RunTest(t, "input_test_5", 132) }

func TestProcess(t *testing.T) { RunTest(t, "input", 136) }
