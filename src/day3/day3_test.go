package main

import (
	"testing"
)

func TestWireCross(t *testing.T) {
	wireA := [4]Wire{Wire{'R', 8}, Wire{'U', 5}, Wire{'L', 5}, Wire{'D', 3}}
	wireB := [4]Wire{Wire{'U', 7}, Wire{'R', 6}, Wire{'D', 4}, Wire{'L', 4}}
	crossPointDist := closestCrossPoint(wireA[:], wireB[:])
	if crossPointDist != 6 {
		t.Errorf("unexpected cross point distance, %v instead of %v", crossPointDist, 6)
	}
}

func TestWireMinimumSteps(t *testing.T) {
	wireA := [4]Wire{Wire{'R', 8}, Wire{'U', 5}, Wire{'L', 5}, Wire{'D', 3}}
	wireB := [4]Wire{Wire{'U', 7}, Wire{'R', 6}, Wire{'D', 4}, Wire{'L', 4}}
	crossPointSteps := shortestStepsCrossPoint(wireA[:], wireB[:])
	if crossPointSteps != 30 {
		t.Errorf("unexpected shorted steps cross point, %v instead of %v", crossPointSteps, 30)
	}
}
func TestWireCoords(t *testing.T) {
	wireA := Wires{Wire{'R', 8}, Wire{'U', 5}, Wire{'L', 5}, Wire{'D', 3}}
	wireCoords := wireA.getCoords()
	if len(wireCoords) != 21 {
		t.Errorf("unexpected number of crossing points, %v instead of %v", len(wireCoords), 21)
	}
}
