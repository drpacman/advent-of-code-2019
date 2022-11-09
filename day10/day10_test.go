package main

import (
	"math"
	"testing"
)

func testAngle(x2, y2 int, expected float64, t *testing.T) {
	grad := angle(0, 0, x2, y2)
	if grad != expected {
		t.Errorf("Unexpected angle for (0,0) to (%v,%v) of %v - expected %v", x2, y2, grad, expected)
	}
}
func TestAngle(t *testing.T) {
	testAngle(0, -1, 0, t)
	testAngle(1, -1, math.Pi/4, t)
	testAngle(1, 0, math.Pi/2, t)
	testAngle(1, 1, math.Pi*3/4, t)
	testAngle(0, 1, math.Pi, t)
	testAngle(-1, 1, math.Pi*5/4, t)
	testAngle(-1, 0, math.Pi*3/2, t)
	testAngle(-1, -1, math.Pi*7/4, t)

}

var asteroidMap = AsteroidMap{contents: [][]int{[]int{0, 1, 0, 0, 1},
	[]int{0, 0, 0, 0, 0},
	[]int{1, 1, 1, 1, 1},
	[]int{0, 0, 0, 0, 1},
	[]int{0, 0, 0, 1, 1}}}

func TestLineOfSightSize(t *testing.T) {
	lineOfSight := asteroidMap.calculateLineOfSight(Coordinate{1, 0})
	if len(lineOfSight) != 7 {
		t.Errorf("Line of sight should have length 7 but we got %v\n %v", len(lineOfSight), lineOfSight)
	}
}

func TestFindBestAsteroid(t *testing.T) {
	asteroid := asteroidMap.findBestAsteroid()
	expected := Coordinate{3, 4}
	if asteroid.location != expected {
		t.Errorf("Did not find correct asteroid - found %v when should have found %v", asteroid.location, expected)
	}
}

func TestFindBestAsteroidFromFile(t *testing.T) {
	asteroidMap := readAsteroidMap("test_input")
	asteroid := asteroidMap.findBestAsteroid()
	expected := Coordinate{11, 13}
	if asteroid.location != expected {
		t.Errorf("Did not find correct asteroid - found %v when should have found %v", asteroid.location, expected)
	}
}

func TestDetectedAsteroids(t *testing.T) {
	asteroidMap := readAsteroidMap("test_input2")
	asteroid := asteroidMap.findBestAsteroid()
	detected := len(asteroid.lineOfSight)
	expected := 33
	if detected != expected {
		t.Errorf("Did not detect correct number of asteroids - found %v when should have found %v", detected, expected)
	}
}

func testTarget(pos int, expected Coordinate, targets []Coordinate, t *testing.T) {
	if targets[pos] != expected {
		t.Errorf("Did not detect correct target asteroid for target number %v - found %v when should have found %v\nTargets:\n %v", pos, targets[pos], expected, targets)
	}
}
func TestVaporisedAsteroids(t *testing.T) {
	asteroidMap := readAsteroidMap("test_input3")
	asteroid := asteroidMap.findBestAsteroid()
	expected := Coordinate{8, 3}
	if asteroid.location != expected {
		t.Errorf("Did not detect correct number of asteroids - found %v when should have found %v", asteroid.location, expected)
	}
	targets := asteroid.getTargets()
	testTarget(0, Coordinate{8, 1}, targets, t)
	testTarget(8, Coordinate{15, 1}, targets, t)
}
