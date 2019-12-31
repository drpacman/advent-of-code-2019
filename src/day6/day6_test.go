package main

import (
	"testing"
)

func makeTestOrbits() map[string]string {
	var orbits = make(map[string]string)
	orbits["B"] = "COM"
	orbits["C"] = "B"
	orbits["D"] = "C"
	orbits["E"] = "D"
	orbits["F"] = "E"
	orbits["G"] = "B"
	orbits["H"] = "G"
	orbits["I"] = "D"
	orbits["J"] = "E"
	orbits["K"] = "J"
	orbits["L"] = "K"
	return orbits
}

func TestDistance(t *testing.T) {
	var orbits = makeTestOrbits()
	var n = distance(orbits, "L", 0)
	if n != 7 {
		t.Errorf("Distance for orbit L should be %v but was %v", 7, n)
	}
}

func TestDistanceTo(t *testing.T) {
	var orbits = makeTestOrbits()
	var n = distanceTo(orbits, "L", "E", 0)
	if n != 2 {
		t.Errorf("Distance from L to E should be %v but was %v", 2, n)
	}
}

func TestOrbitCount(t *testing.T) {
	var orbits = makeTestOrbits()
	var n = countOrbits(orbits)
	if n != 42 {
		t.Errorf("Count of orbits should be %v but was %v", 42, n)
	}
}

func TestOrbitTransfers(t *testing.T) {
	var orbits = makeTestOrbits()
	orbits["YOU"] = "K"
	orbits["SAN"] = "I"
	var transfers = countOrbitTransfers(orbits, "YOU", "SAN")
	if transfers != 4 {
		t.Errorf("Transfers of orbits should be %v but was %v", 4, transfers)
	}
}

func TestAncestors(t *testing.T) {
	var orbits = makeTestOrbits()
	orbits["YOU"] = "K"
	var a = ancestors(orbits, "YOU", []string{})
	var expected = []string{"YOU", "K", "J", "E", "D", "C", "B", "COM"}
	for i, v := range expected {
		if a[i] != v {
			t.Errorf("ancestors of YOU should be %v but was %v", expected, a)
		}
	}
}
