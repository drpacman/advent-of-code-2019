package main

import (
	"testing"
)

func checkDimension(dimension string, actual Dimensions, expected Dimensions, t *testing.T) {
	if actual != expected {
		t.Errorf("Expected moon to have %v %v but got %v", dimension, expected, actual)
	}
}

func checkVelocity(moon Moon, expected Dimensions, t *testing.T) {
	checkDimension("velocity", moon.velocity, expected, t)
}

func checkPosition(moon Moon, expected Dimensions, t *testing.T) {
	checkDimension("position", moon.position, expected, t)
}

func TestVelocityChanges(t *testing.T) {
	moon1 := Moon{Dimensions{-1, 0, 2}, Dimensions{0, 0, 0}}
	moon2 := Moon{Dimensions{2, -10, -7}, Dimensions{0, 0, 0}}
	moon1.applyGravity(moon2)
	checkVelocity(moon1, Dimensions{1, -1, -1}, t)
}

func TestRound(t *testing.T) {
	moon1 := Moon{Dimensions{-1, 0, 2}, Dimensions{0, 0, 0}}
	moon2 := Moon{Dimensions{2, -10, -7}, Dimensions{0, 0, 0}}
	moon3 := Moon{Dimensions{4, -8, 8}, Dimensions{0, 0, 0}}
	moon4 := Moon{Dimensions{3, 5, -1}, Dimensions{0, 0, 0}}

	runTimeStep([]*Moon{&moon1, &moon2, &moon3, &moon4})
	checkVelocity(moon1, Dimensions{3, -1, -1}, t)
	checkPosition(moon4, Dimensions{2, 2, 0}, t)
}
