package main

import (
	"errors"
	"advent-of-code-2019/files"
	"fmt"
	"strings"
)

func readInput() map[string]string {
	lines := files.ReadLines("input")
	orbits := make(map[string]string)
	for _, s := range lines {
		items := strings.Split(s, ")")
		orbits[items[1]] = items[0]
	}
	return orbits
}

func distance(orbits map[string]string, key string, d int) int {
	next, ok := orbits[key]
	if ok {
		return distance(orbits, next, d+1)
	} else {
		return d
	}
}

func distanceTo(orbits map[string]string, key string, dest string, d int) int {
	next, ok := orbits[key]
	if ok && next != dest {
		return distanceTo(orbits, next, dest, d+1)
	} else {
		return d
	}
}

func ancestors(orbits map[string]string, key string, a []string) []string {
	next, ok := orbits[key]
	if ok {
		return ancestors(orbits, next, append(a, key))
	} else {
		return append(a, key)
	}
}

func countOrbits(orbits map[string]string) int {
	count := 0
	for key := range orbits {
		count = count + distance(orbits, key, 0)
	}
	return count
}

func getCommonAncestor(a, b []string) (string, error) {
	for _, s := range a {
		for _, t := range b {
			if t == s {
				return t, nil
			}
		}
	}
	return "", errors.New("No common ancestor")
}

func countOrbitTransfers(orbits map[string]string, src, dest string) int {
	// find ancestors
	var srcAncestors = ancestors(orbits, src, []string{})
	var destAncestors = ancestors(orbits, dest, []string{})
	var commonAncestor, _ = getCommonAncestor(srcAncestors, destAncestors)
	// count both legs
	return distanceTo(orbits, src, commonAncestor, 0) + distanceTo(orbits, dest, commonAncestor, 0)
}

func main() {
	orbits := readInput()
	fmt.Printf("Part 1: %v\n", countOrbits(orbits))
	fmt.Printf("Part 2: %v\n", countOrbitTransfers(orbits, "YOU", "SAN"))
}
